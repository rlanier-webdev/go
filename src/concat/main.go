package main

import "fmt"

func main() {
	fmt.Println(concat("Hello, ", "world!"))
}

func concat(s1, s2 string) string {
	return s1 + s2
}