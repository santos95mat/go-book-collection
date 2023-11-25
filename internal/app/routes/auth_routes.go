package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/internal/handler"
	"github.com/santos95mat/go-book-collection/internal/repository"
)

var (
	authRepository = repository.NewAuthRepository()
	authHandler    = handler.NewAuthHandler(authRepository)
)

func AddAuthRoutes(app *fiber.App) {
	login := app.Group("/login")

	login.Post("/", authHandler.Login)
}
