package main

import (
	"fmt"
	"log"
	"os"

	"github.com/robertobouses/easy-football-tycoon/app"
	"github.com/robertobouses/easy-football-tycoon/http"
	"github.com/robertobouses/easy-football-tycoon/http/lineup"
	"github.com/robertobouses/easy-football-tycoon/http/team"
	"github.com/robertobouses/easy-football-tycoon/internal"
	lineupRepository "github.com/robertobouses/easy-football-tycoon/repository/lineup"
	teamRepository "github.com/robertobouses/easy-football-tycoon/repository/team"
)

func main() {
	fmt.Println("hola")
	db, err := internal.NewPostgres(internal.DBConfig{
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	})
	if err != nil {
		log.Println(err)
		panic(err)
	}

	lineupRepo, err := lineupRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	teamRepo, err := teamRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	app := app.NewApp(lineupRepo, teamRepo)

	lineupHandler := lineup.NewHandler(app)

	teamHandler := team.NewHandler(app)

	s := http.NewServer(lineupHandler, teamHandler)
	s.Run("8080")

}
