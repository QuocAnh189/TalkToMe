package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Friend struct {
	ID         string          `json:"id" gorm:"primaryKey"`
	InviterID  string          `json:"inviter_id" gorm:"column:inviter_id;size:36;not null;uniqueIndex:idx_friend_inviter_accepter"`
	AccepterID string          `json:"accepter_id" gorm:"column:accepter_id;size:36;not null;uniqueIndex:idx_friend_inviter_accepter"`
	CreatedAt  time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time       `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`

	// Relationships
	Inviter  User `json:"inviter" gorm:"foreignKey:InviterID"`
	Accepter User `json:"accepter" gorm:"foreignKey:AccepterID"`
}

func (friend *Friend) BeforeCreate(tx *gorm.DB) error {
	friend.ID = uuid.New().String()

	return nil
}

func (Friend) TableName() string {
	return "friends"
}
