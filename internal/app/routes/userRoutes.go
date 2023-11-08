package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AddUserRoutes(v1 fiber.Router) {
	user := v1.Group("/user")

	user.Post("/", userController.Create)
	user.Get("/", authMiddleware.Auth, userController.GetMany)
	user.Get("/:id", authMiddleware.Auth, userController.GetOne)
	user.Put("/:id", authMiddleware.Auth, userController.Update)
	user.Delete("/:id", authMiddleware.Auth, userController.Delete)
}
