package game

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/xuxadex/backend-mvp-main/db"
	"github.com/xuxadex/backend-mvp-main/pkg/repository"
)

type (
	gameRepositoryImpl struct {
		db *db.DBClient
	}
)

func NewGameRepository(db *db.DBClient) GameRepository {
	return &gameRepositoryImpl{
		db: db,
	}
}

func (r *gameRepositoryImpl) GetByID(ctx context.Context, id string) (*GameEntity, error) {
	game := &GameEntity{}

	query := "SELECT id, name, icon FROM game_entity WHERE id = $1 AND deleted_at IS NULL"

	if err := r.db.GetClient().GetContext(ctx, game, query, id); err != nil {
		return nil, repository.NewNotFoundError("Game")
	}

	return game, nil
}

func (r *gameRepositoryImpl) GetAll(ctx context.Context, queryOpts *repository.QueryOpts) (*[]GameEntity, error) {
	sortableFields := []string{
		"id", "name",
	}

	games := []GameEntity{}
	query := `SELECT id, name, icon FROM game_entity`

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

	if err := r.db.GetClient().SelectContext(ctx, &games, query); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	return &games, nil
}

func (r *gameRepositoryImpl) Create(ctx context.Context, data *GameCreateDTO) (*GameEntity, error) {
	var id string

	query := "INSERT INTO game_entity(name, icon) VALUES ($1, $2) RETURNING id"

	if err := r.db.GetClient().QueryRowContext(ctx, query, data.Name, data.Icon).Scan(&id); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	return r.GetByID(ctx, id)
}

func (r *gameRepositoryImpl) Delete(ctx context.Context, id string) error {
	query := "UPDATE game_entity SET deleted_at = NOW() WHERE id = $1"

	if _, err := r.db.GetClient().ExecContext(ctx, query, id); err != nil {
		return repository.NewInternalError(err.Error())
	}

	return nil
}
