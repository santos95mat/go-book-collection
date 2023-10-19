package main

import (
	"github.com/santos95mat/book-collection/src/app"
	"github.com/santos95mat/book-collection/src/initializer"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectDB()
}

func main() {
	app.Run()
}
