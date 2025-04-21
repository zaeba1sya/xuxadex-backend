package game

import (
	"errors"
)

type (
	GameCreateDTO struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
	}
)

func (dto *GameCreateDTO) Validate() error {
	if dto.Name == "" {
		return errors.New("Validation error: name is required")
	}
	if dto.Icon == "" {
		return errors.New("Validation error: icon is required")
	}
	return nil
}
