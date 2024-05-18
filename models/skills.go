package godnd

// SkillAbilityMap is a map that holds all the skills in D&D and their associated abilities.
// This map is used to determine which ability modifier to use when making a skill check.
// It also contains the abilities themselves, as well as saving throws. All of these are
// treated the same, as d20 tests.
var SkillAbilityMap = map[string]string{
	"Strength":     "STR",
	"StrengthSave": "STR",
	"Athletics":    "STR",

	"Dexterity":     "DEX",
	"DexteritySave": "DEX",
	"Acrobatics":    "DEX",
	"SleightOfHand": "DEX",
	"Stealth":       "DEX",

	"Constitution":     "CON",
	"ConstitutionSave": "CON",

	"Intelligence":     "INT",
	"IntelligenceSave": "INT",
	"Arcana":           "INT",
	"History":          "INT",
	"Investigation":    "INT",
	"Nature":           "INT",
	"Religion":         "INT",

	"Wisdom":         "WIS",
	"WisdomSave":     "WIS",
	"AnimalHandling": "WIS",
	"Insight":        "WIS",
	"Medicine":       "WIS",
	"Perception":     "WIS",
	"Survival":       "WIS",

	"Charisma":     "CHA",
	"CharismaSave": "CHA",
	"Deception":    "CHA",
	"Intimidation": "CHA",
	"Performance":  "CHA",
	"Persuasion":   "CHA",
}
