package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID            string          `gorm:"primaryKey"`
	Name          string          `gorm:"column:name;not null"`
	Description   string          `gorm:"column:description"`
	OwnerID       string          `gorm:"column:owner_id;size:36;not null"`
	LastMessageID *string         `gorm:"column:last_message_id;size:36"`
	CreatedAt     time.Time       `gorm:"column:created_at"`
	UpdatedAt     time.Time       `gorm:"column:updated_at"`
	DeletedAt     *gorm.DeletedAt `gorm:"column:deleted_at"`

	// Relationships
	Owner       User `gorm:"foreignKey:OwnerID"`
	LastMessage *Message
	Members     []*User `gorm:"many2many:group_users;"`
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
