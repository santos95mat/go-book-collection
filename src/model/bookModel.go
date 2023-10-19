package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID
	Name      string    `json:"name" gorm:"primaryKey"`
	Author    string    `json:"author" gorm:"not null"`
	Gender    string    `json:"gender" gorm:"not null"`
	Year      string    `json:"year" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
