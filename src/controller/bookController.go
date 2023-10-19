package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/book-collection/src/dto"
	"github.com/santos95mat/book-collection/src/util"
)

type BookController struct{}

func (BookController) Create(c *fiber.Ctx) error {
	var bookBody dto.CreateBookDto

	err := c.BodyParser(&bookBody)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	err = util.ValidBook(bookBody)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	return c.JSON(bookBody)
}

func (BookController) GetMany(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Encontrar livros na base de dados baseado nas Querys",
	})
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
