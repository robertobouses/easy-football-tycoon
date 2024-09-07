package app

import "github.com/google/uuid"

type Signings struct {
	SigningsId   uuid.UUID
	SigningsName string
	Position     string
	Age          int
	Fee          int
	Salary       int
	Technique    int
	Mental       int
	Physique     int
	InjuryDays   int
	Rarity       int
}
