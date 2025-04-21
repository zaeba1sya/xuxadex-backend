package tournament

import (
	"context"
	"time"

	"github.com/xuxadex/backend-mvp-main/db"
	"github.com/xuxadex/backend-mvp-main/pkg/logger"
	"github.com/xuxadex/backend-mvp-main/pkg/repository"
)

type (
	TournamentRepository interface {
		GetById(ctx context.Context, id string) (*TournamentBaseEntity, error)
		GetAll(ctx context.Context, queryOpts *repository.QueryOpts) (*[]TournamentBaseEntity, error)
		GetDashboard(ctx context.Context) (*TournamentDashboardDTO, error)
		CreateWithRelations(ctx context.Context, data *TournamentCreateDTO) (*TournamentBaseEntity, error)
		JoinTournament(ctx context.Context, data *TournamentJoinDTO) (*TournamentFullEntity, error)
	}

	TournamentService struct {
		repo TournamentRepository
	}
)

func NewTournamentService(db *db.DBClient, log logger.Logger) *TournamentService {
	return &TournamentService{
		repo: NewTournamentRepository(db, log),
	}
}

func (s *TournamentService) GetById(ctx context.Context, id string) (*TournamentBaseEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repo.GetById(timeout, id)
}

func (s *TournamentService) GetAll(ctx context.Context, queryOpts *repository.QueryOpts) (*[]TournamentBaseEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repo.GetAll(timeout, queryOpts)
}

func (s *TournamentService) GetDashboard(ctx context.Context) (*TournamentDashboardDTO, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repo.GetDashboard(timeout)
}

func (s *TournamentService) CreateWithRelations(ctx context.Context, data *TournamentCreateDTO) (*TournamentBaseEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repo.CreateWithRelations(timeout, data)
}

func (s *TournamentService) JoinTournament(ctx context.Context, data *TournamentJoinDTO) (*TournamentFullEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repo.JoinTournament(timeout, data)
}
