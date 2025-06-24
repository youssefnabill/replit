
package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"hexagonal-backend/internal/adapter/repository"
	"hexagonal-backend/internal/adapter/service"
	"hexagonal-backend/internal/app"
	"hexagonal-backend/internal/config"
)

// Server represents the HTTP server
type Server struct {
	config *config.Config
	router *gin.Engine
}

// NewServer creates a new HTTP server
func NewServer(cfg *config.Config) *Server {
	// Initialize repositories
	userRepo := repository.NewMemoryUserRepository()

	// Initialize services
	notificationService := service.NewEmailNotificationService()

	// Initialize application layer
	userApp := app.NewUserApp(userRepo, notificationService)

	// Initialize handlers
	userHandler := NewUserHandler(userApp)

	// Setup router
	router := gin.Default()
	
	// Setup routes
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.GET("", userHandler.ListUsers)
		}
	}

	return &Server{
		config: cfg,
		router: router,
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	addr := fmt.Sprintf("0.0.0.0:%s", s.config.Port)
	return s.router.Run(addr)
}
