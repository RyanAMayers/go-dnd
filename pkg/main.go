package main

import (
	"fmt"

	"RyanAMayers/godnd"
)

func main() {
	firstRoll := godnd.D20()

	fmt.Printf("You rolled a %d!\n", firstRoll)
}
