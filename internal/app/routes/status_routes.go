package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/internal/handler"
)

var statusHandler = handler.NewStatusHandler()

func AddStatusRoute(app *fiber.App) {
	status := app.Group("/status")

	status.Get("/", statusHandler.Get)
}
