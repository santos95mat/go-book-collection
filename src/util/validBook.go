package util

import (
	"errors"
	"strconv"
	"strings"

	"github.com/santos95mat/book-collection/src/dto"
)

func ValidBook(data dto.CreateBookDto) (dto.CreateBookDto, error) {
	data.Name = strings.TrimSpace(data.Name)
	data.Author = strings.TrimSpace(data.Author)
	data.Gender = strings.TrimSpace(data.Gender)
	data.Year = strings.TrimSpace(data.Year)

	if data.Name == "" {
		err := errors.New("error: Name is empty")
		return data, err
	}

	if data.Author == "" {
		err := errors.New("error: Author is empty")
		return data, err
	}

	if data.Gender == "" {
		err := errors.New("error: Gender is empty")
		return data, err
	}

	if len(data.Year) != 4 {
		err := errors.New("error: Year must have 4 numbers")
		return data, err
	}

	_, e := strconv.Atoi(data.Year)

	if e != nil {
		err := errors.New("error: Year must have 4 numbers")
		return data, err
	}

	return data, nil
}
