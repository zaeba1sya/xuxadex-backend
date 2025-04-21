package tournament

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/xuxadex/backend-mvp-main/db"
	"github.com/xuxadex/backend-mvp-main/pkg/logger"
	"github.com/xuxadex/backend-mvp-main/pkg/repository"
)

type tournamentRepositoryImpl struct {
	db  *db.DBClient
	log logger.Logger
}

func NewTournamentRepository(db *db.DBClient, log logger.Logger) TournamentRepository {
	return &tournamentRepositoryImpl{
		db:  db,
		log: log,
	}
}

func (r *tournamentRepositoryImpl) getBaseEntity(ctx context.Context, id string) (*TournamentBaseEntity, error) {
	entity := &TournamentBaseEntity{}
	query := `SELECT
    	te.id,
	    te.title,
	    te.entrance_fee,
	    tue.file_name AS "preview_url",
	    te.teams_count * te.entrance_fee AS "win_up_to",
	    0 AS "free_slots",
	    te.start_timestamp,
	    te.status
    FROM tournament_entity te
	    LEFT JOIN tournament_uploading_entity tue on te.id = tue.tournament_id AND tue.type = 'PREVIEW'
		WHERE te.id = $1`
	if err := r.db.GetClient().GetContext(ctx, entity, query, id); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}
	return entity, nil
}

func (r *tournamentRepositoryImpl) getFullEntity(ctx context.Context, id string) (*TournamentFullEntity, error) {
	entity := &TournamentFullEntity{}
	query := `SELECT
		te.id,
	    te.title,
	    te.entrance_fee,
		CONCAT('/static/uploadings/', tue1.file_uuid, '.', tue1.extension) AS "banner_url",
		CONCAT('/static/uploadings/', tue2.file_uuid, '.', tue2.extension) AS "preview_url",
	    0 AS "prize_pool",
	    0 AS "free_slots",
	    te.start_timestamp,
	    tse.name AS "status",
	    ge.id AS "game.id",
	    ge.name AS "game.name",
		ge.icon AS "game.icon",
	    gme.name AS "game.mode"
    FROM tournament_entity te
	    LEFT JOIN tournament_status_entity tse ON te.status_id = tse.id
		LEFT JOIN tournament_uploading_entity tue1 ON te.id = tue1.tournament_id AND tue1.type_id = (SELECT id FROM tournament_uploading_type_entity WHERE name = 'BANNER')
		LEFT JOIN tournament_uploading_entity tue2 ON te.id = tue2.tournament_id AND tue2.type_id = (SELECT id FROM tournament_uploading_type_entity WHERE name = 'PREVIEW')
	    LEFT JOIN game_entity ge ON te.game_id = ge.id
	    LEFT JOIN game_mode_entity gme ON te.game_mode_id = gme.id
		WHERE te.id = $1`
	if err := r.db.GetClient().GetContext(ctx, entity, query, id); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	return entity, nil
}

func (r *tournamentRepositoryImpl) GetById(ctx context.Context, id string) (*TournamentFullEntity, error) {
	return r.getFullEntity(ctx, id)
}

func (r *tournamentRepositoryImpl) CreateWithRelations(ctx context.Context, data *TournamentCreateDTO) (*TournamentBaseEntity, error) {
	var id string
	query := "INSERT INTO tournament_entity (title, description, creator_id, entrance_fee, teams_count, teams_size, start_timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	if err := r.db.GetClient().QueryRowContext(ctx, query, data.Title, data.Description, data.CreatorID, data.EntranceFee, data.TeamsCount, data.TeamSize, data.StartTimestamp).Scan(&id); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}
	return r.getBaseEntity(ctx, id)
}

func (r *tournamentRepositoryImpl) GetDashboard(ctx context.Context) (*TournamentDashboardDTO, error) {
	dashboard := &TournamentDashboardDTO{
		StartSoon: []TournamentBaseEntity{},
		Upcoming:  []TournamentBaseEntity{},
		Ongoing:   []TournamentBaseEntity{},
		Count:     0,
	}

	baseQuery := `SELECT
	    te.id,
	    te.title,
	    te.entrance_fee,
	    tue.file_name AS "preview_url",
	    te.teams_count * te.entrance_fee AS "win_up_to",
	    0 AS "free_slots",
	    te.start_timestamp,
	    tse.name AS "status",
	    ge.id AS "game.id",
	    ge.name AS "game.name",
		ge.icon AS "game.icon",
	    gme.name AS "game.mode"
	FROM tournament_entity te
	    LEFT JOIN tournament_uploading_entity tue on te.id = tue.tournament_id AND tue.type = (SELECT id FROM tournament_uploading_type_entity WHERE name = 'PREVIEW')
	    LEFT JOIN tournament_status_entity tse on te.status_id = tse.id
	    LEFT JOIN game_entity ge on te.game_id = ge.id
	    LEFT JOIN game_mode_entity gme on te.game_mode_id = gme.id
    `

	if err := r.db.GetClient().SelectContext(
		ctx,
		&dashboard.StartSoon,
		baseQuery+"WHERE te.start_timestamp > now() AND te.start_timestamp < now() + interval '10 minutes' AND te.status_id != (SELECT id FROM tournament_status_entity WHERE name = 'IN_PROGRESS') LIMIT 16;",
	); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	if err := r.db.GetClient().SelectContext(
		ctx,
		&dashboard.Upcoming,
		baseQuery+"WHERE te.start_timestamp > now() AND te.status_id != (SELECT id FROM tournament_status_entity WHERE name = 'IN_PROGRESS') LIMIT 8;",
	); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	if err := r.db.GetClient().SelectContext(
		ctx,
		&dashboard.Ongoing,
		baseQuery+"WHERE te.status_id = (SELECT id FROM tournament_status_entity WHERE name = 'IN_PROGRESS') LIMIT 16;",
	); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	c, err := r.Count(ctx)
	if err != nil {
		return nil, err
	}
	dashboard.Count = c

	return dashboard, nil
}

func (r *tournamentRepositoryImpl) JoinTournament(ctx context.Context, data *TournamentJoinDTO) (*TournamentFullEntity, error) {
	// tournament := &TournamentFullEntity{}
	// query := ""

	return nil, nil
}

func (r *tournamentRepositoryImpl) GetAll(ctx context.Context, queryOpts *repository.QueryOpts) (*[]TournamentBaseEntity, error) {
	sortableFields := []string{
		"id", "title", "created_at",
	}

	tournaments := []TournamentBaseEntity{}
	query := `
	SELECT
	    te.id,
	    te.title,
	    te.entrance_fee,
	    tue.file_name AS "preview_url",
	    te.teams_count * te.entrance_fee AS "win_up_to",
	    0 AS "free_slots",
	    te.start_timestamp,
	    tse.name AS "status",
	    ge.id AS "game.id",
	    ge.name AS "game.name",
		ge.icon AS "game.icon",
	    gme.name AS "game.mode"
	FROM tournament_entity te
	    LEFT JOIN tournament_uploading_entity tue on te.id = tue.tournament_id
	    LEFT JOIN tournament_status_entity tse on te.status_id = tse.id
	    LEFT JOIN game_entity ge on te.game_id = ge.id
	    LEFT JOIN game_mode_entity gme on te.game_mode_id = gme.id
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

	if err := r.db.GetClient().SelectContext(ctx, &tournaments, query); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	return &tournaments, nil
}

func (r *tournamentRepositoryImpl) GetStatuses(ctx context.Context) (*[]TournamentStatusEntity, error) {
	statuses := &[]TournamentStatusEntity{}

	query := "SELECT id, name FROM tournament_status_entity WHERE deleted_at IS NULL"

	if err := r.db.GetClient().Select(statuses, query); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	return statuses, nil
}

func (r *tournamentRepositoryImpl) RandomizeDates(ctx context.Context) error {
	query := "UPDATE tournament_entity SET start_timestamp = now() + interval '10 minutes' WHERE entrance_fee < 4;"
	query2 := "UPDATE tournament_entity SET start_timestamp = now() + interval '50 minutes' WHERE entrance_fee > 4;"

	if _, err := r.db.GetClient().ExecContext(ctx, query); err != nil {
		return repository.NewInternalError(err.Error())
	}

	if _, err := r.db.GetClient().ExecContext(ctx, query2); err != nil {
		return repository.NewInternalError(err.Error())
	}

	return nil
}

func (r *tournamentRepositoryImpl) Count(ctx context.Context) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM tournament_entity WHERE deleted_at IS NULL"

	if err := r.db.GetClient().QueryRowContext(ctx, query).Scan(&count); err != nil {
		return 0, repository.NewInternalError(err.Error())
	}

	return count, nil
}
