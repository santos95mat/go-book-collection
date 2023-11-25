package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/internal/handler"
	"github.com/santos95mat/go-book-collection/internal/repository"
)

var (
	bookRepository = repository.NewBookRepository()
	bookHandler    = handler.NewBookHandler(bookRepository)
)

func AddBookRoutes(app *fiber.App) {
	book := app.Group("/book")

	book.Post("/", authMiddleware.AuthADM, bookHandler.Create)
	book.Get("/", authMiddleware.Auth, bookHandler.GetMany)
	book.Get("/:id", authMiddleware.Auth, bookHandler.GetOne)
	book.Put("/:id", authMiddleware.AuthADM, bookHandler.Update)
	book.Delete("/:id", authMiddleware.AuthADM, bookHandler.Delete)
}
