package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID            string          `json:"id" gorm:"primaryKey"`
	Name          string          `json:"name" gorm:"column:name;not null"`
	Description   string          `json:"description" gorm:"column:description"`
	OwnerID       string          `json:"owner_id" gorm:"column:owner_id;size:36;not null"`
	LastMessageID *string         `json:"last_message_id" gorm:"column:last_message_id;size:36"`
	CreatedAt     time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time       `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt     *gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`

	// Relationships
	Owner       User `json:"owner" gorm:"foreignKey:OwnerID"`
	LastMessage *Message
	Members     []*User `json:"members" gorm:"many2many:group_users;"`
	Messages    []Message
	GroupUsers  []GroupUser
}

func (group *Group) BeforeCreate(tx *gorm.DB) error {
	group.ID = uuid.New().String()

	return nil
}

func (Group) TableName() string {
	return "groups"
}
