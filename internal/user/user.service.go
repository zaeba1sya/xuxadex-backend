package user

import (
	"context"
	"time"

	"gitlab.com/xyxa.gg/backend-mvp-main/db"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/types"
)

type (
	UserRepository interface {
		GetById(ctx context.Context, id string) (*UserEntity, error)
		Authenticate(ctx context.Context, wallet types.Wallet) (*UserEntity, error)
	}

	UserService struct {
		log        logger.Logger
		repository UserRepository
	}
)

func NewUserService(db *db.DBClient, log logger.Logger) *UserService {
	return &UserService{
		log:        log,
		repository: NewUserRepository(db, log),
	}
}

func (s *UserService) GetById(ctx context.Context, id string) (*UserEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repository.GetById(timeout, id)
}

func (s *UserService) Authenticate(ctx context.Context, wallet types.Wallet) (*UserEntity, error) {
	timeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	return s.repository.Authenticate(timeout, wallet)
}
