package main

import (
	"fmt"
	"math"
)

// Define the function y = -(sin(2x) + cos(3x))
func f(x float64) float64 {
	return math.Pow(x, 2) - 4 * x + 7 //math.Pow(x, 3) + 2 * math.Pow(x, 2) + x - 1
}

// Bisection method to find the minimum
func bisection(a, b, epsilon float64) float64 {
	for math.Abs(b-a) >= epsilon {
		c := (a + b) / 2
		//fmt.Println(c)
		if f(a) < f(b) {
			b = c // Root lies between a and c
		} else if f(a) >= f(b) {
			fmt.Println(c)
			a = c // Root lies between b and c
		} 
	}
	return (a + b) / 2 // Approximate root
}
func main() {
	a := 0.
	b := 3.
	epsilon := 0.0001 // Adjust the precision as needed

	minimum := bisection(a, b, epsilon)
	minimumValue := f(minimum)

	fmt.Printf("Approximate minimum value: y = %.4f at x = %.4f\n", minimumValue, minimum)
}