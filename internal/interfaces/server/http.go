package server

import (
	"fmt"
	"gochat/config"
	"gochat/pkg/logger"

	"github.com/gin-gonic/gin"
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
	if err := s.engine.Run(fmt.Sprintf(":%d", s.config.HttpPort)); err != nil {
		logger.Fatalf("Running HTTP server: %v", err)
	}

	logger.Info("HTTP server is listening on PORT: ", s.config.HttpPort)

	return nil
}
