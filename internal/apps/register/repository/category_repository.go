package repository

import (
	"context"
	"database/sql"
	"rest_api/internal/apps/register/model/domain"
)

type CategoryRepository interface {
	/*
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	*/
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
