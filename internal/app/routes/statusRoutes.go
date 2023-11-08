package routes

import (
	"github.com/gofiber/fiber/v2"
)

func AddStatusRoute(v1 fiber.Router) {
	status := v1.Group("/status")

	status.Get("/", statusController.Get)
}
