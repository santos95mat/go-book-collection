package repository

import (
	"github.com/google/uuid"
	"github.com/santos95mat/go-book-collection/initializer/database"
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/model"
	"gorm.io/gorm/clause"
)

type BookRepository struct {
}

func (BookRepository) Create(data dto.BookInputDTO) (model.Book, error) {
	id := uuid.New()

	book := model.Book{
		ID:     id,
		Name:   data.Name,
		Author: data.Author,
		Gender: data.Gender,
		Year:   data.Year,
	}

	err := database.DB.Create(&book).Error

	return book, err
}

func (BookRepository) GetMany(str string) ([]model.Book, error) {
	var books []model.Book

	search := "%" + str + "%"

	err := database.DB.Preload(clause.Associations).Where("name LIKE ?", search).
		Or("author LIKE ?", search).Or("gender LIKE ?", search).Or("year LIKE ?", search).
		Find(&books).Error

	return books, err
}

func (BookRepository) GetOne(id string) (model.Book, error) {
	var book model.Book

	err := database.DB.Preload(clause.Associations).First(&book, "id = ?", id).Error

	return book, err
}

func (b BookRepository) Update(id string, data dto.BookInputDTO) (model.Book, error) {
	var book model.Book
	book, err := b.GetOne(id)

	if err != nil {
		return book, err
	}

	err = database.DB.Model(&book).Updates(
		model.Book{
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

	err = database.DB.Delete(&model.Book{}, "id = ?", id).Error

	return err
}
