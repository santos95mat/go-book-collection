package interfaces

import (
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/entity"
)

type AuthRepositoryInterface interface {
	Login(data dto.UserLoginDTO) (entity.User, string, error)
	VerifyEmail(email string) error
}
