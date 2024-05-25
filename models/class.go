package models

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
	MulticlassMin       int
	// Subclasses          []Subclass
}

type Level struct {
	LevelNumber      int
	ProficiencyBonus int
	Features         []*ClassFeature
}

type ClassFeature struct {
	Name        string
	Description string
}

type PlayerClass struct {
	Class interface{}
	Level int
}
