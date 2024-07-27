package main

import "fmt"

// Define the Person struct
type Person struct {
	Name  string
	Age   int
	Email string
}

// Implement the PrintDetails method for Person
func (p Person) PrintDetails() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Printf("Email: %s\n", p.Email)
}

func main() {
	// Create an instance of Person
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	// Print the details of the person
	person.PrintDetails()
}
