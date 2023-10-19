package service

import (
	"github.com/google/uuid"
	"github.com/santos95mat/book-collection/src/dto"
	"github.com/santos95mat/book-collection/src/initializer"
	"github.com/santos95mat/book-collection/src/model"
	"gorm.io/gorm/clause"
)

type BookService struct{}

func (BookService) Create(data dto.BodyBookDto) (model.Book, error) {
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

func (BookService) GetMany(data dto.SearchBookDto) ([]model.Book, error) {
	var books []model.Book

	queries := "name LIKE ? AND author LIKE ? AND gender LIKE ? AND year LIKE ?"
	nameP := "%" + data.Name + "%"
	authorP := "%" + data.Author + "%"
	genderP := "%" + data.Gender + "%"
	yearP := "%" + data.Year + "%"

	err := initializer.DB.Where(queries, nameP, authorP, genderP, yearP).Find(&books).Error

	return books, err
}

func (BookService) GetOne(id string) (model.Book, error) {
	var book model.Book

	err := initializer.DB.Preload(clause.Associations).First(&book, "id = ?", id).Error

	return book, err
}

func (BookService) Update(id string) {}

func (BookService) Delete(id string) {}
