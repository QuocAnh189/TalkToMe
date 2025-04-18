package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	ID            string          `gorm:"primaryKey"`
	UserIDOne     string          `gorm:"column:user_id_one;size:36;not null;uniqueIndex:idx_conversation_users"`
	UserIDTwo     string          `gorm:"column:user_id_two;size:36;not null;uniqueIndex:idx_conversation_users"`
	LastMessageID *string         `gorm:"column:last_message_id;size:36;"`
	CreatedAt     time.Time       `gorm:"column:created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at"`
	DeletedAt     *gorm.DeletedAt `gorm:"column:deleted_at"`

	// Relationships
	UserOne     User      `gorm:"foreignKey:UserIDOne"`
	UserTwo     User      `gorm:"foreignKey:UserIDTwo"`
	LastMessage *Message  //`gorm:"foreignKey:LastMessageID"`
	Messages    []Message //`gorm:"foreignKey:ConversationID"`
}

func (conversation *Conversation) BeforeCreate(tx *gorm.DB) error {
	conversation.ID = uuid.New().String()

	return nil
}

func (Conversation) TableName() string {
	return "conversations"
}
