package models

import (
	models "github.com/RyanAMayers/godnd/models"
)

type barbarianClass struct {
	models.Class
	levels []barbarianLevel
}

type barbarianLevel struct {
	models.Level
	RageCount  int
	RageDamage int
}
