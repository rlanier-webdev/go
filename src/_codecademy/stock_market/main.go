/*
LEARN GO: LOOPS, ARRAYS, MAPS, AND STRUCTS
Stock Market
In this project, we will represent a stock market in the Go language. We will create the following:

A struct that represents a stock.
A slice of Stock structs that will represent a stock market.
A function that displays the stock market (stock names and prices).
A function to randomly move a stock’s price up or down.
A main which tests our functionality.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGen(min float32, max float32) float32 {
	return min + (max-min)*rand.Float32()
}

// Task implementation goes here

type Stock struct {
	name  string
	price float32
}

func (stock Stock) updateStock() {
	change := randomNumberGen(-10000, 10000)
	stock.price = change
}

func displayMarket(market Stock) {
	for i := 0; i < len(market); i++ {
		fmt.Println(market)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Function calls go here
	stockMarket := []{{"GOOG", 2313.50},{"AAPL", 157.28},{"FB", 203.77},{"TWTR", 50.00}}

	fmt.Println(stockMarket)
}
