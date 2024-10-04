package main

import (
	"fmt"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
	"github.com/robertobouses/easy-football-tycoon/app/signings"
	"github.com/robertobouses/easy-football-tycoon/app/staff"
	"github.com/robertobouses/easy-football-tycoon/http"

	analyticsServer "github.com/robertobouses/easy-football-tycoon/http/analytics"
	calendaryServer "github.com/robertobouses/easy-football-tycoon/http/calendary"
	"github.com/robertobouses/easy-football-tycoon/http/lineup"
	resumeServer "github.com/robertobouses/easy-football-tycoon/http/resume"
	rivalServer "github.com/robertobouses/easy-football-tycoon/http/rival"
	signingsServer "github.com/robertobouses/easy-football-tycoon/http/signings"
	staffServer "github.com/robertobouses/easy-football-tycoon/http/staff"
	strategyServer "github.com/robertobouses/easy-football-tycoon/http/strategy"
	"github.com/robertobouses/easy-football-tycoon/http/team"
	teamStaffServer "github.com/robertobouses/easy-football-tycoon/http/team_staff"
	"github.com/robertobouses/easy-football-tycoon/internal"
	analyticsRepository "github.com/robertobouses/easy-football-tycoon/repository/analytics"
	bankRepository "github.com/robertobouses/easy-football-tycoon/repository/bank"
	calendaryRepository "github.com/robertobouses/easy-football-tycoon/repository/calendary"
	lineupRepository "github.com/robertobouses/easy-football-tycoon/repository/lineup"
	matchRepository "github.com/robertobouses/easy-football-tycoon/repository/match"
	rivalRepository "github.com/robertobouses/easy-football-tycoon/repository/rival"
	signingsRepository "github.com/robertobouses/easy-football-tycoon/repository/signings"
	staffRepository "github.com/robertobouses/easy-football-tycoon/repository/staff"
	strategyRepository "github.com/robertobouses/easy-football-tycoon/repository/strategy"
	teamRepository "github.com/robertobouses/easy-football-tycoon/repository/team"
	teamStaffRepository "github.com/robertobouses/easy-football-tycoon/repository/team_staff"
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

	signingsRepo, err := signingsRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}
	staffRepo, err := staffRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	teamStaffRepo, err := teamStaffRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	calendaryRepo, err := calendaryRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	analyticsRepo, err := analyticsRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	bankRepo, err := bankRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	matchRepo, err := matchRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	strategyRepo, err := strategyRepository.NewRepository(db)
	if err != nil {
		panic(err)
	}

	app := app.NewApp(
		lineupRepo,
		teamRepo,
		rivalRepo,
		signingsRepo,
		staffRepo,
		teamStaffRepo,
		calendaryRepo,
		analyticsRepo,
		bankRepo,
		matchRepo,
		strategyRepo)

	staffApp := staff.NewApp(staffRepo)

	signingsApp := signings.NewApp(signingsRepo)

	lineupHandler := lineup.NewHandler(app)

	teamHandler := team.NewHandler(app)

	rivalHandler := rivalServer.NewHandler(app)

	signingsHandler := signingsServer.NewHandler(&signingsApp)

	staffHandler := staffServer.NewHandler(staffApp)

	teamStaffHandler := teamStaffServer.NewHandler(app)

	calendaryHandler := calendaryServer.NewHandler(app)

	analyticsHandler := analyticsServer.NewHandler(app)

	resumeHandler := resumeServer.NewHandler(&app)

	strategyHandler := strategyServer.NewHandler(app)

	s := http.NewServer(
		lineupHandler,
		teamHandler,
		rivalHandler,
		signingsHandler,
		staffHandler,
		teamStaffHandler,
		calendaryHandler,
		analyticsHandler,
		resumeHandler,
		strategyHandler)
	s.Run("8080")
}
