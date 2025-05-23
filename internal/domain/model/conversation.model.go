package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Conversation struct {
	ID              string          `json:"id" gorm:"primaryKey"`
	UserIDOne       string          `json:"user_id_one" gorm:"column:user_id_one;size:36;not null;uniqueIndex:idx_conversation_users"`
	UserIDTwo       string          `json:"user_id_two" gorm:"column:user_id_two;size:36;not null;uniqueIndex:idx_conversation_users"`
	LastMessageID   *string         `json:"last_message_id" gorm:"column:last_message_id;size:36;default:null"`
	Background      string          `json:"background" gorm:"column:background;size:255;default:null"`
	UserNicknameOne string          `json:"user_nickname_one" gorm:"column:user_nickname_one;size:255"`
	UserNicknameTwo string          `json:"user_nickname_two" gorm:"column:user_nickname_two;size:255"`
	CreatedAt       time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt       *gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`

	// Relationships
	UserOne     User `json:"user_one" gorm:"foreignKey:UserIDOne"`
	UserTwo     User `json:"user_two" gorm:"foreignKey:UserIDTwo"`
	LastMessage *Message
	Messages    []Message
}

func (conversation *Conversation) BeforeCreate(tx *gorm.DB) error {
	conversation.ID = uuid.New().String()

	return nil
}

func (Conversation) TableName() string {
	return "conversations"
}
