package interfaces

import (
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/entity"
)

type UserRepositoryInterface interface {
	Create(data dto.UserInputDTO) (entity.User, error)
	GetMany(str string) ([]entity.User, error)
	GetOne(id string) (entity.User, error)
	Update(id string, data dto.UserInputDTO) (entity.User, error)
	Delete(id string) error
	FavoriteBook(data dto.UserFavoriteBookDTO) (*entity.User, error)
}
