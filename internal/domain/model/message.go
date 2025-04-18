package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID             string          `gorm:"primaryKey"`
	Message        string          `gorm:"column:message"`
	GroupID        *string         `gorm:"column:group_id;index:idx_message_group_time"`
	SenderID       string          `gorm:"column:sender_id;not null"`
	ReceiverID     *string         `gorm:"column:receiver_id"`
	ConversationID *string         `gorm:"column:conversation_id;index:idx_message_conversation_time"`
	CreatedAt      time.Time       `gorm:"column:created_at;index:idx_message_group_time;index:idx_message_conversation_time"`
	UpdatedAt      time.Time       `gorm:"column:updated_at"`
	DeletedAt      *gorm.DeletedAt `gorm:"column:deleted_at"`

	// Relationships
	Sender       User                `gorm:"foreignKey:SenderID"`
	Receiver     *User               `gorm:"foreignKey:ReceiverID"`
	Group        *Group              `gorm:"foreignKey:GroupID"`
	Conversation *Conversation       `gorm:"foreignKey:ConversationID"`
	Attachments  []MessageAttachment `gorm:"foreignKey:MessageID"`
}

func (message *Message) BeforeCreate(tx *gorm.DB) error {
	message.ID = uuid.New().String()

	return nil
}

func (Message) TableName() string {
	return "messages"
}
