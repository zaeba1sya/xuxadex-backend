package main

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/xyxa.gg/backend-mvp-main/config"
	"gitlab.com/xyxa.gg/backend-mvp-main/db"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/random"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.InitConfig()

	log := logger.NewApiLogger(cfg)
	log.InitLogger()

	db := db.New(cfg)
	db.MustConnect()

	tx, err := db.GetClient().BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	var id string
	queryUser := "INSERT INTO user_entity (nickname, avatar, wallet) VALUES ($1, $2, $3) RETURNING id"
	nickname, err := random.RandomNickname("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")
	if err != nil {
		nickname = "Lihaha"
	}

	if err = tx.QueryRowContext(
		ctx,
		queryUser,
		nickname, fmt.Sprintf("/static/mock/avatars/mock-avatar-%d.svg", random.RandomIntFromRange(1, 3)), "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
	).Scan(&id); err != nil {
		tx.Rollback()
		panic(err)
	}

	for i := range random.RandomIntFromRange(100, 500) {
		var tournamentID string
		queryTournament := "INSERT INTO tournament_entity (title, description, creator_id, entrance_fee, teams_count, teams_size, status_id, game_id, game_mode_id, start_timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"

		title := fmt.Sprintf("Tournament #%d", i)
		description := fmt.Sprintf("Description for tournament #%d", i)
		creatorID := id
		entranceFee := random.RandomIntFromRange(1, 1000)
		teamsCount := random.RandomIntFromRange(2, 10)
		teamsSize := random.RandomIntFromRange(2, 10)
		statusID := random.RandomStringFromGiven([]string{"2c60ea0a-f337-4752-8556-b630e6f6cd32", "d4114c4b-ebd8-4aca-84eb-dafa5ca474f7", "80261546-b3be-4420-af21-c6b529e4276e", "0a04a3ba-4126-4e81-b39c-d3eec7408cee"})
		gameID := "066f3ae5-8919-4a3e-8933-744767ddeef7"
		gameModeID := "ced861d0-6f1f-4c9c-8643-ec997f6d9fd3"
		startTimestamp := time.Now().Add(time.Duration(random.RandomIntFromRange(1, 1000)) * time.Minute)

		if err = tx.QueryRowContext(
			ctx,
			queryTournament,
			title, description, creatorID, entranceFee, teamsCount, teamsSize, statusID, gameID, gameModeID, startTimestamp,
		).Scan(&tournamentID); err != nil {
			tx.Rollback()
			panic(err)
		}

		queryUploading := "INSERT INTO tournament_uploading_entity (tournament_id, type_id, file_name, file_uuid, extension, size) VALUES ($1, $2, $3, $4, $5, $6)"
		if _, err = tx.ExecContext(
			ctx,
			queryUploading,
			tournamentID, "504e3166-b8db-407f-a7d1-5825413566a9", "asdf", random.RandomStringFromGiven([]string{"ead9b89f-6a73-4196-9e3e-bfe3cd4ea5ac", "a412ff17-8a57-4343-b6d3-118b8dcd8612", "cb70cd94-6add-4f0a-bff7-278ab7c13dcc", "df37b93c-21f7-4386-85ba-ce7a03d9716a"}), "svg", 0,
		); err != nil {
			tx.Rollback()
			panic(err)
		}

		if _, err = tx.ExecContext(
			ctx,
			queryUploading,
			tournamentID, "e6ffee14-ef02-40ee-9c20-af81ab9e3ca1", "asdf", "ced861d0-6f1f-4c9c-8643-ec997f6d9fd3", "svg", 0,
		); err != nil {
			tx.Rollback()
			panic(err)
		}
	}

	tx.Commit()
}
