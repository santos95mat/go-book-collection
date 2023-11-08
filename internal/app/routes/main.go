package routes

import (
	"github.com/santos95mat/go-book-collection/internal/controller"
	"github.com/santos95mat/go-book-collection/internal/middleware"
)

var (
	// Controllers
	authController   controller.AuthController
	bookController   controller.BookController
	statusController controller.StatusController
	userController   controller.UserController

	// Middleware
	authMiddleware middleware.AuthMiddleware
)
