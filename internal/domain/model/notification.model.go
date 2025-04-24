package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	ID        string          `json:"id" gorm:"primaryKey"`
	FromID    string          `json:"from_id" gorm:"column:from_id;size:36;not null"`
	ToID      string          `json:"to_id" gorm:"column:to_id;size:36;not null;index:idx_notification_to_read"`
	IsRead    bool            `json:"is_read" gorm:"column:is_read;default:false;index:idx_notification_to_read"`
	Type      string          `json:"type" gorm:"column:type;size:20;not null"`
	Content   string          `json:"content" gorm:"column:content;size:255;not null"`
	IsAccept  bool            `json:"is_accept" gorm:"column:is_accept;default:false"`
	IsMarked  bool            `json:"is_marked" gorm:"column:is_marked;default:false"`
	CreatedAt time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`

	// Relationships
	FromUser User `json:"from_user" gorm:"foreignKey:FromID"`
	ToUser   User `json:"to_user" gorm:"foreignKey:ToID"`
}

func (notification *Notification) BeforeCreate(tx *gorm.DB) error {
	notification.ID = uuid.New().String()

	return nil
}

func (Notification) TableName() string {
	return "notifications"
}
