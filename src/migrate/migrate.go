package migrate

import (
	"github.com/santos95mat/go-book-collection/src/initializer"
	"github.com/santos95mat/go-book-collection/src/model"
)

func Migrate() {
	initializer.DB.AutoMigrate(
		&model.Book{},
	)
}
