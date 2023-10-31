package service

import (
	"github.com/santos95mat/go-book-collection/src/dto"
	"github.com/santos95mat/go-book-collection/src/model"
)

type UserService struct{}

func (UserService) Create(data dto.UserBody) (model.User, error) {
	return model.User{}, nil
}

func (UserService) GetMany(str string) ([]model.User, error) {
	return []model.User{}, nil
}

func (UserService) GetOne(id string) (model.User, error) {
	return model.User{}, nil
}

func (UserService) Update(id string, data dto.UserBody) (model.User, error) {
	return model.User{}, nil
}

func (UserService) Delete(id string) error {
	return nil
}
