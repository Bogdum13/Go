package main

import (
	"log"
	"net/http"
)

// writer is a value for updating	the response that will be sent to the browser. request is a value representing a request from the browser.
func write(writer http.ResponseWriter, message string) {
	// The string is converted to a byte segment, as before, and written to the response.
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!") // We add the line "Hello, web!" in response.
}
func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut web!")
}
func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, web!")
}

func main() {
	// If a request was received for URL ending with the suffix "/hello" then the function is calle viewHandler that generates the response.
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("/namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil) // Processes the browser request and responds to it.
	log.Fatal(err)
}



