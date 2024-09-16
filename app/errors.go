package app

import (
	"errors"
)

var ErrPlayerNotFound = errors.New("el jugador no existe en la base de datos")
var ErrPlayerAlreadyInLineup = errors.New("El jugador ya está alineado")
var ErrCalendaryDatesNoAvaliable = errors.New("no hay más fechas disponibles en el calendario")
var ErrTransferNotFound = errors.New("transfer fee is not set")
var ErrBalanceNotFound = errors.New("Error al obtener el balance anterior")
var ErrPaidNotFound = errors.New("Error al obtener el pago")
var ErrInvalidUUID = errors.New("UUID inválido")
var ErrTeamsQualityNil = errors.New("la calidad total de ambos equipos es 0")
var ErrNoRivalsAvailable = errors.New("No hay equipos válidos")
var ErrProcessingMatch = errors.New("Error de procesamiento de partido: Error al enviar los datos del partido a la db")
var ErrLineupCompleted = errors.New("Ya hay once jugadores en la alineación")
var ErrLineupIncompleted = errors.New("La alineación debe incluir once jugadores")
