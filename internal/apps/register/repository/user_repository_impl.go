package repository

import (
    "context"
    "database/sql"
    "errors"
    "rest_api/internal/apps/register/helper"
    "rest_api/internal/apps/register/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
    return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
    SQL := "INSERT INTO users(name) VALUES ($1) RETURNING id"
    err := tx.QueryRowContext(ctx, SQL, user.Name).Scan(&user.Id)
    helper.PanicIfError(err)
    return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
    SQL := "UPDATE users SET name = $1 WHERE id = $2"
    _, err := tx.ExecContext(ctx, SQL, user.Name, user.Id)
    helper.PanicIfError(err)
    return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
    SQL := "DELETE FROM users WHERE id = $1"
    _, err := tx.ExecContext(ctx, SQL, user.Id)
    helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
    SQL := "SELECT id, name FROM users WHERE id = $1"
    row := tx.QueryRowContext(ctx, SQL, userId)
    user := domain.User{}
    err := row.Scan(&user.Id, &user.Name)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return user, errors.New("user is not found")
        }
        return user, err
    }
    return user, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
    SQL := "SELECT id, name FROM users"
    rows, err := tx.QueryContext(ctx, SQL)
    helper.PanicIfError(err)
    defer rows.Close()

    var users []domain.User
    for rows.Next() {
        user := domain.User{}
        err := rows.Scan(&user.Id, &user.Name)
        helper.PanicIfError(err)
        users = append(users, user)
    }
    return users
}

