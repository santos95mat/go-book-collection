package util

import (
	"errors"
	"net/mail"
	"strings"

	"github.com/santos95mat/go-book-collection/src/dto"
)

func ValidUser(data dto.UserBody) (dto.UserBody, error) {
	data.Name = strings.TrimSpace(data.Name)
	data.Number = strings.TrimSpace(data.Number)
	data.Email = strings.TrimSpace(data.Email)
	data.Password = strings.TrimSpace(data.Password)
	data.Role = strings.TrimSpace(data.Role)

	// Validação do nome
	if data.Name == "" {
		err := errors.New("error: Name is empty")
		return data, err
	}

	// Validação do número
	if data.Number == "" {
		err := errors.New("error: Number is empty")
		return data, err
	}

	// Validação do email
	_, err := mail.ParseAddress(data.Email)

	if err != nil {
		err = errors.New("error: invalid Email")
		return data, err
	}

	// Validação da senha
	if len(data.Password) <= 8 {
		err := errors.New("error: Password must have 8 characters")
		return data, err
	}

	// Validação da Role
	if data.Role == "admin" || data.Role == "normal" {
		return data, nil
	}

	err = errors.New("error: Role must be 'admin' or 'normal'")

	return data, err
}
