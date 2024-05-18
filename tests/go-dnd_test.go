package godnd_test

import (
	"fmt"
	"testing"

	gameActions "github.com/RyanAMayers/godnd/gameActions"
)

func TestRollDice(t *testing.T) {
	// Test rolling a single 6-sided dice
	result := gameActions.RollSingle(6)
	fmt.Println("TEST: RollSingle(6) - 1d6 - ", result)
	if result < 1 || result > 6 {
		t.Errorf("RollSingle(6) returned an invalid result: %d", result)
	}

	// Test rolling multiple 10-sided dice
	total, rolls := gameActions.RollN(10, 3)
	fmt.Println("TEST: RollN(10, 3) - 3d10 - ", total, rolls)
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
	fmt.Println("TEST: RollN(6, -4) - (-4)d6 - ", total, rolls)
	if total != 0 {
		t.Errorf("RollN(6, -4) returned an invalid total: %d", total)
	}
	if len(rolls) != 0 {
		t.Errorf("RollN(6, -4) returned invalid rolls: %v", rolls)
	}

	// Test rolling a die with negative sides
	total, rolls = gameActions.RollN(-6, 3)
	fmt.Println("TEST: RollN(-6, 3) - 3d(-6) - ", total, rolls)
	if total != 0 {
		t.Errorf("RollN(-6, 3) returned an invalid total: %d", total)
	}
	if len(rolls) != 0 {
		t.Errorf("RollN(-6, 3) returned invalid rolls: %v", rolls)
	}

	// Test rolling a 0-sided dice
	total, rolls = gameActions.RollN(0, 3)
	fmt.Println("TEST: RollN(0, 3) - 3d0 - ", total, rolls)
	if total != 0 {
		t.Errorf("RollN(0, 3) returned an invalid total: %d", total)
	}
	if len(rolls) != 0 {
		t.Errorf("RollN(0, 3) returned invalid rolls: %v", rolls)
	}

	// Test rolling a d20
	result = gameActions.D20()
	fmt.Println("TEST: D20() - 1d20 - ", result)
	if result < 1 || result > 20 {
		t.Errorf("D20() returned an invalid result: %d", result)
	}

	// Test rolling with advantage
	result = gameActions.RollAdvantage()
	fmt.Println("TEST: RollAdvantage() - 2d20 - ", result)
	if result < 1 || result > 20 {
		t.Errorf("RollAdvantage() returned an invalid result: %d", result)
	}

	// Test rolling with disadvantage
	result = gameActions.RollDisadvantage()
	fmt.Println("TEST: RollDisadvantage() - 2d20 - ", result)
	if result < 1 || result > 20 {
		t.Errorf("RollDisadvantage() returned an invalid result: %d", result)
	}

	// Test rolling with a die string
	total, rolls = gameActions.Roll("3d6")
	fmt.Println("TEST: Roll(\"3d6\") - 3d6 - ", total, rolls)
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
	fmt.Println("TEST: Roll(\"3d\") - 3d - ", total, rolls)
	if total != 0 {
		t.Errorf("Roll(\"3d\") returned an invalid total: %d", total)
	}
	if len(rolls) != 0 {
		t.Errorf("Roll(\"3d\") returned invalid rolls: %v", rolls)
	}

	// Test rolling with a high number of dice
	total, rolls = gameActions.Roll("100d100")
	fmt.Println("TEST: Roll(\"100d100\") - 100d100 - ", total, rolls)
	if total < 100 || total > 10000 {
		t.Errorf("Roll(\"100d100\") returned an invalid total: %d", total)
	}
	for _, roll := range rolls {
		if roll < 1 || roll > 100 {
			t.Errorf("Roll(\"100d100\") returned an invalid roll: %d", roll)
		}
	}
}

func TestHeroRolls(t *testing.T) {

	result := gameActions.Check("Athletics", mitch.AbilityScores)
	if result < 1 || result > 30 {
		t.Errorf("Check(\"Athletics\") returned an invalid result: %d", result)
	}
	fmt.Println("TEST: Mitch - Check(\"Athletics\") - Athletics - ", result)
}
