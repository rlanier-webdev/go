package main

import (
	"fmt"
)

func askOrder() string {
	var input string
	fmt.Print("What would you like to eat: ")
	// Get the input from the user
	fmt.Scan(&input)
	return input
}

// Helper function checks the menu
func contains(menu []string, order string) bool {
	for _, item := range menu {
		if order == item {
			return true
		}
	}
	return false
}

func main() {
	fastfoodMenu := []string{"Burgers", "Nuggets", "Fries"}
	fmt.Println("The fast food menu has these items:", fastfoodMenu)

	var total int
	var order string
	// ask for orders
	for {
		order = askOrder()
		if order == "quit" {
			break // Exit the loop if the user wants to quit
		}
		if contains(fastfoodMenu, order) {
			total += 4
		} else {
			fmt.Println("Sorry, that item is not on the menu.")
		}
	}

	// count $2 bills
	for amount := total; amount > 0; amount -= 2 {
		fmt.Println("You paid with $2 bills. Remaining amount: $", amount)
	}

	fmt.Println("The total for the order is $", total)
}
