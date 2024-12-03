package main

import (
	"fmt"
	"math"
)

// Define the function y = x^2 - 4x + 7
func f(x float64) float64 {
	return - math.Sqrt(20 - math.Pow((x-1), 2))  //x*x - 4*x + 7
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
	a := 0. // Lower bound
	b := 4. // Upper bound
	tol := 1e-6 // Tolerance (adjust as needed)

	minimum := goldenSectionSearch(a, b, tol)
	fmt.Printf("Minimum unction value: y = %.3f at x = %.3f\n", f(minimum), minimum)
}
