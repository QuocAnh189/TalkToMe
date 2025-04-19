package server

import (
	"fmt"
	"gochat/config"
	"gochat/pkg/logger"
	"gochat/pkg/mail"
	"gochat/pkg/storage"
	"gochat/pkg/token"
	"gochat/pkg/validation"

	"github.com/gin-gonic/gin"

	"gochat/internal/infrashstructrure/cache"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/internal/interfaces/api/routes"
	"gochat/internal/interfaces/websocket"
)

type Server struct {
	engine    *gin.Engine
	config    *config.Config
	hub       *websocket.Hub
	validator validation.Validation
	db        db.IDatabase
	cache     cache.IRedis
	token     token.IMarker
	storage   storage.IUploadService
	mailer    mail.IMailer
}

func NewHttpServer(
	hub *websocket.Hub,
	validator validation.Validation,
	db db.IDatabase,
	cache cache.IRedis,
	token token.IMarker,
	storage storage.IUploadService,
	mailer mail.IMailer,
) *Server {
	return &Server{
		engine:    gin.Default(),
		config:    config.GetConfig(),
		hub:       hub,
		validator: validator,
		db:        db,
		cache:     cache,
		token:     token,
		storage:   storage,
		mailer:    mailer,
	}
}

func (s *Server) Run() error {
	if s.config.Environment == config.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	go s.hub.Run()

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
	s.engine.GET("/ws", func(c *gin.Context) {
		websocket.WsHandler(s.hub, c)
	})

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
