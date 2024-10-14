package app

import "github.com/google/uuid"

type PlayerStats struct {
	PlayerId    uuid.UUID
	Appearances int
	Chances     int
	Assists     int
	Goals       int
	MVP         int
	Rating      float64
}
