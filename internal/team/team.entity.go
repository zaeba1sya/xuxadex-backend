package team

import "github.com/xuxadex/backend-mvp-main/internal/player"

type TeamEntity struct {
	ID         string                `db:"id" json:"id"`
	Name       string                `db:"name" json:"name"`
	MaxPlayers int                   `db:"max_players" json:"max_players"`
	MinPlayers int                   `db:"min_players" json:"min_players"`
	Players    []player.PlayerEntity `db:"players" json:"players"`
}
