package main

import (
	"github.com/santos95mat/go-book-collection/src/app"
	"github.com/santos95mat/go-book-collection/src/initializer"
	"github.com/santos95mat/go-book-collection/src/migrate"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectDB()
	migrate.Migrate()
}

func main() {
	app.Run()
}
