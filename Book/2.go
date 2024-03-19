package main

//Альтернативный формат инструкции «import» позволяет импортировать сразу несколько пакетов.
import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect" // Импортируем пакет 	«reflect», чтобы использовать его функцию TypeOf.
)

func main() {
	// Функция fmt.Println вызывается для возвращаемого значения функции math.Floor.
	// Получает число, округляет его в меньшую сторону и возвращает полученное значение.
	fmt.Println(math.Floor(2.75))
	// strings.Title Получает строку и возвращает новую строку, в которой все слова начинаются с буквы верхнего регистра.
	fmt.Println(strings.Title("head fitst go"))
	fmt.Println(cases.Title(language.Und).String("head fitst go"))
	task1()
	various()
}

func task1() {
	//Go — язык со статической типизацией. Если вы используете неправильный тип значения в неподходящем месте, Go сообщит вам об этом.
	//fmt.Println(math.Floor("head first go"))
	//fmt.Println(strings.Title(2.75))

	// Чтобы узнать тип любого значения, передайте его функции TypeOf из пакета reflect
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))

}

func various() {
	// Переменная представляет собой область памяти, в которой хранится значение. Чтобы к переменной можно было обращаться по имени, используйтся
	// объявление переменной. Объявление состоит из ключевого слова var, за которым следует имя и тип значений, которые будут храниться в переменной.

	// Объявление переменных
	/*var quantity int
	var length, width float64
	var customerName string

	// присваивание значений переменным
	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Damon Cole"*/

	// Объявление переменных И присваивание значений
	// var quantity int = 4
	// var length, width float64 = 1.2, 2.4
	// var customerName string = "Damon Cole"

	// Если значение переменной присваивается одновременно с ее объявлением, тип переменной в объявлении обычно не указывают.
	// quantity := 4
	// length, width := 1.2, 2.4
	// customerName := "Damon Cole"

	// // использование переменных
	// fmt.Println(customerName)
	// fmt.Println("has ordered", quantity, "sheets")
	// fmt.Println("each with an area of")
	// fmt.Println(length*width, "square meters")

	// // Если переменная объявляется без присваивания значения, то она будет содержать нулевое значение для этого типа.
	// var myInt int
	// var myFloat float64
	// var myString string
	// var myBool bool
	// fmt.Println("myInt:", myInt, "; myFloat:", myFloat, "; myString:", myString, "; myBool:", myBool)

	// var length float64 = 1.2
	// var width int = 2
	// length = float64(width) // преобразование значения одного типа к другому типу.Следует указать тип, к которому должно быть преобразовано значение,
	// // а за ним преобразуемое значение в скобках.
	// fmt.Println(length)

	// var myInt int = 2
	// fmt.Println(reflect.TypeOf(myInt))
	// fmt.Println(reflect.TypeOf(float64(myInt)))

	// var length float64 = 1.2
	// var width int = 2
	// fmt.Println("Area is", length*float64(width))
	// fmt.Println("length > width?", length > float64(width))

	myInt := 2.
	fmt.Println(reflect.TypeOf(myInt))

	var length float64 = 3.75
	var width int = 5
	width = int(length)
	fmt.Println(width)

}
