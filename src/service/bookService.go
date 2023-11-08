package service

import (
	"github.com/google/uuid"
	"github.com/santos95mat/go-book-collection/src/dto"
	"github.com/santos95mat/go-book-collection/src/initializer"
	"github.com/santos95mat/go-book-collection/src/model"
	"gorm.io/gorm/clause"
)

type BookService struct{}

func (BookService) Create(data dto.BookInputDTO) (model.Book, error) {
	id := uuid.New()

	book := model.Book{
		ID:     id,
		Name:   data.Name,
		Author: data.Author,
		Gender: data.Gender,
		Year:   data.Year,
	}

	err := initializer.DB.Create(&book).Error

	return book, err
}

func (BookService) GetMany(str string) ([]model.Book, error) {
	var books []model.Book

	search := "%" + str + "%"

	err := initializer.DB.Preload(clause.Associations).Where("name LIKE ?", search).
		Or("author LIKE ?", search).Or("gender LIKE ?", search).Or("year LIKE ?", search).
		Find(&books).Error

	return books, err
}

func (BookService) GetOne(id string) (model.Book, error) {
	var book model.Book

	err := initializer.DB.Preload(clause.Associations).First(&book, "id = ?", id).Error

	return book, err
}

func (b BookService) Update(id string, data dto.BookInputDTO) (model.Book, error) {
	var book model.Book
	book, err := b.GetOne(id)

	if err != nil {
		return book, err
	}

	err = initializer.DB.Model(&book).Updates(
		model.Book{
			Name:   data.Name,
			Author: data.Author,
			Gender: data.Gender,
			Year:   data.Year,
		},
	).Error

	return book, err
}

func (b BookService) Delete(id string) error {
	_, err := b.GetOne(id)

	if err != nil {
		return err
	}

	err = initializer.DB.Delete(&model.Book{}, "id = ?", id).Error

	return err
}
