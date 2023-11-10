package repository

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/santos95mat/go-book-collection/initializer/database"
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

type AuthRepository struct{}

func (AuthRepository) Login(data dto.UserLoginDTO) (entity.User, string, error) {
	var user entity.User

	err := database.DB.Preload(clause.Associations).First(&user, "email = ?", data.Email).Error

	if err != nil {
		return entity.User{}, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return entity.User{}, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	rokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return entity.User{}, "", err
	}

	return user, rokenString, nil
}
