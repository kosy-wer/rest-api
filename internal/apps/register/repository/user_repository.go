package repository

import (
	"context"
	"database/sql"
	"rest_api/internal/apps/register/model/domain"
)

type UserRepository interface {
	Delete(ctx context.Context, tx *sql.Tx, user domain.Student)
	FindByEmail(ctx context.Context, tx *sql.Tx, userEmail string) (domain.Student, error)
	FindByName(ctx context.Context, tx *sql.Tx, userName string) (domain.Student, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.Student) domain.Student
	Save(ctx context.Context, tx *sql.Tx, user domain.Student) domain.Student
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Student

	UserExist(ctx context.Context, tx *sql.Tx, userEmail string) (bool, error)
}
