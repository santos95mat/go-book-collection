package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/service"
	"github.com/santos95mat/go-book-collection/internal/util"
)

type BookController struct {
	bookService service.BookService
}

func (b BookController) Create(c *fiber.Ctx) error {
	var createBookDTO dto.BookInputDTO
	err := c.BodyParser(&createBookDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	createBookDTO, err = util.ValidBook(createBookDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	book, err := b.bookService.Create(createBookDTO)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (b BookController) GetMany(c *fiber.Ctx) error {
	q := c.Queries()
	search := q["search"]

	books, err := b.bookService.GetMany(search)

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
	var updateBookDTO dto.BookInputDTO
	id := c.Params("id")
	err := c.BodyParser(&updateBookDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	book, err := b.bookService.Update(id, updateBookDTO)

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
