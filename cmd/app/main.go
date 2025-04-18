package main

import (
	"gochat/config"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/internal/infrashstructrure/persistence/migration"
	"gochat/internal/interfaces/server"
	"gochat/pkg/logger"
	"sync"
)

var wg sync.WaitGroup

func main() {
	config := config.LoadConfig()
	logger.Initialize(config.Environment)

	database, err := db.NewDatabase(config.DatabaseURI)
	if err != nil {
		logger.Fatal("Database connection error:", err)

	}

	err = migration.Migrate(database)
	if err != nil {
		logger.Fatal("Database migration error:", err)
	}

	httpServer := server.NewHttpServer()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := httpServer.Run(); err != nil {
			logger.Fatal("Running HTTP server error:", err)
		}
	}()
	wg.Wait()
}
