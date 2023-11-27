package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/internal/handler"
	"github.com/santos95mat/go-book-collection/internal/middleware"
	"github.com/santos95mat/go-book-collection/internal/repository"
)

var (
	bookRepository = repository.NewBookRepository()
	bookHandler    = handler.NewBookHandler(bookRepository)
)

func AddBookRoutes(app *fiber.App) {
	book := app.Group("/book")

	book.Post("/", middleware.AuthADM, bookHandler.Create)
	book.Get("/", middleware.Auth, bookHandler.GetMany)
	book.Get("/:id", middleware.Auth, bookHandler.GetOne)
	book.Put("/:id", middleware.AuthADM, bookHandler.Update)
	book.Delete("/:id", middleware.AuthADM, bookHandler.Delete)
}
