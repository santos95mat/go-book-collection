package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/santos95mat/go-book-collection/internal/app/routes"
)

func Run() {
	app := fiber.New()
	v1 := app.Group("/v1")

	app.Use(cors.New())
	getRoutes(v1)

	err := app.Listen(os.Getenv("PORT"))

	if err != nil {
		log.Fatalln(err)
	}
}

func getRoutes(v1 fiber.Router) {
	routes.AddStatusRoute(v1)
	routes.AddBookRoutes(v1)
	routes.AddUserRoutes(v1)
	routes.AddAuthRoutes(v1)
}
