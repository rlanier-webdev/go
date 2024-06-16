package main

import (
	"fmt"
)

func main() {
/* 	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 0))
	fmt.Println(SquareSum([]int{1, 2, 2}))
	fmt.Println(numToString(123)) */
	//string to reverse
	s := "A string to reverse with golang."
	r := []rune(s)
	fmt.Printf("text reversed %s\n", reverse(r))
}

/* func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend/divisor, nil
}

func SquareSum(numbers []int) int {
	// your code here
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	return sum
}

func numToString(num int) string {
	s := fmt.Sprint(num)
	return s
} */

func reverse(r []rune) string {
	// your code here
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	result := string(r)
	return result
}