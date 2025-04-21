package match

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"strings"

	"github.com/xuxadex/backend-mvp-main/db"
	"github.com/xuxadex/backend-mvp-main/pkg/random"
	"github.com/xuxadex/backend-mvp-main/pkg/repository"
)

var (
	randomNames = []string{
		"Fraggin' Fantastic",
		"Ctrl Alt Defeat",
		"Terrorists Are Us",
		"Pixelated Pistoleros",
		"Banana Peel Busters",
		"No Scope, No Problem",
		"Spray and Pray",
		"The Lagging Legends",
		"Smokescreen Surfers",
		"Headshot Enthusiasts",
	}
)

type matchRepositoryImpl struct {
	db *db.DBClient
}

func NewMatchRepository(db *db.DBClient) MatchRepository {
	return &matchRepositoryImpl{
		db: db,
	}
}

func (r *matchRepositoryImpl) GetById(ctx context.Context, id string) (*MatchEntity, error) {
	match := &MatchEntity{}
	query := `SELECT
	   me.id,
	   me.creator_id,
       te1.id          AS "team1.id",
       te1.name        AS "team1.name",
       te1.max_players AS "team1.max_players",
       te1.min_players AS "team1.min_players",
       me.team1_score,
       te2.id          AS "team2.id",
       te2.name        AS "team2.name",
       te2.max_players AS "team2.max_players",
       te2.min_players AS "team2.min_players",
       me.team2_score,
       me.start_time
	FROM match_entity me
        LEFT JOIN team_entity te1 ON te1.id = me.team1_id
        LEFT JOIN team_entity te2 ON te2.id = me.team2_id
    WHERE me.id = $1`

	err := r.db.GetClient().GetContext(ctx, match, query, id)
	if err != nil && err == sql.ErrNoRows {
		return nil, repository.NewNotFoundError("Match")
	} else if err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	return match, nil
}

func (r *matchRepositoryImpl) GetAll(ctx context.Context, queryOpts *repository.QueryOpts) (*[]MatchEntity, error) {
	sortableFields := []string{
		"id", "name", "max_players", "min_players",
	}

	matches := []MatchEntity{}
	query := `
	SELECT
	   me.id,
	   me.creator_id,
       te1.id          AS "team1.id",
       te1.name        AS "team1.name",
       te1.max_players AS "team1.max_players",
       te1.min_players AS "team1.min_players",
       me.team1_score,
       te2.id          AS "team2.id",
       te2.name        AS "team2.name",
       te2.max_players AS "team2.max_players",
       te2.min_players AS "team2.min_players",
       me.team2_score,
       me.start_time
	FROM match_entity me
        LEFT JOIN team_entity te1 ON te1.id = me.team1_id
        LEFT JOIN team_entity te2 ON te2.id = me.team2_id
    `

	if queryOpts.Filter != "" {
		query += fmt.Sprintf(" WHERE %s AND me.deleted_at IS NULL", queryOpts.Filter)
	} else {
		query += " WHERE me.deleted_at IS NULL"
	}

	if slices.Contains(sortableFields, queryOpts.Sort.Field) {
		query += fmt.Sprintf(" ORDER BY me.%s %s", queryOpts.Sort.Field, strings.ToUpper(queryOpts.Sort.Order))
	}

	query += fmt.Sprintf(" LIMIT %d", queryOpts.Limit)
	query += fmt.Sprintf(" OFFSET %d", (queryOpts.Page-1)*queryOpts.Limit)

	if err := r.db.GetClient().SelectContext(ctx, &matches, query); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	return &matches, nil
}

func (r *matchRepositoryImpl) Create(ctx context.Context, data *MatchCreateDTO) (*MatchEntity, error) {
	return nil, nil
}

func (r *matchRepositoryImpl) CreateQuickMatch(ctx context.Context, data *QuickMatchCreateDTO) (*MatchEntity, error) {
	tx, err := r.db.GetClient().BeginTx(ctx, nil)
	if err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	var matchID, team1ID, team2ID string

	queryTeam := "INSERT INTO team_entity (name, max_players) VALUES ($1, $2) RETURNING id"
	randomNameTeam1 := random.RandomStringFromGiven(randomNames)
	if err := tx.QueryRowContext(ctx, queryTeam, randomNameTeam1, data.MaxPlayers).Scan(&team1ID); err != nil {
		err := repository.RollbackTx(tx, err)
		return nil, repository.NewInternalError(err.Error())
	}

	randomNameTeam2 := random.RandomStringFromGiven(randomNames)
	if err := tx.QueryRowContext(ctx, queryTeam, randomNameTeam2, data.MaxPlayers).Scan(&team2ID); err != nil {
		err := repository.RollbackTx(tx, err)
		return nil, repository.NewInternalError(err.Error())
	}

	matchQuery := "INSERT INTO match_entity (creator_id, team1_id, team2_id, start_time) VALUES ($1, $2, $3, $4) RETURNING id"
	if err := tx.QueryRowContext(ctx, matchQuery, data.CreatorID, team1ID, team2ID, data.StartTime).Scan(&matchID); err != nil {
		err := repository.RollbackTx(tx, err)
		return nil, repository.NewInternalError(err.Error())
	}

	if err := tx.Commit(); err != nil {
		err := repository.RollbackTx(tx, err)
		return nil, repository.NewInternalError(err.Error())
	}

	return r.GetById(ctx, matchID)
}

func (r *matchRepositoryImpl) Update(ctx context.Context, data *MatchUpdateDTO) (*MatchEntity, error) {
	return nil, nil
}

func (r *matchRepositoryImpl) Delete(ctx context.Context, id string) (bool, error) {
	query := "UPDATE match_entity SET deleted_at = NOW() WHERE id = $1"
	if _, err := r.db.GetClient().ExecContext(ctx, query, id); err != nil {
		return false, repository.NewInternalError(err.Error())
	}

	return true, nil
}

func (r *matchRepositoryImpl) Count(ctx context.Context) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM match_entity WHERE deleted_at IS NULL"
	if err := r.db.GetClient().QueryRowContext(ctx, query).Scan(&count); err != nil {
		return 0, repository.NewInternalError(err.Error())
	}

	return count, nil
}
