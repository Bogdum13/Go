// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	// Define the search range
// 	startX := -math.Pi / 4
// 	endX := 3 * math.Pi / 2
// 	step := 0.001 // Adjust the step size as needed

// 	// Initialize variables for maximum and minimum values
// 	maxY := math.Inf(-1)
// 	minY := math.Inf(1)
// 	maxX := 0.0
// 	minX := 0.0

// 	// Brute-force search
// 	for x := startX; x <= endX; x += step {
// 		y := -math.Sin(2*x) - math.Cos(3*x)

// 		// Update maximum and minimum values
// 		if y > maxY {
// 			maxY = y
// 			maxX = x
// 		}
// 		if y < minY {
// 			minY = y
// 			minX = x
// 		}
// 	}

// 	// Print results
// 	fmt.Printf("Maximum value: y = %.4f at x = %.4f\n", maxY, maxX)
// 	fmt.Printf("Minimum value: y = %.4f at x = %.4f\n", minY, minX)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	// Define the search range
// 	startX := 0.
// 	endX := 3.
// 	step := 0.001 // Adjust the step size as needed

// 	// Initialize variables for maximum and minimum values
// 	maxY := math.Inf(-1)
// 	minY := math.Inf(1)
// 	maxX := 0.0
// 	minX := 0.0

// 	// Brute-force search
// 	for x := startX; x <= endX; x += step {
// 		y :=  math.Pow(x, 2) - 4 * x + 7

// 		// Update maximum and minimum values
// 		if y > maxY {
// 			maxY = y
// 			maxX = x
// 		}
// 		if y < minY {
// 			minY = y
// 			minX = x
// 		}
// 	}

// 	// Print results
// 	fmt.Printf("Maximum value: y = %.3f at x = %.3f\n", maxY, maxX)
// // 	fmt.Printf("Minimum value: y = %.3f at x = %.3f\n", minY, minX)
// // }

// package main

// // import (
// // 	"fmt"
// 	"math"
// )

// // Define the function y = x^2 - 4x + 7
// func f(x float64) float64 {
// 	return math.Pow(x, 2) - 4*x + 7
// }

// // Fibonacci search method
// func fibonacciSearch(a, b float64, n int) float64 {
// 	// Calculate the Fibonacci sequence up to n
// 	fib := make([]float64, n+1)
// 	fib[0], fib[1] = 1, 1
// 	for i := 2; i <= n; i++ {
// 		fib[i] = fib[i-1] + fib[i-2]
// 		fmt.Println(fib[i])
// 	}

// 	// Perform the search
// 	for k := 1; k < n; {
// 		c := a + (fib[n-k-1]/fib[n-k+1])*(b-a)
// 		if f(b) > f(c) {
// 			b = c
// 		} else if f(b) <= f(c) {
// 			a = c
// 		}
// 		k++
// 	}

// 	return (a + b) / 2
// }

// func main() {
// 	a := 0.0 // Start of the interval
// 	b := 3.0 // End of the interval
// 	n := 50  // Number of Fibonacci numbers to use (adjust as needed)

// 	minimum := fibonacciSearch(a, b, n)
// 	minimumValue := f(minimum)

// 	fmt.Printf("Approximate minimum value: y = %.4f at x = %.4f\n", minimumValue, minimum)
// }
