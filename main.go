package main

import (
	"fmt"

	gameActions "github.com/RyanAMayers/godnd/gameActions"
)

func main() {
	firstRoll := gameActions.D20()

	fmt.Printf("You rolled a %d!\n", firstRoll)
}
