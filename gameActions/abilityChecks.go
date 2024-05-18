package gameActions

import (
	godnd "github.com/RyanAMayers/godnd/models"
)

func getAbilityScore(ability string, AbilityScores godnd.AbilityScores) int {
	switch ability {
	case "STR":
		return AbilityScores.STR
	case "DEX":
		return AbilityScores.DEX
	case "CON":
		return AbilityScores.CON
	case "INT":
		return AbilityScores.INT
	case "WIS":
		return AbilityScores.WIS
	case "CHA":
		return AbilityScores.CHA
	default:
		return 0
	}
}

func calculateAbilityModifier(abilityScore int) int {
	switch {
	case abilityScore <= 1:
		return -5
	case abilityScore >= 30:
		return 10
	case abilityScore >= 10:
		return (abilityScore - 10) / 2
	default:
		return (abilityScore - 10) / 2
	}
}

func Check(skill string, AbilityScores godnd.AbilityScores) int {
	ability := godnd.SkillAbilityMap[skill]
	abilityScore := getAbilityScore(ability, AbilityScores)
	if abilityScore == 0 {
		return 0
	}
	modifier := calculateAbilityModifier(abilityScore)
	return D20() + modifier
}

func CheckAdv(skill string, AbilityScores godnd.AbilityScores) int {
	firstRoll := Check(skill, AbilityScores)
	secondRoll := Check(skill, AbilityScores)
	if firstRoll > secondRoll {
		return firstRoll
	}
	return secondRoll
}

func CheckDis(skill string, AbilityScores godnd.AbilityScores) int {
	firstRoll := Check(skill, AbilityScores)
	secondRoll := Check(skill, AbilityScores)
	if firstRoll < secondRoll {
		return firstRoll
	}
	return secondRoll
}

func CheckWithAdvantage(skill string, AbilityScores godnd.AbilityScores) int {
	return CheckAdv(skill, AbilityScores)
}

func CheckWithDisadvantage(skill string, AbilityScores godnd.AbilityScores) int {
	return CheckDis(skill, AbilityScores)
}
