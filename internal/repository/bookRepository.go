package repository

import (
	"github.com/google/uuid"
	"github.com/santos95mat/go-book-collection/initializer/database"
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/entity"
	"gorm.io/gorm/clause"
)

type BookRepository struct {
}

func (BookRepository) Create(data dto.BookInputDTO) (entity.Book, error) {
	id := uuid.New()

	book := entity.Book{
		ID:     id,
		Name:   data.Name,
		Author: data.Author,
		Gender: data.Gender,
		Year:   data.Year,
	}

	err := database.DB.Create(&book).Error

	return book, err
}

func (BookRepository) GetMany(str string) ([]entity.Book, error) {
	var books []entity.Book

	search := "%" + str + "%"

	err := database.DB.Preload(clause.Associations).Where("name LIKE ?", search).
		Or("author LIKE ?", search).Or("gender LIKE ?", search).Or("year LIKE ?", search).
		Find(&books).Error

	return books, err
}

func (BookRepository) GetOne(id string) (entity.Book, error) {
	var book entity.Book

	err := database.DB.Preload(clause.Associations).First(&book, "id = ?", id).Error

	return book, err
}

func (b BookRepository) Update(id string, data dto.BookInputDTO) (entity.Book, error) {
	var book entity.Book
	book, err := b.GetOne(id)

	if err != nil {
		return book, err
	}

	err = database.DB.Model(&book).Updates(
		entity.Book{
			Name:   data.Name,
			Author: data.Author,
			Gender: data.Gender,
			Year:   data.Year,
		},
	).Error

	return book, err
}

func (b BookRepository) Delete(id string) error {
	_, err := b.GetOne(id)

	if err != nil {
		return err
	}

	err = database.DB.Delete(&entity.Book{}, "id = ?", id).Error

	return err
}
