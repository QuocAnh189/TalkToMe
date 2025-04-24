package model

import (
	"gochat/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string          `json:"id" gorm:"primaryKey"`
	Name      string          `json:"name" gorm:"column:name;not null"`
	Email     string          `json:"email" gorm:"column:email;unique;not null"`
	AvatarURL string          `json:"avatar_url" gorm:"column:avatar_url"`
	Password  string          `json:"password" gorm:"column:password;not null"`
	Role      string          `json:"role" gorm:"column:role;default:'user'"`
	CreatedAt time.Time       `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	user.Password = utils.HashAndSalt([]byte(user.Password))

	return nil
}

func (User) TableName() string {
	return "users"
}
