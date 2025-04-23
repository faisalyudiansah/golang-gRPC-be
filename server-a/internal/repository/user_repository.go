package repository

import (
	"context"
	"database/sql"
	"server/internal/entity"
	"server/pkg/database/transactor"
)

type UserRepositoryInterface interface {
	FindByID(ctx context.Context, id int64) (*entity.User, error)
}

type UserRepositoryImplementation struct {
	db *sql.DB
}

func NewUserRepositoryImplementation(db *sql.DB) *UserRepositoryImplementation {
	return &UserRepositoryImplementation{
		db: db,
	}
}

func (rs *UserRepositoryImplementation) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	query := `
			select id, role, email, is_verified, is_oauth, created_at, updated_at, deleted_at 
			from users 
			where id = $1 and deleted_at is null
		`
	tx := transactor.ExtractTx(ctx)
	user := &entity.User{ID: id}
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, id).Scan(
			&user.ID,
			&user.Role,
			&user.Email,
			&user.IsVerified,
			&user.IsOAuth,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
	} else {
		err = rs.db.QueryRowContext(ctx, query, id).Scan(
			&user.ID,
			&user.Role,
			&user.Email,
			&user.IsVerified,
			&user.IsOAuth,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
