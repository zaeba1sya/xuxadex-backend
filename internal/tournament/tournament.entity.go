package tournament

import (
	"time"

	"gitlab.com/xyxa.gg/backend-mvp-main/internal/game"
	"gitlab.com/xyxa.gg/backend-mvp-main/internal/player"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/repository"
)

const (
	StatusPreparation = "PREPARATION"
	StatusStarted     = "STARTED"
	StatusEnded       = "ENDED"
	StatusCanceled    = "CANCELED"
)

type (
	TournamentBaseEntity struct {
		ID             string                  `db:"id" json:"id"`
		Title          string                  `db:"title" json:"title"`
		EntranceFee    float64                 `db:"entrance_fee" json:"entrance_fee"`
		PreviewURL     repository.NullString   `db:"preview_url" json:"preview_url"`
		WinUpTo        float64                 `db:"win_up_to" json:"win_up_to"`
		FreeSlots      int                     `db:"free_slots" json:"free_slots"`
		Game           game.GameForMatchEntity `db:"game" json:"game"`
		StartTimestamp time.Time               `db:"start_timestamp" json:"start_timestamp"`
		Status         string                  `db:"status" json:"status"`
	}

	TournamentFullEntity struct {
		ID          string                  `db:"id" json:"id"`
		Title       string                  `db:"title" json:"title"`
		EntranceFee float64                 `db:"entrance_fee" json:"entrance_fee"`
		PreviewURL  repository.NullString   `db:"preview_url" json:"preview_url"`
		BannerURL   repository.NullString   `db:"banner_url" json:"banner_url"`
		PrizePool   float64                 `db:"prize_pool" json:"prize_pool"`
		FreeSlots   int                     `db:"free_slots" json:"free_slots"`
		Game        game.GameForMatchEntity `db:"game" json:"game"`
		Players     struct {
			Count int                   `db:"count" json:"count"`
			List  []player.PlayerEntity `db:"list" json:"list"`
		} `db:"players" json:"players"`
		StartTimestamp time.Time `db:"start_timestamp" json:"start_timestamp"`
		Status         string    `db:"status" json:"status"`
	}

	TournamentStatusEntity struct {
		ID   string `db:"id" json:"id"`
		Name string `db:"name" json:"name"`
	}
)
