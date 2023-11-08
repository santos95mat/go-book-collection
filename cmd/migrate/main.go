package migrate

import (
	"github.com/santos95mat/go-book-collection/initializer/database"
	"github.com/santos95mat/go-book-collection/initializer/loadenv"
	"github.com/santos95mat/go-book-collection/internal/model"
)

func init() {
	loadenv.LoadEnvVariables()
	database.Connect()
}

func Migrate() {
	database.DB.AutoMigrate(
		&model.Book{},
		&model.User{},
	)
}