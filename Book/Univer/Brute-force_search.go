package main

import (
	"fmt"
	"math"
)

func main() {
	// Define the search range
	startX := 0.
	endX := 3.
	step := 0.001 // Adjust the step size as needed

	// Initialize variables for maximum and minimum values
	minY := math.Inf(1)
	minX := 0.0

	// Brute-force search
	for x := startX; x <= endX; x += step {
		y := math.Pow(x, 2) - 4*x + 7

		// Update minimum values
		if y < minY {
			minY = y
			minX = x
		}
	}

	// Print results
	fmt.Printf("Minimum value: y = %.3f at x = %.3f\n", minY, minX)
}
