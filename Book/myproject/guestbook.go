// page 481 //
package main

import (
	"bufio"
	// "text/template"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	// The new type is defined at the beginning of guestbook.go.
	SignatureCount int
	Signatures     []string
}

// The error message code is moved to this function
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// As usual, the ResponseWriter value is passed to the handler functions.. ..as well as a pointer to the Request value
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt") // Adding the getstring call
	//fmt.Printf("%#v\n", signatures) // We display the uploaded records.
	html, err := template.ParseFiles("view.html")
	check(err)
	// Creating a new Guestbook structure
	guestbook := Guestbook{
		// The Signature Count field stores the length of the record segment.
		SignatureCount: len(signatures),
		//The Signature s field stores the record segment itself.
		Signatures: signatures,
	}
	// The structure is passed to the Execute method of the Template value.
	err = html.Execute(writer, guestbook)
	check(err)
}


////////////////////////////////////////////////////////////////////////////


// Adding another handler function with the same parameters as the ViewHandler
func newHandler(writer http.ResponseWriter, request *http.Request) {
	// Downloads the content new.html as the template text
	html, err := template.ParseFiles("new.html")
	check(err)
	// Writing the template in response (you don't need to insert the data yet)
	err = html.Execute(writer, nil)
	check(err)
}

// Defines another request handler function with similar parameters
func createHandler(writer http.ResponseWriter, request *http.Request) {
	// Gets the value of the "signature" field of the form.
	signature := request.FormValue("signature")
	// File Opening Options
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	// Open file
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	// Writes the text in a new line of the file.
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	// page 506
	http.Redirect(writer, request, "/guestbook", http.StatusFound)

	// // Writes the value of the field in the response.
	// _, err := writer.Write([]byte(signature))
	// check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	// We assign the new Handler function to process requests with the path "/guestbook/new".
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
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
