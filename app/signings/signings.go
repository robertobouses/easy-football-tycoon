package signings

import "github.com/google/uuid"

type Signings struct {
	SigningsId  uuid.UUID
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
	Rarity      int
	Fitness     int
}
