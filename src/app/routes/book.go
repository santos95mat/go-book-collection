package routes

import "github.com/gofiber/fiber/v2"

func AddBookRoutes(v1 fiber.Router) {
	book := v1.Group("/book")

	book.Post("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Cadastrar um livro na base de dados",
		})
	})

	book.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Encontrar livros na base de dados baseado nas Querys",
		})
	})

	book.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Encontrar um livro na base de dados",
		})
	})

	book.Put("/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Editar um livro na base de dados",
		})
	})

	book.Delete("/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Deletar um livro na base de dados",
		})
	})
}
