package http

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/http/lineup"
	rivalServer "github.com/robertobouses/easy-football-tycoon/http/rival"
	"github.com/robertobouses/easy-football-tycoon/http/team"
)

type Server struct {
	lineup lineup.Handler
	team   team.Handler
	rival  rivalServer.Handler
	engine *gin.Engine
}

func NewServer(
	lineup lineup.Handler,
	team team.Handler,
	rival rivalServer.Handler,

) Server {
	return Server{
		lineup: lineup,
		team:   team,
		rival:  rival,
		engine: gin.Default(),
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
	rival.POST("/rival", s.rival.PostRival)

	log.Printf("running api at %s port\n", port)
	return s.engine.Run(fmt.Sprintf(":%s", port))
}
