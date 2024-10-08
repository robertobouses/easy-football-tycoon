package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/http/analytics"
	"github.com/robertobouses/easy-football-tycoon/http/calendary"
	"github.com/robertobouses/easy-football-tycoon/http/lineup"
	"github.com/robertobouses/easy-football-tycoon/http/resume"
	rivalServer "github.com/robertobouses/easy-football-tycoon/http/rival"
	signingsServer "github.com/robertobouses/easy-football-tycoon/http/signings"
	"github.com/robertobouses/easy-football-tycoon/http/staff"
	"github.com/robertobouses/easy-football-tycoon/http/strategy"
	"github.com/robertobouses/easy-football-tycoon/http/team"
	"github.com/robertobouses/easy-football-tycoon/http/team_staff"
)

type Server struct {
	lineup     lineup.Handler
	team       team.Handler
	rival      rivalServer.Handler
	signings   signingsServer.Handler
	staff      staff.Handler
	team_staff team_staff.Handler
	calendary  calendary.Handler
	analytics  analytics.Handler
	resume     resume.Handler
	strategy   strategy.Handler

	engine *gin.Engine
}

func NewServer(
	lineup lineup.Handler,
	team team.Handler,
	rival rivalServer.Handler,
	signings signingsServer.Handler,
	staff staff.Handler,
	team_staff team_staff.Handler,
	calendary calendary.Handler,
	analytics analytics.Handler,
	resume resume.Handler,
	strategy strategy.Handler,

) Server {
	return Server{
		lineup:     lineup,
		team:       team,
		rival:      rival,
		signings:   signings,
		staff:      staff,
		team_staff: team_staff,
		calendary:  calendary,
		analytics:  analytics,
		resume:     resume,
		strategy:   strategy,
		engine:     gin.Default(),
	}
}

func (s *Server) Run(port string) error {
	s.engine.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET, PUT, POST, DELETE, PATCH, OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "X-Accept-Language"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	s.engine.GET("", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	team := s.engine.Group("/team")
	team.GET("", s.team.GetTeam)
	team.POST("/player", s.team.PostTeam)

	lineup := s.engine.Group("/lineup")
	lineup.GET("", s.lineup.GetLineup)
	lineup.POST("/player", s.lineup.PostLineup)

	rival := s.engine.Group("/rival")
	rival.GET("", s.rival.GetRival)
	rival.POST("/team", s.rival.PostRival)

	signings := s.engine.Group("/signings")
	signings.GET("", s.signings.GetSignings)
	signings.POST("/person", s.signings.PostSignings)
	signings.POST("/auto", s.signings.PostAutoPlayerGenerator)

	staff := s.engine.Group("/staff")
	staff.POST("/create", s.staff.PostStaff)
	staff.POST("/auto", s.staff.PostAutoStaffGenerator)

	team_staff := s.engine.Group("/team_staff")
	team_staff.POST("/create", s.team_staff.PostTeamStaff)
	team_staff.GET("/", s.team_staff.GetTeamStaff)

	calendary := s.engine.Group("/calendary")
	calendary.GET("", s.calendary.GetCalendary)
	calendary.POST("/create", s.calendary.PostCalendary)

	resume := s.engine.Group("/resume")
	resume.GET("", s.resume.GetResume)
	resume.POST("/player-signing-decision", s.resume.PostPlayerSigningDecision)
	resume.POST("/player-sale-decision", s.resume.PostPlayerSaleDecision)
	resume.POST("/staff-signing-decision", s.resume.PostStaffSigningDecision)
	resume.POST("/staff-sale-decision", s.resume.PostStaffSaleDecision)

	analytics := s.engine.Group("/analytics")
	analytics.GET("", s.analytics.GetAnalytics)
	analytics.POST("/create", s.analytics.PostAnalytics)

	strategy := s.engine.Group("/strategy")
	strategy.GET("", s.strategy.GetStrategy)
	strategy.POST("/create", s.strategy.PostStrategy)

	log.Printf("running api at %s port\n", port)
	return s.engine.Run(fmt.Sprintf(":%s", port))
}
