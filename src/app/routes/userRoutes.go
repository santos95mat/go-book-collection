package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/src/controller"
)

var userController controller.UserController

func AddUserRoutes(v1 fiber.Router) {
	user := v1.Group("/user")

	user.Post("/", userController.Create)

	user.Get("/", userController.GetMany)

	user.Get("/:id", userController.GetOne)

	user.Put("/:id", userController.Update)

	user.Delete("/:id", userController.Delete)

}
