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
	bookBody    dto.BodyBookDto
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
	err := c.QueryParser(&b.bookSearch)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	books, err := b.bookService.GetMany(b.bookSearch)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

func (b BookController) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := b.bookService.GetOne(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (b BookController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	err := c.BodyParser(&b.bookBody)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	book, err := b.bookService.Update(id, b.bookBody)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (b BookController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := b.bookService.Delete(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Livro deletado com sucesso",
	})
}
