package guest

import (
	"fmt"
	"math/rand"
	"time"
)

func getObject() []string {
	objects := [...]string{"toy chest", "crate", "box", "cabinet", "trunk", "closet"}

	fmt.Println("\nObjects:", objects)

	oSlice := objects[:]

	return oSlice
}