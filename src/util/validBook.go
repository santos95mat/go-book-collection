package util

import (
	"errors"
	"strconv"
	"strings"

	"github.com/santos95mat/book-collection/src/dto"
)

func ValidBook(book dto.CreateBookDto) error {
	if strings.TrimSpace(book.Name) == "" {
		err := errors.New("error: Name is empty")
		return err
	}

	if strings.TrimSpace(book.Author) == "" {
		err := errors.New("error: Author is empty")
		return err
	}

	if strings.TrimSpace(book.Gender) == "" {
		err := errors.New("error: Gender is empty")
		return err
	}

	if len(strings.TrimSpace(book.Year)) != 4 {
		err := errors.New("error: Year must have 4 numbers")
		return err
	}

	_, e := strconv.Atoi(book.Year)

	if e != nil {
		err := errors.New("error: Year must have 4 numbers")
		return err
	}

	return nil
}
