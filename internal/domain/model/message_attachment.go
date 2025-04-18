package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageAttachment struct {
	ID        string          `gorm:"primaryKey"`
	MessageID string          `gorm:"column:message_id;size:36;not null"`
	Type      string          `gorm:"column:type;not null"`
	Filename  string          `gorm:"column:filename"`
	URL       string          `gorm:"column:url;not null"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at"`

	// Relationships
	Message Message `gorm:"foreignKey:MessageID"`
}

func (message_attachments *MessageAttachment) BeforeCreate(tx *gorm.DB) error {
	message_attachments.ID = uuid.New().String()

	return nil
}

func (MessageAttachment) TableName() string {
	return "message_attachments"
}
