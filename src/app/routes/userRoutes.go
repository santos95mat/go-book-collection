package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/src/controller"
	"github.com/santos95mat/go-book-collection/src/middleware"
)

var userController controller.UserController

func AddUserRoutes(v1 fiber.Router) {
	user := v1.Group("/user")
	login := v1.Group("/login")

	user.Post("/", userController.Create)

	user.Get("/", middleware.Auth, userController.GetMany)

	user.Get("/:id", middleware.Auth, userController.GetOne)

	user.Put("/:id", middleware.Auth, userController.Update)

	user.Delete("/:id", middleware.Auth, userController.Delete)

	login.Post("/", userController.Login)

}
