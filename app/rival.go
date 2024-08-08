package app

import "github.com/google/uuid"

type Rival struct {
	RivalId   uuid.UUID
	RivalName string
	Technique int
	Mental    int
	Physique  int
}
