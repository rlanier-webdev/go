//Create a program that converts temperatures between Celsius and Fahrenheit.
//Fahrenheit °F to Celsius °C Formula: C = (°F - 32) ÷ 1.8

package main

import "fmt"

func main() {
	c := convert(75.5)
	s := fmt.Sprintf("%.2f", c)
	fmt.Println(s)

}

func convert(f float32) float32 {
	c := (f - 32) / 1.8
	return c
}
