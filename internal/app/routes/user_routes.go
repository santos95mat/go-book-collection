package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/internal/handler"
	"github.com/santos95mat/go-book-collection/internal/middleware"
	"github.com/santos95mat/go-book-collection/internal/repository"
)

var (
	userRepository = repository.NewUserRepository()
	userHandler    = handler.NewUserHandler(userRepository)
)

func AddUserRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Post("/", userHandler.Create)
	user.Get("/", middleware.Auth, userHandler.GetMany)
	user.Get("/:id", middleware.Auth, userHandler.GetOne)
	user.Put("/:id", middleware.Auth, userHandler.Update)
	user.Delete("/:id", middleware.Auth, userHandler.Delete)
}
