package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupUser struct {
	ID        string          `gorm:"primaryKey"`
	UserID    string          `gorm:"column:user_id;size:36;not null"`
	GroupID   string          `gorm:"column:group_id;size:36;not null"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at"`

	// Relationships
	User  User  `gorm:"foreignKey:UserID"`
	Group Group `gorm:"foreignKey:GroupID"`
}

func (group_user *GroupUser) BeforeCreate(tx *gorm.DB) error {
	group_user.ID = uuid.New().String()

	return nil
}

func (GroupUser) TableName() string {
	return "group_users"
}
