package repository

import (
	"github.com/google/uuid"
	"github.com/santos95mat/go-book-collection/initializer/database"
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (*userRepository) Create(data dto.UserInputDTO) (entity.User, error) {
	id := uuid.New()
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
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

func (*userRepository) GetMany(str string) ([]entity.User, error) {
	var users []entity.User
	search := "%" + str + "%"

	err := database.DB.Preload(clause.Associations).Where("name LIKE ?", search).
		Or("email LIKE ?", search).Find(&users).Error

	return users, err
}

func (*userRepository) GetOne(id string) (entity.User, error) {
	var user entity.User

	err := database.DB.Preload(clause.Associations).First(&user, "id = ?", id).Error

	return user, err
}

func (r *userRepository) Update(id string, data dto.UserInputDTO) (entity.User, error) {
	var user entity.User
	user, err := r.GetOne(id)

	if err != nil {
		return user, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		return user, err
	}

	err = database.DB.Model(&user).Updates(
		entity.User{
			Name:     data.Name,
			Number:   data.Number,
			Email:    data.Email,
			Password: string(hash),
			Role:     data.Role,
		},
	).Error

	return user, err
}

func (r *userRepository) Delete(id string) error {
	_, err := r.GetOne(id)

	if err != nil {
		return err
	}

	err = database.DB.Delete(&entity.User{}, "id = ?", id).Error

	return err
}

func (r *userRepository) FavoriteBook(data dto.UserFavoriteBookDTO) (*entity.User, error) {
	var book entity.Book
	err := database.DB.Preload(clause.Associations).First(&book, "id = ?", data.BookID).Error

	if err != nil {
		return nil, err
	}

	user, err := r.GetOne(data.UserID.String())

	if err != nil {
		return nil, err
	}

	database.DB.Model(&user).Association("Books").Append(&book)

	return &user, nil
}
