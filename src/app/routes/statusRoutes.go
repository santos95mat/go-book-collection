package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/src/controller"
)

var statusController controller.StatusController

func AddStatusRoute(v1 fiber.Router) {
	status := v1.Group("/status")

	status.Get("/", statusController.Get)
	status.Get("/", statusController.Get)
}
