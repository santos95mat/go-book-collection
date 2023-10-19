package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/book-collection/src/dto"
	"github.com/santos95mat/book-collection/src/service"
	"github.com/santos95mat/book-collection/src/util"
)

type BookController struct {
	bookService service.BookService
	bookSearch  dto.SearchBookDto
	bookBody    dto.CreateBookDto
}

func (b BookController) Create(c *fiber.Ctx) error {
	err := c.BodyParser(&b.bookBody)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	data, err := util.ValidBook(b.bookBody)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	book, err := b.bookService.Create(data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (b BookController) GetMany(c *fiber.Ctx) error {
	q := c.Queries()
	b.bookSearch.Name = q["name"]
	b.bookSearch.Author = q["author"]
	b.bookSearch.Gender = q["gender"]
	b.bookSearch.Year = q["year"]

	books, err := b.bookService.GetMany(b.bookSearch)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

func (BookController) GetOne(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Encontrar um livro na base de dados",
	})
}

func (BookController) Update(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Editar um livro na base de dados",
	})
}

func (BookController) Delete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Deletar um livro na base de dados",
	})
}
