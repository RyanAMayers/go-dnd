package godnd_test

import (
	"testing"

	"github.com/RyanAMayers/godnd/gameActions"
)

func TestRollDice(t *testing.T) {
	// Test rolling a single 6-sided dice
	result := gameActions.RollSingle(6)
	if result < 1 || result > 6 {
		t.Errorf("RollSingle(6) returned an invalid result: %d", result)
	}

	// Test rolling multiple 10-sided dice
	total, rolls := gameActions.RollN(10, 3)
	if total < 3 || total > 30 {
		t.Errorf("RollN(10, 3) returned an invalid total: %d", total)
	}
	for _, roll := range rolls {
		if roll < 1 || roll > 10 {
			t.Errorf("RollN(10, 3) returned an invalid roll: %d", roll)
		}
	}

	// Test rolling a negative number of dice
	total, rolls = gameActions.RollN(6, -4)
	if total != 0 {
		t.Errorf("RollN(6, -4) returned an invalid total: %d", total)
	}
	if len(rolls) != 0 {
		t.Errorf("RollN(6, -4) returned invalid rolls: %v", rolls)
	}

	// Test rolling a 0-sided dice
	total, rolls = gameActions.RollN(0, 3)
	if total != 0 {
		t.Errorf("RollN(0, 3) returned an invalid total: %d", total)
	}
	if len(rolls) != 0 {
		t.Errorf("RollN(0, 3) returned invalid rolls: %v", rolls)
	}

	// Test rolling a d20
	result = gameActions.D20()
	if result < 1 || result > 20 {
		t.Errorf("D20() returned an invalid result: %d", result)
	}

	// Test rolling with advantage
	result = gameActions.RollAdvantage()
	if result < 1 || result > 20 {
		t.Errorf("RollAdvantage() returned an invalid result: %d", result)
	}

	// Test rolling with disadvantage
	result = gameActions.RollDisadvantage()
	if result < 1 || result > 20 {
		t.Errorf("RollDisadvantage() returned an invalid result: %d", result)
	}

	// Test rolling with a die string
	total, rolls = gameActions.Roll("3d6")
	if total < 3 || total > 18 {
		t.Errorf("Roll(\"3d6\") returned an invalid total: %d", total)
	}
	for _, roll := range rolls {
		if roll < 1 || roll > 6 {
			t.Errorf("Roll(\"3d6\") returned an invalid roll: %d", roll)
		}
	}

	// Test rolling with an invalid die string
	total, rolls = gameActions.Roll("3d")
	if total != 0 {
		t.Errorf("Roll(\"3d\") returned an invalid total: %d", total)
	}
	if len(rolls) != 0 {
		t.Errorf("Roll(\"3d\") returned invalid rolls: %v", rolls)
	}
}
