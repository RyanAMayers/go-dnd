package godnd

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

// PlayerClass is a struct that represents a Dungeons and Dragons class.
// A hero can have multiple classes, and each class has a subclass and
// an independent level.
type PlayerClass struct {
	Class Class
	// subclass *Subclass
	Level int
}

type Class struct {
	Name                string
	HitDie              int
	PrimaryAbility      string
	SavingThrows        []string
	ArmorProficiencies  []string
	WeaponProficiencies []string
	ToolProficiencies   []string
	SkillProficiencies  []string
	Spellcaster         bool
	// Subclasses          []Subclass
}
