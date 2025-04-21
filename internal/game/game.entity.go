package game

type (
	GameEntity struct {
		ID   string `db:"id" json:"id"`
		Name string `db:"name" json:"name"`
		Icon string `db:"icon" json:"icon"`
	}

	GameForMatchEntity struct {
		GameEntity
		Mode string `db:"mode" json:"mode"`
	}
)
