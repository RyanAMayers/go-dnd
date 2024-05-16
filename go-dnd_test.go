package godnd_test

import (
	"testing"

	. "github.com/RyanAMayers/godnd"
)

func TestRollDice(t *testing.T) {
	// Test rolling a single 6-sided dice
	result := RollDice(6)
	if result < 1 || result > 6 {
		t.Errorf("RollDice(6) returned an invalid result: %d", result)
	}

	// Test rolling multiple 10-sided dice
	result = RollDice(10, 3)
	if result < 3 || result > 30 {
		t.Errorf("RollDice(10, 3) returned an invalid result: %d", result)
	}

	// Test rolling a negative number of dice
	result = RollDice(-4)
	if result != 0 {
		t.Errorf("RollDice(-4) returned an invalid result: %d", result)
	}
}
