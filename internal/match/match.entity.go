package match

import (
	"time"

	"github.com/xuxadex/backend-mvp-main/internal/team"
)

type (
	MatchEntity struct {
		ID         string          `db:"id" json:"id"`
		CreatorID  string          `db:"creator_id" json:"creator"`
		Team1      team.TeamEntity `db:"team1" json:"team1"`
		Team1Score int             `db:"team1_score" json:"team1_score"`
		Team2      team.TeamEntity `db:"team2" json:"team2"`
		Team2Score int             `db:"team2_score" json:"team2_score"`
		StartTime  time.Time       `db:"start_time" json:"start_time"`
	}
)
