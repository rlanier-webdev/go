// Write a program to print the multiplication table of a given number up to 10.

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number:")
	fmt.Scan(&num)

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n",num, i, num * i)
	}
}