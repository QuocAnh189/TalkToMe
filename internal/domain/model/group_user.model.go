package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupUser struct {
	ID        string          `json:"id" gorm:"primaryKey"`
	UserID    string          `json:"user_id" gorm:"column:user_id;size:36;not null;uniqueIndex:idx_group_user"`
	GroupID   string          `json:"group_id" gorm:"column:group_id;size:36;not null;uniqueIndex:idx_group_user"`
	IsAdmin   bool            `json:"is_admin" gorm:"column:is_admin;default:false"`
	CreatedAt time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`

	// Relationships
	User  User  `json:"user" gorm:"foreignKey:UserID"`
	Group Group `json:"group" gorm:"foreignKey:GroupID"`
}

func (group_user *GroupUser) BeforeCreate(tx *gorm.DB) error {
	group_user.ID = uuid.New().String()

	return nil
}

func (GroupUser) TableName() string {
	return "group_users"
}
