package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AddBookRoutes(v1 fiber.Router) {
	book := v1.Group("/book")

	book.Post("/", authMiddleware.AuthADM, bookController.Create)
	book.Get("/", authMiddleware.Auth, bookController.GetMany)
	book.Get("/:id", authMiddleware.Auth, bookController.GetOne)
	book.Put("/:id", authMiddleware.AuthADM, bookController.Update)
	book.Delete("/:id", authMiddleware.AuthADM, bookController.Delete)
}
