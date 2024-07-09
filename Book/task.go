package main

import (
	"fmt"
	// "log"
	// "os"
)

func main() {
	// fmt.Println("Cannon\nall!!! gjldt\" ghnrt\\\tu")
	// fmt.Println('в') // возвращает числовой код символа в Юникоде
	// fmt.Println('\n')
	// fmt.Println(true)
	// fmt.Println(33.3 / 7)
	// fmt.Println(2+2 != 5)
	task()
}

func task() {
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

	for x := 1; x >= 3; x++ {
		fmt.Print(x)
	}
    fmt.Print("\n")
	for x := 3; x >= 1; x-- {
		fmt.Print(x)
		}

}
