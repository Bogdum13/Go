package main

//Необходимо импортировать пакет «time», чтобы использовать тип time.Time.
import (
	"fmt"
	// "reflect"
	// "strings"
	// "time"
)

/*
	func main() {
		var now time.Time = time.Now() // Метод time.Now возвращает значение time.Time,	представляющее текущую дату и время
		var year int = now.Year()      // У значений time.Time имеется метод Year, который возвращает текущий год.
		fmt.Println(year)

		now2 := time.Now()
		year2 := now2.Year()
		day := now2.Day()
		fmt.Println(year2, day, "\n\n")

		broken := "G# r#cks!"
		// Возвращает значение strings.Replacer, настроенное для замены всех «#» в строке на «o».
		replacer := strings.NewReplacer("#", "o")

		// Вызываем метод Replace для strings.Replacer и передаем ему строку, в которой должны осуществляться замены.
		fixed := replacer.Replace(broken)

		// Вывод строки, возвращенной методом Replace.
		fmt.Println(fixed, reflect.TypeOf(replacer))
	}
*/
func main() {

	// Избегайте замещения имен
	var /*int*/ count int = 12
	var /*append*/ suffix string = "minutes of bonus footage"
	var /*fmt*/ format string = "DVD"
	var languages = append([]string{}, "Español")
	// fmt.Println(int, append, "on", fmt, languages)
	fmt.Println(count, suffix, "on", format, languages)

}
