package main

import (
	"fmt"
)

// read write protection
func sendMessage(msgch<- chan string) {
	msg := <- msgch
	//msgch <- "hello!"
}

// Buffered2
func main() {
	userch := make(chan string, 2)

	userch <- "Bob" //blocking
	userch <- "Ray" //blocking

	user := <-userch
	fmt.Print(user)
}

// Buffered
/* func main() {
	userch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		userch <- "Bob" //blocking
	}()
	user := <-userch
	fmt.Print(user)
}
*/
