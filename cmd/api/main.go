package main

import (
	"github.com/santos95mat/go-book-collection/initializer/database"
	"github.com/santos95mat/go-book-collection/initializer/loadenv"
	"github.com/santos95mat/go-book-collection/internal/app"
)

func init() {
	loadenv.LoadEnvVariables()
	database.Connect()
	//migrate.Migrate()
}

func main() {
	app.Run()
}
