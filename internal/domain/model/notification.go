package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	ID        string          `gorm:"primaryKey"`
	FromID    string          `gorm:"column:from_id;size:36;not null"`
	ToID      string          `gorm:"column:to_id;size:36;not null;index:idx_notification_to_read"`
	IsRead    bool            `gorm:"column:is_read;default:false;index:idx_notification_to_read"`
	IsAccept  bool            `gorm:"column:is_accept;default:false"`
	IsMarked  bool            `gorm:"column:is_marked;default:false"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at"`

	// Relationships
	FromUser User `gorm:"foreignKey:FromID"`
	ToUser   User `gorm:"foreignKey:ToID"`
}

func (notification *Notification) BeforeCreate(tx *gorm.DB) error {
	notification.ID = uuid.New().String()

	return nil
}

func (Notification) TableName() string {
	return "notifications"
}
