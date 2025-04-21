package tournament

import (
	"context"

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

func (r *tournamentRepositoryImpl) GetById(ctx context.Context, id string) (*TournamentBaseEntity, error) {
	return r.getBaseEntity(ctx, id)
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
	    LEFT JOIN tournament_uploading_entity tue on te.id = tue.tournament_id AND tue.type = (SELECT id FROM tournament_status_entity WHERE name = 'PREPARATION')
	    LEFT JOIN tournament_status_entity tse on te.status_id = tse.id
	    LEFT JOIN game_entity ge on te.game_id = ge.id
	    LEFT JOIN game_mode_entity gme on te.game_mode_id = gme.id
    `

	if err := r.db.GetClient().SelectContext(
		ctx,
		&dashboard.StartSoon,
		baseQuery+"WHERE te.start_timestamp > now() AND te.start_timestamp < now() + interval '10 minutes' LIMIT 16;",
	); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	if err := r.db.GetClient().SelectContext(
		ctx,
		&dashboard.Upcoming,
		baseQuery+"WHERE te.start_timestamp > now() LIMIT 8;",
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

	return dashboard, nil
}

func (r *tournamentRepositoryImpl) JoinTournament(ctx context.Context, data *TournamentJoinDTO) (*TournamentFullEntity, error) {
	// tournament := &TournamentFullEntity{}
	// query := ""

	return nil, nil
}

func (r *tournamentRepositoryImpl) GetAll(ctx context.Context, queryOpts *repository.QueryOpts) (*[]TournamentBaseEntity, error) {
	return nil, nil
}
