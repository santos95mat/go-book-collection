package dto

import "github.com/google/uuid"

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

type UserFavoriteBookDTO struct {
	BookID uuid.UUID `json:"book_id"`
	UserID uuid.UUID `json:"user_id"`
}
