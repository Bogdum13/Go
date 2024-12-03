// package main

// import (
// 	"fmt"
// 	// "log"
// 	// "os"
// )

// func main() {
// 	// fmt.Println("Cannon\nall!!! gjldt\" ghnrt\\\tu")
// 	// fmt.Println('в') // возвращает числовой код символа в Юникоде
// 	// fmt.Println('\n')
// 	// fmt.Println(true)
// 	// fmt.Println(33.3 / 7)
// 	// fmt.Println(2+2 != 5)
// 	task()
// }

// func task() {

// fmt.Println(8 <= 6)
// Стр 52
// var originalCount, eatenCount int = 10, 4
// // var originalCount, eatenCount int = 10, 4

// fmt.Println("I started with", originalCount, "apples.\nSome jerk ate", eatenCount, "apples.\nThere are", originalCount-eatenCount,
// 	"apples left.")

// // //Page 54
// quantity := 4
// length, width := 1.2, 2.4
// customerName := "Damon Cole"
// fmt.Println(customerName)
// fmt.Println("has ordered", quantity, "sheets")
// fmt.Println("each with an area of")
// fmt.Println(length*width, "square meters")

// Page 58
// var price int = 100
// fmt.Println("Price is", price, "dollars.")
// var taxRate float64 = 0.08
// var tax float64 = float64(price) * taxRate
// fmt.Println("Tax is", tax, "dollars.")
// var total float64 = float64(price) + tax
// fmt.Println("Total cost is", total, "dollars.")
// var availableFunds int = 120
// fmt.Println(availableFunds, "dollars available.")
// fmt.Println("Within budget?", total <= float64(availableFunds))

// Page 77
// fileInfo, err := os.Stat("go.mod")
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Println(fileInfo.Size())

// Page 98

// for x := 1; x >= 3; x++ {
// 	fmt.Print(x)
// }
// fmt.Print("\n")
// for x := 3; x >= 1; x-- {
// 	fmt.Print(x)
// 	}

// page 117
//The first field must have a minimum width of 12 characters. The width is not defined for the second field
// fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
// fmt.Println("-----------------------------")
// // Again, the minimum width is 12. The minimum width is 2 for the second field
// fmt.Printf("%12s | %2d\n", "Stamps", 50)
// fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
// fmt.Printf("%12s | %2d\n", "Tape", 99)

// fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
// fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
// fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
// fmt.Printf("%%.1f: %.1f\n", 12.3456)
// fmt.Printf("%%.2f: %.2f\n", 12.3456)
// fmt.Println([]byte("Hello Б 宝"))

// 	fmt.Println("Hi")

// }

// func fun() {
// 	var myFunction func()
// 	myFunction = task
// 	myFunction()
// 	}

// page 437

// func callFunction(passedFunction func()) {
// 	passedFunction()
// 	}
// 	func callTwice(passedFunction func()) {
// 	passedFunction()
// 	passedFunction()
// 	}
// 	func callWithArguments(passedFunction func(string, bool)) {
// 	passedFunction("This sentence is", false)
// 	}
// 	func printReturnValue(passedFunction func() string) {
// 	fmt.Println(passedFunction())
// 	}
// 	func functionA() {
// 	fmt.Println("function called")
// 	}
// 	func functionB() string {
// 	fmt.Println("function called")
// 	return "Returning from function"
// 	}
// 	func functionC(a string, b bool) {
// 	fmt.Println("function called")
// 	fmt.Println(a, b)
// 	}
// 	func main() {
// 	callFunction(functionA)
// 	callTwice(functionA)
// 	callWithArguments(functionC)
// 	printReturnValue(functionB/)
// /}

// package main
// import (
// "log"
// "os"
// "text/template"
// )

// }
// func main() {
// text := "Here's my template!\n"
// tmpl, err := template.New("test").Parse(text)
// check(err)
// err = tmpl.Execute(os.Stdout, nil)
// check(err)
// }

// func main() {
// 	templateText := "Template start\nAction: {{.}}\nTemplate end\n"
// 	tmpl, err := template.New("test").Parse(templateText)
// 	check(err)
// 	err = tmpl.Execute(os.Stdout, "ABC")
// 	check(err)
// 	err = tmpl.Execute(os.Stdout, 42)
// 	check(err)
// 	err = tmpl.Execute(os.Stdout, true)
// 	check(err) }

// page 497
package main

import (
	// "bufio"
	// "text/template"
	"html/template"
	"log"
	// "net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Invoice struct {
	Name    string
	Paid    bool
	Charges []float64
	Total   float64
}

func main() {
	html, err := template.ParseFiles("bill.html")
	check(err)
	bill := Invoice{
		Name:    "Mary Gibbs",
		Paid:    true,
		Charges: []float64{23.19, 1.13, 42.79},
		Total:   67.11,
	}
	err = html.Execute(os.Stdout, bill)
	check(err)
}
