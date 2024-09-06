package app

import (
	"errors"
)

var ErrPlayerNotFound = errors.New("el jugador no existe en la base de datos")
var ErrPlayerAlreadyInLineup = errors.New("El jugador ya está alineado")
var ErrCalendaryDatesNoAvaliable = errors.New("no hay más fechas disponibles en el calendario")
var ErrTransferNotFound = errors.New("transfer fee is not set")
var ErrBalanceNotFound = errors.New("Error al obtener el balance anterior")
