package models

// Hero is a struct that represents a Dungeons and Dragons character.
// It has fields for the character's name, level, raw ability scores,
// maximum hit points, class, subclass, background, and other base data.
// This will represent the BASE STATS of a character, at level one.
type Hero struct {
	FullName        string
	ShortName       string
	TotalLevel      int
	MaxHP           int
	Classes         []PlayerClass
	Background      string
	Race            string
	AlignmentSocial string
	AlignmentMoral  string
	AbilityScores   AbilityScores
}

// AbilityScores is a struct that represents a Dungeons and Dragons character's
// raw ability scores. These scores are used to calculate the character's
// modifiers, which are used in various calculations throughout the game.
type AbilityScores struct {
	STR int
	DEX int
	CON int
	INT int
	WIS int
	CHA int
}
