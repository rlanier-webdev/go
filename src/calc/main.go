package main

import (
	"fmt"
	"math/rand"
)

func isUndefined(y int) bool {
	if y == 0 {
		return true
	}
	return false
}
func add(x, y int) int {
	return x + y
}
func subtract(x, y int) int {
	return x - y
}
func multipy(x, y int) int {
	return x * y
}
func divide(x, y int) float32 {
	if isUndefined(y) {
		return 0
	} else {
		return float32(x / y)
	}

}

func main() {
	x := rand.Intn(10)
	y := rand.Intn(20)

	// Addition
	add := add(x, y)
	fmt.Println("Result of", x, "+", y, "=", add)

	// Subtraction
	sub := subtract(x, y)
	fmt.Println("Result of", x, "-", y, "=", sub)

	// Multiplication
	mult := multipy(x, y)
	fmt.Println("Result of", x, "*", y, "=", mult)

	// Division
	// redo this part...figure out the best approach
	div := divide(x, y)
	fmt.Printf("\nResult of %d / %d = %.3f", x, y, div)
}
