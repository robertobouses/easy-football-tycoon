package app

import (
	"errors"
)

var ErrPlayerNotFound = errors.New("el jugador no existe en la base de datos")
var ErrPlayerAlreadyInLineup = errors.New("El jugador ya está alineado")
