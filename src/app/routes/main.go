package routes

import (
	"github.com/santos95mat/go-book-collection/src/controller"
	"github.com/santos95mat/go-book-collection/src/middleware"
)

// Controllers
var authController controller.AuthController
var bookController controller.BookController
var statusController controller.StatusController
var userController controller.UserController

// Middleware
var authMiddleware middleware.AuthMiddleware
