/* Codecademy
LEARN GO: LOOPS, ARRAYS, MAPS, AND STRUCTS
Catnapped
In this project, we will be setting up a single-player mystery board game.

A number of guests are attending a party in your parlor when the lights go out. You hear an angry “MEOW!” When the lights come back on, you realize — your cat is gone! It’s up to you to find out who hid your cat and where the cat is hidden!

Imagine that we have a gameboard with a collection of locations and guests.

In order to make our game work we need some lists:

The guests in the house
The objects the cat could be hidden inside of
Each game, we would also need to create a solution to the mystery. The solution would contain:

The guest who stole the cat
The object the cat is in
In this project, we will create these lists and the solution to our mystery!
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Setting Up The House
	guest := [...]string{"Jill Scott", "Max Dot", "Beyonce", "Prince"}
	objects := [...]string{"toy chest", "crate", "box", "cabinet", "trunk", "closet"}

	fmt.Println("Guest:", guest, "\nObjects:", objects)
	
	// convert arrays to slices
	gSlice := guest[:]
	oSlice := objects[:]

	// get random elements
	culprit := getRandomElement(gSlice)
	secretObject := getRandomElement(oSlice)

	// solution
	solution := fmt.Sprintln(culprit, "hid the cat by putting it in the", secretObject)
	fmt.Print(solution)
}

func getRandomElement(slice []string) string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random index between 0 and the length of the array
	index := rand.Intn(len(slice))

	// Return the slice element at that index.
	return slice[index]

}
