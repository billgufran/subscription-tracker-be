package main

import (
	"subscription-tracker/internal/config"
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
	config *config.Config
}

func NewServer(db *gorm.DB, cfg *config.Config) *Server {
	server := &Server{
		router: gin.Default(),
		db:     db,
		config: cfg,
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	// Initialize repositories
	userRepo := repository.NewUserRepository(s.db)
	categoryRepo := repository.NewCategoryRepository(s.db)
	currencyRepo := repository.NewCurrencyRepository(s.db)
	billingCycleRepo := repository.NewBillingCycleRepository(s.db)
	subscriptionRepo := repository.NewSubscriptionRepository(s.db)
	paymentMethodRepo := repository.NewPaymentMethodRepository(s.db)

	// Initialize services with config
	authService := services.NewAuthService(userRepo, s.config)
	categoryService := services.NewCategoryService(categoryRepo)
	currencyService := services.NewCurrencyService(currencyRepo)
	billingCycleService := services.NewBillingCycleService(billingCycleRepo)
	subscriptionService := services.NewSubscriptionService(
		subscriptionRepo,
		categoryRepo,
		currencyRepo,
		billingCycleRepo,
		paymentMethodRepo,
	)
	paymentMethodService := services.NewPaymentMethodService(paymentMethodRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	currencyHandler := handlers.NewCurrencyHandler(currencyService)
	billingCycleHandler := handlers.NewBillingCycleHandler(billingCycleService)
	subscriptionHandler := handlers.NewSubscriptionHandler(subscriptionService)
	paymentMethodHandler := handlers.NewPaymentMethodHandler(paymentMethodService)

	// Public routes
	public := s.router.Group("/api/v1")
	{
		public.POST("/auth/register", authHandler.Register)
		public.POST("/auth/login", authHandler.Login)
		public.GET("/currencies", currencyHandler.GetAll)
	}

	// Protected routes
	protected := s.router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(s.config))
	{
		// Category routes
		categories := protected.Group("/categories")
		{
			categories.GET("/", categoryHandler.GetAll)
			categories.POST("/", categoryHandler.Create)
			categories.PUT("/:id", categoryHandler.Update)
			categories.DELETE("/:id", categoryHandler.Delete)
		}

		// Billing cycle routes
		billingCycles := protected.Group("/billing-cycles")
		{
			billingCycles.POST("/", billingCycleHandler.Create)
			billingCycles.GET("/", billingCycleHandler.GetAll)
			billingCycles.PUT("/:id", billingCycleHandler.Update)
			billingCycles.DELETE("/:id", billingCycleHandler.Delete)
		}

		// Payment method routes
		paymentMethods := protected.Group("/payment-methods")
		{
			paymentMethods.POST("/", paymentMethodHandler.Create)
			paymentMethods.GET("/", paymentMethodHandler.GetAll)
			paymentMethods.PUT("/:id", paymentMethodHandler.Update)
			paymentMethods.DELETE("/:id", paymentMethodHandler.Delete)
		}

		// Subscription routes
		subscriptions := protected.Group("/subscriptions")
		{
			subscriptions.POST("/", subscriptionHandler.Create)
			subscriptions.GET("/", subscriptionHandler.GetAll)
			subscriptions.GET("/:id", subscriptionHandler.GetByID)
			subscriptions.GET("/category/:categoryId", subscriptionHandler.GetByCategory)
			subscriptions.GET("/billing-cycle/:billingCycleId", subscriptionHandler.GetByBillingCycle)
			subscriptions.GET("/payment-method/:paymentMethodId", subscriptionHandler.GetByPaymentMethod)
			subscriptions.PUT("/:id", subscriptionHandler.Update)
			subscriptions.DELETE("/:id", subscriptionHandler.Delete)
		}
	}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
