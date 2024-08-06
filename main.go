package main

import (
	"fmt"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
	"github.com/robertobouses/easy-football-tycoon/http"
	"github.com/robertobouses/easy-football-tycoon/http/lineup"
	prospectServer "github.com/robertobouses/easy-football-tycoon/http/prospect"
	rivalServer "github.com/robertobouses/easy-football-tycoon/http/rival"
	"github.com/robertobouses/easy-football-tycoon/http/team"
	"github.com/robertobouses/easy-football-tycoon/internal"
	lineupRepository "github.com/robertobouses/easy-football-tycoon/repository/lineup"
	prospectRepository "github.com/robertobouses/easy-football-tycoon/repository/prospect"
	rivalRepository "github.com/robertobouses/easy-football-tycoon/repository/rival"
	teamRepository "github.com/robertobouses/easy-football-tycoon/repository/team"
)

func main() {
	fmt.Println("hola")
	// db, err := internal.NewPostgres(internal.DBConfig{
	// 	User:     os.Getenv("DB_USER"),
	// 	Pass:     os.Getenv("DB_PASS"),
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	Database: os.Getenv("DB_DATABASE"),
	// })
	//TODO EN MIS VARIABLES DE ENTORNO TENGO USER ROBERTO EN DOCKER Y POSTGRES TENGO postgres
	//TODO NO DEBE ESTAR VISIBLE ESTA INFO
	db, err := internal.NewPostgres(internal.DBConfig{
		User:     "postgres",
		Pass:     "mysecretpassword",
		Host:     "localhost",
		Port:     "5432",
		Database: "easy_football_tycoon",
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

	rivalRepo, err := rivalRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	prospectRepo, err := prospectRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	app := app.NewApp(lineupRepo, teamRepo, rivalRepo, prospectRepo)

	lineupHandler := lineup.NewHandler(app)

	teamHandler := team.NewHandler(app)

	rivalHandler := rivalServer.NewHandler(app)

	prospectHandler := prospectServer.NewHandler(app)

	s := http.NewServer(lineupHandler, teamHandler, rivalHandler, prospectHandler)
	s.Run("8080")

}
