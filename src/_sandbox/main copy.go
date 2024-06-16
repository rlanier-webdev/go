package main

import (
	"fmt"
	"os/user"
)

type messageToSend struct {
	message string
	sender user
	recipient user
}

type user struct {
	name string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	// ?
	return true
}
func main() {


}