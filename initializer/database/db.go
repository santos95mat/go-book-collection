package database

import (
	"log"
	"os"

	"github.com/santos95mat/go-book-collection/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	DB.AutoMigrate(
		&entity.Book{},
		&entity.User{},
	)
}
