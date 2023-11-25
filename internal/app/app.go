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

	app.Use(cors.New())
	getRoutes(app)

	err := app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		log.Fatalln(err)
	}
}

func getRoutes(app *fiber.App) {
	routes.AddStatusRoute(app)
	routes.AddBookRoutes(app)
	routes.AddUserRoutes(app)
	routes.AddAuthRoutes(app)
}
