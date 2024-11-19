package main

import (
	"subscription-tracker/internal/handlers"
	"subscription-tracker/internal/middleware"
	"subscription-tracker/internal/repository"
	"subscription-tracker/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	server := &Server{
		router: gin.Default(),
		db:     db,
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	// Initialize repositories
	userRepo := repository.NewUserRepository(s.db)

	// Initialize services
	authService := services.NewAuthService(userRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Public routes
	public := s.router.Group("/api/v1")
	{
		public.POST("/auth/register", authHandler.Register)
		public.POST("/auth/login", authHandler.Login)
	}

	// Protected routes
	protected := s.router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		// Add protected routes here later
	}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
