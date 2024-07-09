package main

import (
	"fmt"
	"math"
)

// Define the function y = x^(2)- 4 x + 7
func f(x float64) float64 {
	return math.Pow(x, 2) - 4*x + 7
}

// Bisection method to find the minimum
func bisection(a, b, epsilon float64) float64 {
	for math.Abs(b-a) >= epsilon {
		c := (a + b) / 2
		if f(a) < f(b) {
			b = c // Root lies between a and c
		} else {
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

	fmt.Printf("Approximate minimum value: y = %.3f at x = %.3f\n", minimumValue, minimum)
}
