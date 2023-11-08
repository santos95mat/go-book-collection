package service

import (
	"github.com/google/uuid"
	"github.com/santos95mat/go-book-collection/initializer/database"
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

type UserService struct{}

func (UserService) Create(data dto.UserInputDTO) (model.User, error) {
	id := uuid.New()
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		ID:       id,
		Name:     data.Name,
		Number:   data.Number,
		Email:    data.Email,
		Password: string(hash),
		Role:     data.Role,
	}

	err = database.DB.Create(&user).Error

	return user, err
}

func (UserService) GetMany(str string) ([]model.User, error) {
	var users []model.User
	search := "%" + str + "%"

	err := database.DB.Preload(clause.Associations).Where("name LIKE ?", search).
		Or("email LIKE ?", search).Find(&users).Error

	return users, err
}

func (UserService) GetOne(id string) (model.User, error) {
	var user model.User

	err := database.DB.Preload(clause.Associations).First(&user, "id = ?", id).Error

	return user, err
}

func (b UserService) Update(id string, data dto.UserInputDTO) (model.User, error) {
	var user model.User
	user, err := b.GetOne(id)

	if err != nil {
		return user, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		return user, err
	}

	err = database.DB.Model(&user).Updates(
		model.User{
			Name:     data.Name,
			Number:   data.Number,
			Email:    data.Email,
			Password: string(hash),
			Role:     data.Role,
		},
	).Error

	return user, err
}

func (b UserService) Delete(id string) error {
	_, err := b.GetOne(id)

	if err != nil {
		return err
	}

	err = database.DB.Delete(&model.User{}, "id = ?", id).Error

	return err
}
