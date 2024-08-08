package app

import "github.com/google/uuid"

type Rival struct {
	TeamId    uuid.UUID
	TeamName  string
	Technique int
	Mental    int
	Physique  int
}
