package main

//Альтернативный формат инструкции «import» позволяет импортировать сразу несколько пакетов.
import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// Функция fmt.Println вызывается для возвращаемого значения функции math.Floor.
	// Получает число, округляет его в меньшую сторону и возвращает полученное значение.
	fmt.Println(math.Floor(2.75))
	// strings.Title Получает строку и возвращает новую строку, в которой все слова начинаются с буквы верхнего регистра.
	fmt.Println(strings.Title("head fitst go"))
	fmt.Println(cases.Title(language.Und).String("head fitst go"))
}
