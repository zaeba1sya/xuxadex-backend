package game

import (
	"context"

	"gitlab.com/xyxa.gg/backend-mvp-main/db"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/repository"
)

type (
	GameRepository interface {
		GetByID(ctx context.Context, id string) (*GameEntity, error)
		GetAll(ctx context.Context, queryOpts *repository.QueryOpts) (*[]GameEntity, error)
		Create(ctx context.Context, data *GameCreateDTO) (*GameEntity, error)
		Delete(ctx context.Context, id string) error
	}

	GameService struct {
		repository GameRepository
	}
)

func NewGameService(db *db.DBClient) *GameService {
	return &GameService{
		repository: NewGameRepository(db),
	}
}

func (s *GameService) GetByID(ctx context.Context, id string) (*GameEntity, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *GameService) GetAll(ctx context.Context, queryOpts *repository.QueryOpts) (*[]GameEntity, error) {
	return s.repository.GetAll(ctx, queryOpts)
}

func (s *GameService) Create(ctx context.Context, data *GameCreateDTO) (*GameEntity, error) {
	return s.repository.Create(ctx, data)
}

func (s *GameService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
