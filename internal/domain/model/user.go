package model

import (
	"gochat/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string          `gorm:"primaryKey"`
	Name      string          `gorm:"column:name;not null"`
	Email     string          `gorm:"column:email;unique;not null"`
	AvatarURL string          `gorm:"column:avatar_url"`
	Password  string          `gorm:"column:password;not null"`
	Role      string          `gorm:"column:role;default:'user'"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	user.Password = utils.HashAndSalt([]byte(user.Password))

	return nil
}

func (User) TableName() string {
	return "users"
}
