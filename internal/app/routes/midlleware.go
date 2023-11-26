package routes

import (
	"github.com/santos95mat/go-book-collection/internal/middleware"
)

var authMiddleware = middleware.NewAuthMiddleware()
