package interfaces

import (
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/entity"
)

type BookRepositoryInterface interface {
	Create(data dto.BookInputDTO) (entity.Book, error)
	GetMany(str string) ([]entity.Book, error)
	GetOne(id string) (entity.Book, error)
	Update(id string, data dto.BookInputDTO) (entity.Book, error)
	Delete(id string) error
}
