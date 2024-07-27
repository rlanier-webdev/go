package main

import (
	"errors"
	"fmt"
	"regexp"
)

// Define the Person struct
type Person struct {
	Name  string
	Age   int
	Email string
}

// Method to initialize a Person with validation
func NewPerson(name string, age int, email string) (*Person, error) {
	if name == "" {
		return nil, errors.New("name cannot be blank")
	}
	if age < 0 {
		return nil, errors.New("age cannot be negative")
	}
	if !validateEmail(email) {
		return nil, errors.New("email is not valid")
	}
	return &Person{Name: name, Age: age, Email: email}, nil
}

// Implement the PrintDetails method for Person
func (p Person) PrintDetails() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Printf("Email: %s\n", p.Email)
}

func main() {
	// Create an instance of Person
	person, err := NewPerson("John Doe", 15, "john.doe@example.com")
	if err != nil {
		fmt.Println("Error creating person:", err)
		return
	}

	// Print the details of the person
	person.PrintDetails()
}

func validateEmail(email string) bool {
	// Define the email regex pattern
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex pattern
	re := regexp.MustCompile(emailPattern)

	// Use the regex to match the email string
	return re.MatchString(email)
}
