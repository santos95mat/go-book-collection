package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/src/controller"
	"github.com/santos95mat/go-book-collection/src/middleware"
)

var bookController controller.BookController

func AddBookRoutes(v1 fiber.Router) {
	book := v1.Group("/book", middleware.Auth)

	book.Post("/", bookController.Create)
	book.Get("/", bookController.GetMany)
	book.Get("/:id", bookController.GetOne)
	book.Put("/:id", bookController.Update)
	book.Delete("/:id", bookController.Delete)
}
