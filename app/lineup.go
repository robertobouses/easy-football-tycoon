package app

import "github.com/google/uuid"

type Lineup struct {
	PlayerId       uuid.UUID
	LastName       string
	Position       string
	Technique      int
	Mental         int
	Physique       int
	TotalTechnique int
	TotalMental    int
	TotalPhysique  int
}

//TODO ROBERTO TABLA LINEUP NO DUPLICAR DATOS Y PONER SOLO PLAYERID Y HACER JOIN EN LA TABLA SQL PARA OBTENER DATOS
