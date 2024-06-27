package main

import "fmt"

var w, l int = 5, 10

func calculateRecArea() {
	// A = wl
	area := w * l
	fmt.Println(area)
}

func main() {
	calculateRecArea()
}
