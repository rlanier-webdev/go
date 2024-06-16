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
	var input int
	x := rand.Intn(10)
	
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)

	// Addition
	add := add(x, input)
	fmt.Println("Result of", x, "+", input, "=", add)

	// Subtraction
	sub := subtract(x, input)
	fmt.Println("Result of", x, "-", input, "=", sub)

	// Multiplication
	mult := multipy(x, input)
	fmt.Println("Result of", x, "*", input, "=", mult)

	// Division
	// redo this part...figure out the best approach
	div := divide(x, input)
	fmt.Printf("\nResult of %d / %d = %.3f", x, input, div)
}
