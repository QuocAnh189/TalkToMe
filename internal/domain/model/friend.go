package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Friend struct {
	ID         string          `gorm:"primaryKey"`
	InviterID  string          `gorm:"column:inviter_id;not null;uniqueIndex:idx_friend_inviter_accepter"`
	AccepterID string          `gorm:"column:accepter_id;not null;uniqueIndex:idx_friend_inviter_accepter"`
	CreatedAt  time.Time       `gorm:"column:created_at"`
	UpdatedAt  time.Time       `gorm:"column:updated_at"`
	DeletedAt  *gorm.DeletedAt `gorm:"column:updated_at"`

	// Relationships
	Inviter  User `gorm:"foreignKey:InviterID"`
	Accepter User `gorm:"foreignKey:AccepterID"`
}

func (friend *Friend) BeforeCreate(tx *gorm.DB) error {
	friend.ID = uuid.New().String()

	return nil
}

func (Friend) TableName() string {
	return "friends"
}
