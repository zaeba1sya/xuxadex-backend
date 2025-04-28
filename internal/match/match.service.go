package match

import (
	"context"
	"time"

	"gitlab.com/xyxa.gg/backend-mvp-main/db"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/repository"
)

type (
	MatchRepository interface {
		GetById(ctx context.Context, id string) (*MatchEntity, error)
		GetAll(ctx context.Context, queryOpts *repository.QueryOpts) (*[]MatchEntity, error)
		Create(ctx context.Context, data *MatchCreateDTO) (*MatchEntity, error)
		CreateQuickMatch(ctx context.Context, creatorID string, data *QuickMatchCreateDTO) (*MatchEntity, error)
		Update(ctx context.Context, data *MatchUpdateDTO) (*MatchEntity, error)
		Delete(ctx context.Context, id string) (bool, error)
		Count(ctx context.Context) (int, error)
	}

	MatchService struct {
		repository MatchRepository
	}
)

func NewMatchService(db *db.DBClient) *MatchService {
	return &MatchService{
		repository: NewMatchRepository(db),
	}
}

func (s *MatchService) GetById(ctx context.Context, id string) (*MatchEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repository.GetById(timeout, id)
}

func (s *MatchService) GetAll(ctx context.Context, queryOpts *repository.QueryOpts) (*[]MatchEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repository.GetAll(timeout, queryOpts)
}

func (s *MatchService) Create(ctx context.Context, data *MatchCreateDTO) (*MatchEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repository.Create(timeout, data)
}

func (s *MatchService) CreateQuickMatch(ctx context.Context, creatorID string, data *QuickMatchCreateDTO) (*MatchEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repository.CreateQuickMatch(timeout, creatorID, data)
}

func (s *MatchService) Update(ctx context.Context, data *MatchUpdateDTO) (*MatchEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repository.Update(timeout, data)
}

func (s *MatchService) Delete(ctx context.Context, id string) (bool, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repository.Delete(timeout, id)
}

func (s *MatchService) Count(ctx context.Context) (int, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repository.Count(timeout)
}
