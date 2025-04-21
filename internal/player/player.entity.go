package player

import (
	"github.com/xuxadex/backend-mvp-main/internal/user"
)

type PlayerEntity struct {
	ID   string              `db:"id" json:"id"`
	User user.UserBaseEntity `db:"user" json:"user"`
}
