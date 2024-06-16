/* 
Given an array of integers.

Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.

Example
For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].
*/

package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}

	for _, value := range SumNegatives(arr) {
		fmt.println(value)
	}
	for _, value := range CountPositives(arr) {
		fmt.println(value)
	}
}

func SumNegatives(arr []int) int {
	var sum int

	for _, value := range arr {
		if value < 0
			sum += value
	fmt.print(sum)
}

func CountPositives(arr []int) int {
	var count int

	for _, value := range arr {
		if value > 0
			count++
	}
	fmt.print(count)
}