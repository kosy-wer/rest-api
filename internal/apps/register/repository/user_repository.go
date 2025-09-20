package repository

import (
	"context"
	"database/sql"
	"rest_api/internal/apps/register/model/domain"
)

type UserRepository interface {
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindByEmail(ctx context.Context, tx *sql.Tx, userEmail string) (domain.User, error)
	FindByName(ctx context.Context, tx *sql.Tx, userName string) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User

	UserExist(ctx context.Context, tx *sql.Tx, userEmail string) (bool, error)
}
