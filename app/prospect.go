package app

import "github.com/google/uuid"

type Prospect struct {
	ProspectId   uuid.UUID
	ProspectName string
	Position     string
	Age          int
	Fee          int
	Salary       int
	Technique    int
	Mental       int
	Physique     int
	InjuryDays   int
	Job          string
	Rarity       int
}
