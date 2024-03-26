/* Codecademy
LEARN GO: LOOPS, ARRAYS, MAPS, AND STRUCTS
Currency Converter
In this project, we will be setting up an interactive currency converter.

You will input a dollar amount, then a target currency. The currency converter will convert your dollars into the target currency and display the result to you.

We will need to write code for the following:

Creating a map of currency conversion rates.
Inputting an amount and a currency type.
Using the map to convert the dollars.
*/

package main

import (
	"fmt"
)

func main() {
	// Create a map of currency conversion rates.
	currencies := map[string]float32{
		"JPY": 130.2,
		"EUR": 0.95,
	}
	fmt.Println(currencies)

	var dollarAmount float32
	var currency string

	// Get dollar amount from user
	fmt.Print("Enter a dollar amount: ")
	fmt.Scan(&dollarAmount)

	if dollarAmount == 0 {
		fmt.Println("Invalid stock price.")
	} else {
		// Get currency from user
		fmt.Print("Enter currency: ")
		fmt.Scan(&currency)

		// Use the map to convert the dollars.
		rate, isValid := currencies[currency]

		// Check if valid
		if isValid {
			rate *= dollarAmount
			fmt.Println(rate)
		} else {
			fmt.Println("Currency is not available.")
		}

	}

}
