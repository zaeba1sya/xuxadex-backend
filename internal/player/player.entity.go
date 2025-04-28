package player

import (
	"gitlab.com/xyxa.gg/backend-mvp-main/internal/user"
)

type PlayerEntity struct {
	ID   string              `db:"id" json:"id"`
	User user.UserBaseEntity `db:"user" json:"user"`
}
