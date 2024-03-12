package main

import (
	"fmt"      // fmt спецальный покет придоставляющий доступ к функциям для вывода текста в терминале
	"net/http" // net/http за счёт этого покета что либо показывать пользоателю, и отслеживать куда пользователь перешёл на странице
)

func home_page(w http.ResponseWriter, r *http.Request) { // при помощи параметра page ResponseWriter мы сможем обращиться к определённой страничке и что либо показывать пользователю, параметр r Request
	fmt.Fprintf(w, "Go is super easy %s!", r.URL.Path[1:])
}

func Bogdum_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bogdum page!")
}

/*
func main() {
	//fmt.Println("Go рулит!")
	http.HandleFunc("/", home_page)   // принимает два параметра, и позволяет отследить переход по определённому url адресу / , по переходу по которому будет вызват метод home_page
	http.ListenAndServe(":8080", nil) // локальный сервер, первый параметр порт по каторому будем слушать наш локальный сервер, вторым параметром передаются настройки для самого сервера (nil -> nul ничего)
}*/

func handleRequest() {
	http.HandleFunc("/", home_page) // принимает два параметра, и позволяет отследить переход по определённому url адресу / , по переходу по которому будет вызват метод home_page
	http.HandleFunc("/Bogdum/", Bogdum_page)
	http.ListenAndServe(":8080", nil) // локальный сервер, первый параметр порт по каторому будем слушать наш локальный сервер, вторым параметром передаются настройки для самого сервера (nil -> nul ничего)
}

func main() {
	handleRequest()
}
