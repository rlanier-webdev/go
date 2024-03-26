package main

import "fmt"

type Book struct {
	publisher  string
	writer     string
	artist     string
	title      string
	genre      string
	year       uint
	pageNumber uint
	grade      float32
}

func main() {

	bookOne := Book{
		"Mr. GoToSleep",
		"Tracey Hatchet",
		"Jewel Tampson",
		"DizzyBooks Publishing Inc.",
		"Horror",
		1997,
		14,
		6.5,
	}
	fmt.Println(bookOne)

	bookTwo := Book{
		"Dolly Vol. 1",
		"Ray N. Shawn",
		"Joy Nails",
		"Ckode Publishing Inc.",
		"Non Fiction",
		2024,
		18,
		9.0,
	}
	fmt.Println(bookTwo)

	bookThree := Book{
		"Epic Vol. 1",
		"Ryan N. Shawn",
		"Phoebe Paperclips",
		"DizzyBooks Publishing Inc.",
		"Crime",
		2013,
		160,
		9.0,
	}
	fmt.Println(bookThree)
}
