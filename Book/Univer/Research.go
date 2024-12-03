package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Define the function y = x^2 - 4x + 7
func f(x float64) float64 {
	return x*x - 4*x + 7
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

func main() {
	//Create a "buffered reader" to receive text from the keyboard.
	reader := bufio.NewReader(os.Stdin)
	// Request a value from the user.
	fmt.Print("Lower bound: ")
	// Returns the text entered by the user before pressing the Enter key. Будет прочитан весь текст до руны новой строки.
	inputL, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err) //Report an error and stop the program.
	}
	fmt.Print("Upper bound: ")
	inputU, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err) //Report an error and stop the program.
	}
	// Removing the newline character from the input string
	inputL = strings.TrimSpace(inputL)
	inputU = strings.TrimSpace(inputU)

	//The arguments pass the converted string "input" and the number in float format
	a, err := strconv.ParseFloat(inputL, 32)

	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.ParseFloat(inputL, 32)
	if err != nil {
		log.Fatal(err)
	}

	// a := 0.     // Lower bound
	// b := 3.     // Upper bound
	tol := 1e-6 // Tolerance (adjust as needed)

	minimum := goldenSectionSearch(a, b, tol)
	fmt.Printf("Minimum unction value: y = %.3f at x = %.3f\n", f(minimum), minimum)
}
