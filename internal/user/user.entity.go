package user

import (
	"time"

	"github.com/xuxadex/backend-mvp-main/pkg/repository"
)

type (
	UserEntity struct {
		ID        string                `db:"id" json:"id"`
		Nickname  string                `db:"nickname" json:"nickname"`
		Avatar    string                `db:"avatar" json:"avatar"`
		Wallet    string                `db:"wallet" json:"wallet"`
		SteamID   repository.NullString `db:"steam_id" json:"steam_id"`
		LastLogin time.Time             `db:"last_login" json:"last_login"`
	}
)
