package app

import "github.com/google/uuid"

type Team struct {
	PlayerId   uuid.UUID
	PlayerName string
	Position   string
	Age        int
	Fee        int
	Salary     int
	Technique  int
	Mental     int
	Physique   int
	InjuryDays int
	Lined      bool
}
