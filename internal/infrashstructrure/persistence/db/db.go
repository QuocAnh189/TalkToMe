package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

const (
	DatabaseTimeout = time.Second * 5
)

type IDatabase interface {
	AutoMigrate(models ...any) error
}

type Database struct {
	db *gorm.DB
}

func NewDatabase(uri string) (*Database, error) {
	database, err := gorm.Open(mysql.Open(uri), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})
	if err != nil {
		return nil, err
	}

	// Set up connection pool
	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)

	return &Database{
		db: database,
	}, nil
}

func (d *Database) AutoMigrate(models ...any) error {
	return d.db.AutoMigrate(models...)
}
