package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageAttachment struct {
	ID        string          `json:"id" gorm:"primaryKey"`
	MessageID string          `json:"message_id" gorm:"column:message_id;size:36;not null"`
	Type      string          `json:"type" gorm:"column:type;not null"`
	Filename  string          `json:"filename" gorm:"column:filename"`
	URL       string          `json:"url" gorm:"column:url;not null"`
	CreatedAt time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`

	// Relationships
	Message Message `json:"message" gorm:"foreignKey:MessageID"`
}

func (message_attachments *MessageAttachment) BeforeCreate(tx *gorm.DB) error {
	message_attachments.ID = uuid.New().String()

	return nil
}

func (MessageAttachment) TableName() string {
	return "message_attachments"
}
