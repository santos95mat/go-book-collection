package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app = fiber.New()

func Run() {
	app.Use(cors.New())

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("Book-collection API is running")
	})

	app.Listen(os.Getenv("PORT"))
}
