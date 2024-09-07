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
	"github.com/robertobouses/easy-football-tycoon/http/team"
)

type Server struct {
	lineup    lineup.Handler
	team      team.Handler
	rival     rivalServer.Handler
	signings  signingsServer.Handler
	staff     staff.Handler
	calendary calendary.Handler
	analytics analytics.Handler
	resume    resume.Handler
	engine    *gin.Engine
}

func NewServer(
	lineup lineup.Handler,
	team team.Handler,
	rival rivalServer.Handler,
	signings signingsServer.Handler,
	staff staff.Handler,
	calendary calendary.Handler,
	analytics analytics.Handler,
	resume resume.Handler,

) Server {
	return Server{
		lineup:    lineup,
		team:      team,
		rival:     rival,
		signings:  signings,
		staff:     staff,
		calendary: calendary,
		analytics: analytics,
		resume:    resume,
		engine:    gin.Default(),
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

	staff := s.engine.Group("/staff")
	staff.POST("/create", s.staff.PostStaff)

	calendary := s.engine.Group("/calendary")
	calendary.GET("", s.calendary.GetCalendary)
	calendary.POST("/create", s.calendary.PostCalendary)

	resume := s.engine.Group("/resume")
	resume.GET("", s.resume.GetResume)
	resume.POST("/purchase-decision", s.resume.PostPurchaseDecision)
	resume.POST("/sale-decision", s.resume.PostSaleDecision)

	analytics := s.engine.Group("/analytics")
	analytics.GET("", s.analytics.GetAnalytics)
	analytics.POST("/create", s.analytics.PostAnalytics)

	log.Printf("running api at %s port\n", port)
	return s.engine.Run(fmt.Sprintf(":%s", port))
}
