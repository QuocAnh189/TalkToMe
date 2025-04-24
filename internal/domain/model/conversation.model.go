package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	ID            string          `json:"id" gorm:"primaryKey"`
	UserIDOne     string          `json:"user_id_one" gorm:"column:user_id_one;size:36;not null;uniqueIndex:idx_conversation_users"`
	UserIDTwo     string          `json:"user_id_two" gorm:"column:user_id_two;size:36;not null;uniqueIndex:idx_conversation_users"`
	LastMessageID *string         `json:"last_message_id" gorm:"column:last_message_id;size:36;"`
	CreatedAt     time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time       `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt     *gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`

	// Relationships
	UserOne     User      `json:"user_one" gorm:"foreignKey:UserIDOne"`
	UserTwo     User      `json:"user_two" gorm:"foreignKey:UserIDTwo"`
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
