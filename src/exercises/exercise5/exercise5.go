package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel to communicate the result
	result := make(chan int)

	// Start the factorial calculation in a goroutine
	go func() {
		result <- factorial(5)
	}()

	// Calculate and print the factorial
	fmt.Println(factorial(12))

	// Receive and print the result from the goroutine
	fmt.Println(<-result)
}

// factorial function calculates the factorial of n.
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		time.Sleep(100 * time.Millisecond)
		return n * factorial(n-1)
	}

}