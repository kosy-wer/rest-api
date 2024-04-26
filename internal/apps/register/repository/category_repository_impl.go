package repository

import (
	"context"
	"database/sql"
	"errors"
	"rest_api/internal/apps/register/helper"
	"rest_api/internal/apps/register/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
    SQL := "INSERT INTO category(name) VALUES ($1) RETURNING id"
    // Menggunakan $1 sebagai placeholder untuk parameter
    err := tx.QueryRowContext(ctx, SQL, category.Name).Scan(&category.Id)
    helper.PanicIfError(err)
    return category
}


func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
    SQL := "UPDATE category SET name = $1 WHERE id = $2"
    _, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
    helper.PanicIfError(err)
    return category
}


func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
    SQL := "DELETE FROM category WHERE id = $1"
    _, err := tx.ExecContext(ctx, SQL, category.Id)
    helper.PanicIfError(err)
}


func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
    SQL := "SELECT id, name FROM category WHERE id = $1"
    row := tx.QueryRowContext(ctx, SQL, categoryId)
    category := domain.Category{}
    err := row.Scan(&category.Id, &category.Name)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return category, errors.New("category is not found")
        }
        return category, err
    }
    return category, nil
}



func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
