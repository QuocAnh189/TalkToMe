package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	ID             string          `json:"id" gorm:"primaryKey"`
	Message        string          `json:"message" gorm:"column:message"`
	GroupID        *string         `json:"group_id" gorm:"column:group_id;size:36;index:idx_message_group_time"`
	SenderID       string          `json:"sender_id" gorm:"column:sender_id;size:36;not null"`
	ReceiverID     *string         `json:"receiver_id" gorm:"column:receiver_id;size:36;"`
	ConversationID *string         `json:"conversation_id" gorm:"column:conversation_id;size:36;index:idx_message_conversation_time"`
	CreatedAt      time.Time       `json:"created_at" gorm:"column:created_at;index:idx_message_group_time;index:idx_message_conversation_time"`
	UpdatedAt      time.Time       `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt      *gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`

	// Relationships
	Sender       User                `json:"sender" gorm:"foreignKey:SenderID"`
	Receiver     *User               `json:"receiver" gorm:"foreignKey:ReceiverID"`
	Group        *Group              `json:"group" gorm:"foreignKey:GroupID"`
	Conversation *Conversation       `json:"conversation" gorm:"foreignKey:ConversationID"`
	Attachments  []MessageAttachment `json:"attachments"`
}

func (message *Message) BeforeCreate(tx *gorm.DB) error {
	message.ID = uuid.New().String()

	return nil
}

func (Message) TableName() string {
	return "messages"
}
