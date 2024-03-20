package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	rand.Seed(time.Now().UnixNano())

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedValut := rand.Intn(100)

	if isHeistOn && openedValut >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Vault can't be opened.")
	}

	leftSafely := rand.Intn(5)
	
	if isHeistOn {
		switch leftSafely {
			case 0:
				isHeistOn = false
				fmt.Println("Failed")
			case 1:
				isHeistOn = false
				fmt.Println("Turns out vault doors don't open from the inside...")
			case 2:
				isHeistOn = false
				fmt.Println("Close but not yet...")
			case 3:
				isHeistOn = false
				fmt.Println("Still failed bruh...")
			default:
				fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println(amtStolen)
	}
	
	fmt.Println(isHeistOn)
}
