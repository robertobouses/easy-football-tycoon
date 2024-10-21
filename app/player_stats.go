package app

import "github.com/google/uuid"

type PlayerStats struct {
	PlayerId    uuid.UUID
	Appearances int
	Blocks      int
	Saves       int
	AerialDuel  int
	KeyPass     int
	Assists     int
	Chances     int
	Goals       int
	MVP         int
	Rating      float64
}
