/*
1. check to see if array contains a 0
2. finding the sum of all the elements
3. updating a value for a particular element
4. searching for a particular value within the array
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x, y, z := rand.Intn(10), rand.Intn(10), rand.Intn(10)

	arr := []int{x, y, z}
	
	// updating a value for a particular element
	if arr[1] == 5 {
		arr[1] = 25
		fmt.Println(arr)
	}

	fmt.Println("Sum of all elements:", findSum(arr))

	if hasZero(arr) {
		fmt.Println("Array contains zero:", arr)
	} else {
		fmt.Println("Array doesn't contain any zeros:", arr)
	}

	
}

// checks if the array contains a zero value and returns true if it does

func hasZero(arr []int) bool {
	for _, value := range arr {
		if value == 0 {
			return true
		}
	}
	return false
}

// finding the sum of all the elements
func findSum(arr [] int) int {
	var sum int
	for _, value := range arr {
		sum += value
	}
	return sum
}


