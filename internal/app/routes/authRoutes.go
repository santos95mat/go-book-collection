package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AddAuthRoutes(v1 fiber.Router) {
	login := v1.Group("/login")

	login.Post("/", authController.Login)
}
