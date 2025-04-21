package tournament

import (
	"errors"
	"time"
)

type (
	TournamentDashboardDTO struct {
		StartSoon []TournamentBaseEntity `json:"start_soon"`
		Upcoming  []TournamentBaseEntity `json:"upcoming"`
		Ongoing   []TournamentBaseEntity `json:"ongoing"`
	}

	TournamentJoinDTO struct {
		TournamentID string `json:"tournament_id"`
		TeamID       string `json:"team_id"`
		UserID       string `json:"user_id"`
	}

	TournamentCreateDTO struct {
		Title          string    `json:"title"`
		CreatorID      string    `json:"creator_id"`
		Description    string    `json:"description"`
		EntranceFee    float64   `json:"entrance_fee"`
		TeamsCount     int       `json:"teams_count"`
		TeamSize       int       `json:"team_size"`
		MatchDelay     int       `json:"match_delay"`
		StartTimestamp time.Time `json:"start_timestamp"`
	}
)

func (dto *TournamentDashboardDTO) Validate() error {
	return nil
}

func (dto *TournamentJoinDTO) Validate() error {
	if dto.UserID == "" {
		return errors.New("user id is required")
	}

	if dto.TeamID == "" {
		return errors.New("team id is required")
	}

	if dto.TournamentID == "" {
		return errors.New("tournament id is required")
	}
	return nil
}

func (dto *TournamentCreateDTO) Validate() error {
	if dto.Title == "" {
		return errors.New("title is required")
	}

	if dto.CreatorID == "" {
		return errors.New("creator id is required")
	}

	if dto.EntranceFee == 0 {
		return errors.New("entrance fee is required")
	}

	if dto.MatchDelay == 0 {
		return errors.New("match delay is required")
	}

	if dto.StartTimestamp.IsZero() {
		return errors.New("start timestamp is required")
	}
	return nil
}
