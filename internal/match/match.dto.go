package match

import (
	"errors"
	"time"
)

type (
	QuickMatchCreateDTO struct {
		StartTime  time.Time `json:"start_time"`
		MaxPlayers int       `json:"max_players"`
	}

	MatchCreateDTO struct {
	}

	JoinMatchDTO struct {
		UserID  string `json:"user_id"`
		MatchID string `json:"match_id"`
		TeamID  string `json:"team_id"`
	}

	MatchUpdateDTO struct {
	}
)

func (dto *QuickMatchCreateDTO) Validate() error {
	if dto.StartTime.IsZero() {
		return errors.New("start time is required")
	}
	if dto.MaxPlayers <= 0 {
		return errors.New("max players must be greater than 0")
	}
	return nil
}
