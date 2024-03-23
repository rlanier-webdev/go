package main

import (
	"fmt"
	"math/rand"
)

func main() {
	isHeistOn := true

	eludedGuards := eludeGuards()

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := openVault(isHeistOn)

	if openedVault {
		leftSafely := escapeTheVault()
		if leftSafely {
			amtStolen := stealMoney()
			fmt.Printf("Success! You stole $%d\n", amtStolen)
		}
	} else {
		fmt.Println("You've failed.")
	}
}

func eludeGuards() int {
	var i int

	for {
		fmt.Println("Enter a number between 1 and 100 (enter 0 to quit):")
		fmt.Scan(&i)

		if i == 0 {
			break
		}

		if i < 0 || i > 100 {
			fmt.Println("Please enter a number between 1 and 100 or enter 0 to quit.")
			continue
		}

		break
	}

	return i
}

func openVault(isHeistOn bool) bool {
	if !isHeistOn {
		return false
	}

	openedValut := rand.Intn(100)

	if openedValut >= 70 {
		fmt.Println("Vault opened!")
		return true
	} else {
		fmt.Println("Vault can't be opened.")
		return false
	}
}

func escapeTheVault() bool {
	// 50% chance
	escaped := rand.Intn(2) == 0

	if !escaped {
		fmt.Println("Failed to escape.")
		return false
	}

	vaultMessages := []string{
		"Hold on, vault doors are tricky...",
		"One more try should do it...",
	}
	messageIndex := rand.Intn(len(vaultMessages))
	fmt.Println(vaultMessages[messageIndex])

	fmt.Println("Escape successful! Start the car!")
	return true
}

func stealMoney() int {
	total := 10000 + rand.Intn(1000000)
	return total
}
