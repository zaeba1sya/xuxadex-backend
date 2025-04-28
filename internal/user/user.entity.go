package user

import (
	"time"

	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/repository"
)

type (
	UserBaseEntity struct {
		ID       string `db:"id" json:"id"`
		Nickname string `db:"nickname" json:"nickname"`
		Avatar   string `db:"avatar" json:"avatar"`
	}

	UserEntity struct {
		UserBaseEntity
		Wallet    string                `db:"wallet" json:"wallet"`
		SteamID   repository.NullString `db:"steam_id" json:"steam_id"`
		LastLogin time.Time             `db:"last_login" json:"last_login"`
	}
)
