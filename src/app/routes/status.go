package routes

import "github.com/gofiber/fiber/v2"

func AddStatusRoute(v1 fiber.Router) {
	status := v1.Group("/status")

	status.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Book-collection API is running")
	})
}
