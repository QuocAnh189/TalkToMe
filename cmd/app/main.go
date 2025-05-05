package main

import (
	"gochat/pkg/mail"
	"gochat/pkg/storage/minio"
	jwtToken "gochat/pkg/token/jwt"
	"gochat/pkg/validation"

	"gochat/config"
	"gochat/internal/driver/server"
	"gochat/internal/driver/websocket"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/internal/infrashstructrure/persistence/migration"
	"gochat/pkg/logger"
	"sync"

	"gochat/internal/infrashstructrure/cache"
)

var wg sync.WaitGroup

func main() {
	cfg := config.LoadConfig()
	logger.Initialize(cfg.Environment)

	//database
	database, err := db.NewDatabase(cfg.DatabaseURI)
	if err != nil {
		logger.Fatal("Database connection error:", err)
	}

	err = migration.Migrate(database)
	if err != nil {
		logger.Fatal("Database migration error:", err)
	}

	// // Seed data one time
	// err = migration.SeedData(database)
	// if err != nil {
	// 	logger.Fatal("Database seeding error:", err)
	// }

	//validation
	validator := validation.New()

	//cache
	cache := cache.New(cache.Config{
		Address:  cfg.RedisURI,
		Password: cfg.RedisPassword,
		Database: cfg.RedisDB,
	})

	//token
	token, err := jwtToken.NewJTWMarker()
	if err != nil {
		logger.Fatal(err)
	}

	//minio
	minioClient, err := minio.NewMinioClient(
		cfg.MinioEndpoint,
		cfg.MinioAccessKey,
		cfg.MinioSecretKey,
		cfg.MinioBucket,
		cfg.MinioBaseurl,
		cfg.MinioUseSSL,
	)
	if err != nil {
		logger.Fatalf("Failed to connect to MinIO: %s", err)
	}

	//mailer
	mailer := mail.NewMailer(
		cfg.MailHost,
		cfg.MailPort,
		cfg.MailUser,
		cfg.MailPassword,
		cfg.MailFrom,
	)

	//websocket
	hub := websocket.NewHub()

	httpServer := server.NewHttpServer(
		hub,
		validator,
		database,
		cache,
		token,
		minioClient,
		mailer,
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := httpServer.Run(); err != nil {
			logger.Fatal("Running HTTP server error:", err)
		}
	}()
	wg.Wait()
}
