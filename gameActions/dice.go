package gameActions

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func Roll(dieString string) (total int, rolls []int) {
	die, n := parseDieString(dieString)
	return RollN(die, n)
}

func RollSingle(die int) int {
	if die < 1 {
		return 0
	}
	return rand.Intn(die) + 1
}

func RollN(die, n int) (total int, rolls []int) {
	rolls = []int{}
	total = 0
	if n < 1 || die < 1 {
		return total, rolls
	}
	for i := 0; i < n; i++ {
		rolls = append(rolls, RollSingle(die))
		total += rolls[i]
	}
	return total, rolls
}

func D20() int {
	return RollSingle(20)
}

func RollAdvantage() int {
	roll1 := D20()
	roll2 := D20()
	if roll1 > roll2 {
		return roll1
	}
	return roll2
}

func RollDisadvantage() int {
	roll1 := D20()
	roll2 := D20()
	if roll1 < roll2 {
		return roll1
	}
	return roll2
}

func parseDieString(dieString string) (int, int) {
	die := 0
	n := 1
	fmt.Sscanf(dieString, "%dd%d", &n, &die)
	return die, n
}
