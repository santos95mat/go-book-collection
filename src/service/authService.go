package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/santos95mat/go-book-collection/src/dto"
	"github.com/santos95mat/go-book-collection/src/initializer"
	"github.com/santos95mat/go-book-collection/src/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

type AuthService struct{}

func (AuthService) Login(data dto.UserLoginDTO) (model.User, string, error) {
	var user model.User

	err := initializer.DB.Preload(clause.Associations).First(&user, "email = ?", data.Email).Error

	if err != nil {
		return model.User{}, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return model.User{}, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	rokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return model.User{}, "", err
	}

	return user, rokenString, nil
}
