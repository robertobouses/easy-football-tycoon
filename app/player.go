package app

import "github.com/google/uuid"

type Player struct {
	PlayerId    uuid.UUID
	FirstName   string
	LastName    string
	Nationality string
	Position    string
	Age         int
	Fee         int
	Salary      int
	Technique   int
	Mental      int
	Physique    int
	InjuryDays  int
	Lined       bool
	Familiarity int
	Fitness     int
	Happiness   int
}
