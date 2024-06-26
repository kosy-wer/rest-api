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
    SQL := "INSERT INTO users(name, email) VALUES ($1, $2)"
    _, err := tx.ExecContext(ctx, SQL, user.Name, user.Email)
    helper.PanicIfError(err)
    return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
    SQL := "UPDATE users SET name = $1 WHERE email = $2"
    _, err := tx.ExecContext(ctx, SQL, user.Name, user.Email)
    helper.PanicIfError(err)
    return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
    SQL := "DELETE FROM users WHERE email = $1"
    _, err := tx.ExecContext(ctx, SQL, user.Email)
    helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, userEmail string) (domain.User, error) {
    SQL := "SELECT name, email FROM users WHERE email = $1"
    row := tx.QueryRowContext(ctx, SQL, userEmail)
    user := domain.User{}
    err := row.Scan(&user.Name, &user.Email)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return user, errors.New("user is not found")
        }
        return user, err
    }
    return user, nil
}

func (repository *UserRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, userName string) (domain.User, error) {
    SQL := "SELECT name, email FROM users WHERE name = $1"
    row := tx.QueryRowContext(ctx, SQL, userName)
    user := domain.User{}
    err := row.Scan(&user.Name, &user.Email)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return user, errors.New("user is not found")
        }
        return user, err
    }
    return user, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
    SQL := "SELECT name, email FROM users"
    rows, err := tx.QueryContext(ctx, SQL)
    helper.PanicIfError(err)
    defer rows.Close()

    var users []domain.User
    for rows.Next() {
        user := domain.User{}
        err := rows.Scan(&user.Name, &user.Email)
        helper.PanicIfError(err)
        users = append(users, user)
    }
    return users
}

