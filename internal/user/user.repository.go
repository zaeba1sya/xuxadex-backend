package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/xuxadex/backend-mvp-main/db"
	"github.com/xuxadex/backend-mvp-main/pkg/logger"
	"github.com/xuxadex/backend-mvp-main/pkg/random"
	"github.com/xuxadex/backend-mvp-main/pkg/repository"
	"github.com/xuxadex/backend-mvp-main/pkg/types"
)

type userRepositoryImpl struct {
	db  *db.DBClient
	log logger.Logger
}

func NewUserRepository(db *db.DBClient, log logger.Logger) UserRepository {
	return &userRepositoryImpl{
		db:  db,
		log: log,
	}
}

func (r *userRepositoryImpl) isExistsByWallet(ctx context.Context, wallet types.Wallet) (bool, string, error) {
	var id string
	query := "SELECT id from user_entity where wallet = $1 AND deleted_at IS NULL"

	err := r.db.GetClient().GetContext(ctx, &id, query, wallet.String())
	if err != nil && err == sql.ErrNoRows {
		return false, "", nil
	} else if err != nil {
		return false, "", repository.NewInternalError(err.Error())
	}

	return true, id, nil
}

func (r *userRepositoryImpl) updateLoginTime(ctx context.Context, id string) (bool, error) {
	query := "UPDATE user_entity SET last_login = NOW() WHERE id = $1 AND deleted_at IS NULL"

	result, err := r.db.GetClient().ExecContext(ctx, query, id)
	if err != nil {
		return false, repository.NewInternalError(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, repository.NewInternalError(err.Error())
	}

	return rowsAffected > 0, nil
}

func (r *userRepositoryImpl) GetById(ctx context.Context, id string) (*UserEntity, error) {
	user := &UserEntity{}
	query := "SELECT id, nickname, avatar, wallet, steam_id, last_login from user_entity WHERE id = $1 AND deleted_at IS NULL"

	err := r.db.GetClient().GetContext(ctx, user, query, id)
	if err != nil && err == sql.ErrNoRows {
		return nil, repository.NewNotFoundError("User")
	} else if err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	return user, nil
}

func (r *userRepositoryImpl) Authenticate(ctx context.Context, wallet types.Wallet) (*UserEntity, error) {
	isExists, id, err := r.isExistsByWallet(ctx, wallet)
	if err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	if isExists {
		_, err = r.updateLoginTime(ctx, id)
		if err != nil {
			r.log.Error("Database Error: Failed to update login time")
		}
		return r.GetById(ctx, id)
	}

	return r.Create(ctx, wallet)
}

func (r *userRepositoryImpl) Create(ctx context.Context, wallet types.Wallet) (*UserEntity, error) {
	var id string
	query := "INSERT INTO user_entity (nickname, avatar, wallet) VALUES ($1, $2, $3) RETURNING id"

	nickname, err := random.RandomNickname(wallet.String())
	if err != nil {
		nickname = "Lihaha"
	}

	if err = r.db.GetClient().QueryRowContext(
		ctx,
		query,
		nickname, fmt.Sprintf("/mock/avatars/mock-avatar-%d.svg", random.RandomIntFromRange(1, 3)), wallet.String(),
	).Scan(&id); err != nil {
		return nil, repository.NewInternalError(err.Error())
	}

	return r.GetById(ctx, id)
}
