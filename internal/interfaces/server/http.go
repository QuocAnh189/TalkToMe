package server

import (
	"fmt"
	"gochat/config"
	"gochat/pkg/logger"

	"github.com/gin-gonic/gin"

	"gochat/internal/interfaces/api/routes"
)

type Server struct {
	engine *gin.Engine
	config *config.Config
}

func NewHttpServer() *Server {
	return &Server{
		engine: gin.Default(),
		config: config.GetConfig(),
	}
}

func (s *Server) Run() error {
	if s.config.Environment == config.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := s.MapRoutes(); err != nil {
		logger.Fatalf("MapRoutes Error: %v", err)
	}

	if err := s.engine.Run(fmt.Sprintf(":%d", s.config.HttpPort)); err != nil {
		logger.Fatalf("Running HTTP server: %v", err)
	}

	logger.Info("HTTP server is listening on PORT: ", s.config.HttpPort)

	return nil
}

func (s *Server) MapRoutes() error {
	routesV1 := s.engine.Group("/api/v1")
	routes.NewAuthRoutes(routesV1)
	routes.NewUserRoutes(routesV1)
	routes.NewFriendRoutes(routesV1)
	routes.NewGroupRoutes(routesV1)
	routes.NewConversationRoutes(routesV1)
	routes.NewMessageRoutes(routesV1)
	routes.NewNotificationRoutes(routesV1)
	return nil
}
