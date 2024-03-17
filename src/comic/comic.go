package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber uint
	var grade float32

	// Book 1
	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5

	fmt.Println(title, "\nWritten by: ", writer, "\nArtist: ", artist, "\nPublisher: ", publisher, "\nYear: ", year, "\n# of Pages: ", pageNumber, "\nGrade: ", grade)

	// Book 2
	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println("")
	fmt.Println("")

	fmt.Println(title, "\nWritten by: ", writer, "\nArtist: ", artist, "\nPublisher: ", publisher, "\nYear: ", year, "\n# of Pages: ", pageNumber, "\nGrade: ", grade)
}
