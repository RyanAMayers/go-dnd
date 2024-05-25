package demos

import (
	models "github.com/RyanAMayers/godnd/models"
)

var Mitch models.Hero = models.Hero{
	FullName:   "The Mighty, Fearsome Mitch",
	ShortName:  "Mitch",
	TotalLevel: 1,
	MaxHP:      14,
	Classes: []models.PlayerClass{
		{
			Class: models.Barbarian,
			Level: 1,
		},
	},
	Background:      "Folk Hero",
	Race:            "Half-Orc",
	AlignmentSocial: "Neutral",
	AlignmentMoral:  "Neutral",
	AbilityScores: models.AbilityScores{
		STR: 20,
		DEX: 14,
		CON: 14,
		INT: 7,
		WIS: 12,
		CHA: 13,
	},
}
