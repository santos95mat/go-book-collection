package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/src/controller"
)

var authController controller.AuthController

func AddAuthRoutes(v1 fiber.Router) {
	login := v1.Group("/login")

	login.Post("/", authController.Login)
}
