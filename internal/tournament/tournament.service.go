package tournament

import (
	"context"
	"time"

	"gitlab.com/xyxa.gg/backend-mvp-main/db"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/repository"
)

type (
	TournamentRepository interface {
		GetById(ctx context.Context, id string) (*TournamentFullEntity, error)
		GetAll(ctx context.Context, queryOpts *repository.QueryOpts) (*[]TournamentBaseEntity, error)
		GetDashboard(ctx context.Context) (*TournamentDashboardDTO, error)
		CreateWithRelations(ctx context.Context, data *TournamentCreateDTO) (*TournamentBaseEntity, error)
		JoinTournament(ctx context.Context, data *TournamentJoinDTO) (*TournamentFullEntity, error)
		GetStatuses(ctx context.Context) (*[]TournamentStatusEntity, error)
		RandomizeDates(ctx context.Context) error
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

func (s *TournamentService) GetById(ctx context.Context, id string) (*TournamentFullEntity, error) {
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

func (s *TournamentService) GetStatuses(ctx context.Context) (*[]TournamentStatusEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repo.GetStatuses(timeout)
}

func (s *TournamentService) RandomizeDates(ctx context.Context) error {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repo.RandomizeDates(timeout)
}
