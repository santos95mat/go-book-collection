package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID           `json:"id" gorm:"primaryKey"`
	Name      string              `json:"name" gorm:"not null"`
	Number    string              `json:"number" gorm:"not null;unique"`
	Email     string              `json:"email" gorm:"not null;unique"`
	Role      string              `json:"role" gorm:"not null"`
	Password  string              `json:"-" gorm:"not null"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Books     []*Book             `json:"books" gorm:"many2many:users_books"`
	Passwords []*RecoveryPassword `json:"passwords" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type RecoveryPassword struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	NewPassword string    `json:"new_password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
