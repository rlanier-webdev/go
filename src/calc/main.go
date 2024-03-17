package main

import "fmt"

func main() {
	x := 13
	y := 31

	// Addition
	add := x + y
	fmt.Println("Result of ", x, " + ", y, " =", add)

	// Subtraction
	sub := add - y
	fmt.Println("Result of ", add, " - ", y, " =", sub)

	// Multiplication
	mult := sub * x
	fmt.Println("Result of ", sub, " * ", x, "=", mult)

	// Division
	div := mult / y
	fmt.Println("Result of ", mult, " / ", y, " = ", div)
}
