package server

import (
	"subscription-tracker/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
	config *config.Config
}

func New(db *gorm.DB, cfg *config.Config) *Server {
	server := &Server{
		router: gin.Default(),
		db:     db,
		config: cfg,
	}

	server.setupRoutes()
	return server
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
