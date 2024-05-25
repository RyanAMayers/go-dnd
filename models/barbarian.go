package models

type BarbarianClass struct {
	ClassDetails Class
	Levels       []BarbarianLevel
}

type BarbarianLevel struct {
	LevelDetails Level
	RageCount    int
	RageDamage   int
}

var Barbarian = BarbarianClass{
	ClassDetails: Class{
		Name:           "Barbarian",
		HitDie:         12,
		PrimaryAbility: "Strength",
		SavingThrows: []string{
			"Strength",
			"Constitution"},
		ArmorProficiencies: []string{
			"Light Armor",
			"Medium Armor",
			"Shields"},
		WeaponProficiencies: []string{
			"Simple Weapons",
			"Martial Weapons"},
		ToolProficiencies: []string{},
		SkillProficiencies: []string{
			"Animal Handling",
			"Athletics",
			"Intimidation",
			"Nature",
			"Perception",
			"Survival"},
		Spellcaster:   false,
		MulticlassMin: 13,
	},
	Levels: []BarbarianLevel{
		{
			LevelDetails: Level{
				LevelNumber:      1,
				ProficiencyBonus: 2,
				Features: []*ClassFeature{
					{
						Name:        "Rage",
						Description: "In battle, you fight with primal ferocity. On your turn, you can enter a rage as a bonus action. While raging, you gain the following benefits if you aren't wearing heavy armor: You have advantage on Strength checks and Strength saving throws. When you make a melee weapon attack using Strength, you gain a bonus to the damage roll that increases as you gain levels as a barbarian, as shown in the Rage Damage column of the Barbarian table. You have resistance to bludgeoning, piercing, and slashing damage.",
					},
					{
						Name:        "Unarmored Defense",
						Description: "While you are not wearing any armor, your Armor Class equals 10 + your Dexterity modifier + your Constitution modifier. You can use a shield and still gain this benefit.",
					},
				},
			},
			RageCount:  2,
			RageDamage: 2,
		},
	},
}
