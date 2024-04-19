package guest

import (
	"fmt"
	"math/rand"
	"time"
)

func getGuest() []string {
	guest := [...]string{"Jill Scott", "Max Dot", "Beyonce", "Prince"}

	fmt.Println("Guest:", guest)

	gSlice := guest[:]

	return gSlice
}