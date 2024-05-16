package godnd

import (
	"fmt"
)

// Hero is a struct that represents a Dungeons and Dragons character.
// It has fields for the character's name, level, raw ability scores,
// maximum hit points, class, subclass, background, and other base data.
// This will represent the BASE STATS of a character, at level one.
type Hero struct {
	Name            string
	Level           int
	MaxHP           int
	Classes         []PlayerClass
	Background      string
	Race            string
	AlignmentSocial string
	AlignmentMoral  string
	STR             int
	DEX             int
	CON             int
	INT             int
	WIS             int
	CHA             int
}

// PlayerClass is a struct that represents a Dungeons and Dragons class.
// A hero can have multiple classes, and each class has a subclass and
// ajn independent level.
type PlayerClass struct {
	class *Class
	// subclass *Subclass
	level int
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

func main() {
	fmt.Println("Hello, World!")
}
