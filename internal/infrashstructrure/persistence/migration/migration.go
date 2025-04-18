package migration

import (
	"gochat/internal/domain/model"
	"gochat/internal/infrashstructrure/persistence/db"
)

func Migrate(db *db.Database) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Friend{},
		&model.Notification{},
		&model.Group{},
		&model.GroupUser{},
		&model.Conversation{},
		&model.Message{},
		&model.MessageAttachment{},
	)
}
