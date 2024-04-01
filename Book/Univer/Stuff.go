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