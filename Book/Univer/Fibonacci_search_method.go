package main

import (
	"fmt"
	"math"
)

// Define the function y = x^2 - 4x + 7
func f(x float64) float64 {
	return math.Pow(x, 2) - 4*x + 7
}

// Fibonacci search method
func fibonacciSearch(a, b float64, n int) float64 {
	// Calculate the Fibonacci sequence up to n
	fib := make([]float64, n+1)
	fib[0], fib[1] = 1, 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	// Perform the search
	for k := 1; k < n; {
		c_1 := a + (fib[n-k-1]/fib[n-k+1])*(b-a)
		c_2 := a + (fib[n-k]/fib[n-k+1])*(b-a)
		if f(c_1) < f(c_2) {
			b = c_2
		} else {
			a = c_1
		}
		k++
	}

	return (a + b) / 2
}

func main() {
	a := 0. // Start of the interval
	b := 3. // End of the interval
	n := 50  // Number of Fibonacci numbers to use (adjust as needed)

	minimum := fibonacciSearch(a, b, n)
	minimumValue := f(minimum)

	fmt.Printf("Approximate minimum value: y = %.4f at x = %.4f\n", minimumValue, minimum)
}
