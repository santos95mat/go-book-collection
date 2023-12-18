package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserInputDTO struct {
	Name     string `json:"name"`
	Number   string `json:"number"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserTokenInputDTO struct {
	UserID     string `json:"user_id"`
	Email      string `json:"email"`
	Validation string `json:"validation"`
}

type UserTokenOutputDTO struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type UserTokenValidateDTO struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}

type UserRecoverPasswordDTO struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type UserFavoriteBookDTO struct {
	BookID uuid.UUID `json:"book_id"`
	UserID uuid.UUID `json:"user_id"`
}
