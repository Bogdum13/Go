package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

const (
	host    = "localhost"
	port    = 5432
	user    = "postgres"
	pasword = "BogdanBogdum"
	dbname  = "app_db"
)

type App struct {
	// The new type is defined at the beginning of app.go.
	AnswersCount int
	Answers      []string
}

// The error message code is moved to this function
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// As usual, the ResponseWriter value is passed to the handler functions.. ..as well as a pointer to the Request value
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	// answers := getStrings("add_number.txt") // Adding the getstring call

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pasword, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	// insertDunStmt := `select from "add_number" ("lower_bound", "upper_bound", "minimum_y", "minimum_x") values($1, $2, $3, $4)`
	// _, err = db.Exec(insertDunStmt, lower_bound, upper_bound, f(minimum), minimum)
	// CheckError(err)

	// Выполнение SELECT-запроса
	rows, err := db.Query("SELECT lower_bound, upper_bound, minimum_y, minimum_x FROM add_number")
	if err != nil {
		log.Fatalf("Failed to execute query: %v\n", err)
	}
	defer rows.Close()

	type Add_number struct {
		Lower_bound float32
		Upper_bound float32
		Minimum_y   float32
		Minimum_x   float32
	}

	var add_number []Add_number
	for rows.Next() {
		var u Add_number
		err := rows.Scan(&u.Lower_bound, &u.Upper_bound, &u.Minimum_y, &u.Minimum_x)
		if err != nil {
			log.Fatalf("Failed to scan row: %v\n", err)
		}
		add_number = append(add_number, u)
	}

	// Обработка ошибок сканирования строк
	if err := rows.Err(); err != nil {
		log.Fatalf("Failed to iterate over rows: %v\n", err)
	}

	// Вывод результатов
	for _, add_number := range add_number {
		fmt.Printf("Lower_bound: %f Upper_bound: %f Minimum_y: %f Minimum_x: %f\n", add_number.Lower_bound,
			add_number.Upper_bound, add_number.Minimum_y, add_number.Minimum_x)
	}

	// Получение количества строк
	var coun int
	err = db.QueryRow("SELECT COUNT(*) FROM add_number").Scan(&coun)
	if err != nil {
		log.Fatalf("Failed to execute query: %v\n", err)
	}

	// Формирование HTML-кода
	htmlOutput := "<h1>Нахождение минимума функции</h1>"
	htmlOutput += "<!-- The number of entries from the Guestbook structure is inserted. -->"
	htmlOutput += fmt.Sprintf("Всего ответов - %d", coun)
	htmlOutput += "<br><a href=\"/answer/new\">Добавить значения</a>"

	for _, add_number := range add_number {
		htmlOutput += fmt.Sprintf("<br>Для диапазона ( %0.1f ; %0.1f ) Минимальное значение функции: y = %0.4f x = %0.4f\n",
			add_number.Lower_bound, add_number.Upper_bound, add_number.Minimum_y, add_number.Minimum_x)
	}

	htmlOutput += "</table></body></html>"

	// Сохранение HTML-кода в файл
	file, err := os.Create("view.html")
	if err != nil {
		log.Fatalf("Failed to create file: %v\n", err)
	}
	defer file.Close()

	writerr := bufio.NewWriter(file)
	_, err = writerr.WriteString(htmlOutput)
	if err != nil {
		log.Fatalf("Failed to write to file: %v\n", err)
	}
	writerr.Flush()

	fmt.Println("HTML file saved as view.html")

	//fmt.Printf("%#v\n", answers) // We display the uploaded records.
	html, err := template.ParseFiles("view.html")
	check(err)
	// Creating a new App structure

	// The structure is passed to the Execute method of the Template value.
	err = html.Execute(writer, nil)
	check(err)

	// answers := getStrings("add_number.txt") // Adding the getstring call

	// //fmt.Printf("%#v\n", answers) // We display the uploaded records.
	// html, err := template.ParseFiles("view.html")
	// check(err)
	// // Creating a new App structure
	// app := App{
	// 	// The Answers Count field stores the length of the record segment.
	// 	AnswersCount: len(answers),
	// 	//The Answers s field stores the record segment itself.
	// 	Answers: answers,
	// }
	// // The structure is passed to the Execute method of the Template value.
	// err = html.Execute(writer, app)
	// check(err)
}

// Adding another handler function with the same parameters as the ViewHandler
func newHandler(writer http.ResponseWriter, request *http.Request) {
	// Downloads the content new_answer.html as the template text
	html, err := template.ParseFiles("add_number.html")
	check(err)
	// Writing the template in response (you don't need to insert the data yet)
	err = html.Execute(writer, nil)
	check(err)
}

// Defines another request handler function with similar parameters
func createHandler(writer http.ResponseWriter, request *http.Request) {
	// Gets the value of the "answer" field of the form.
	lower_bound := request.FormValue("lower_bound")
	upper_bound := request.FormValue("upper_bound")

	// page 506
	http.Redirect(writer, request, "/answer", http.StatusFound)

	a, err := strconv.ParseFloat(lower_bound, 2) // Lower bound
	b, err := strconv.ParseFloat(upper_bound, 2) // Upper bound
	tol := 1e-6                                  // Tolerance (adjust as needed)

	minimum := goldenSectionSearch(a, b, tol)

	fmt.Printf("Минимальное значение функции: y = %.3f at x = %.3f\n", f(minimum), minimum)

	// File Opening Options
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	// Open file
	file, err := os.OpenFile("add_number.txt", options, os.FileMode(0600))
	check(err)
	// Writes the text in a new line of the file.
	_, err = fmt.Fprintln(file, "Для диапазона (", lower_bound, "; ", upper_bound,
		") Минимальное значение функции:", "y =", f(minimum), " x =", minimum)

	// _, err = fmt.Fprintf(file, "y = %.3f at x = %.3f\n", f(minimum), minimum)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pasword, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertDunStmt := `insert into "add_number" ("lower_bound", "upper_bound", "minimum_y", "minimum_x") values($1, $2, $3, $4)`
	_, err = db.Exec(insertDunStmt, lower_bound, upper_bound, f(minimum), minimum)
	CheckError(err)

	check(err)
	err = file.Close()
	check(err)
}

////////////////////////////////////////////////////////////////////////////////////////////

// Define the function y = x^2 - 4x + 7
func f(x float64) float64 {
	return -math.Sqrt(20 - math.Pow((x-1), 2)) //x*x - 4*x + 7
}

// Golden-section search to find the minimum value
func goldenSectionSearch(a, b float64, tol float64) float64 {
	phi := (1 + math.Sqrt(5)) / 2 // Golden ratio

	// Initial points
	x1 := b - (b-a)/phi
	x2 := a + (b-a)/phi

	for math.Abs(b-a) > tol {
		if f(x1) < f(x2) {
			b = x2
			x2 = x1
			x1 = b - (b-a)/phi
		} else {
			a = x1
			x1 = x2
			x2 = a + (b-a)/phi
		}
	}

	return (a + b) / 2
}

//////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pasword, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt := `insert into "Employee"("Name", "EmpId") values('Bogdum', 13)`
	_, err = db.Exec(insertStmt)
	CheckError(err)

	insertDunStmt := `insert into "Employee"("Name", "EmpId") values($1, $2)`
	_, err = db.Exec(insertDunStmt, "Lu", 16)
	CheckError(err)

	http.HandleFunc("/answer", viewHandler)
	// We assign the new Handler function to process requests with the path "/answer/new".
	http.HandleFunc("/answer/new", newHandler)
	http.HandleFunc("/answer/create", createHandler)
	er := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(er)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	// Opens the file
	// If an error is returned indicating that the file does not exist return nil instead of a line segment.
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	// For any other type of error, report it and shut down.
	check(err)
	// After exiting the function, make sure that the file is closed.
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// Report any file processing errors and shut down.
	check(scanner.Err())
	return lines
}
