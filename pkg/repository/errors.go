package repository

import (
	"errors"
	"fmt"
)

func NewNotFoundError(destination string) error {
	return errors.New(fmt.Sprintf("Database Error: %s not found with provided data", destination))
}

func NewInternalError(extra string) error {
	return errors.New(fmt.Sprintf("Database Error: %s", extra))
}
