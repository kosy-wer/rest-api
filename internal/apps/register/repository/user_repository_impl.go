package repository

import (
    "context"
    "database/sql"
    "errors"
    "rest_api/internal/apps/register/exception"
    "rest_api/internal/apps/register/helper"
    "rest_api/internal/apps/register/model/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
    return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
    SQL := `INSERT INTO users("first_name", "last_name", "email", "password", )
            VALUES ($1, $2, $3, $4, $5, $6) RETURNING "id"`
    err := tx.QueryRowContext(ctx, SQL,
        user.FirstName,
        user.LastName,
        user.Email,
        user.Password,
    ).Scan(&user.ID)
    if err != nil {
        panic(exception.NewNotFoundError(err.Error()))
    }
    return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
    SQL := `UPDATE users SET "first_name" = $1, "last_name" = $2, "password" = $3, "updated_at" = $4 WHERE "email" = $5`
    _, err := tx.ExecContext(ctx, SQL,
        user.FirstName,
        user.LastName,
        user.Password,
        user.UpdatedAt,
        user.Email,
    )
    helper.PanicIfError(err)
    return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) {
    SQL := `DELETE FROM users WHERE "email" = $1`
    _, err := tx.ExecContext(ctx, SQL, user.Email)
    helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, userEmail string) (domain.User, error) {
    SQL := `SELECT "id", "first_name", "last_name", "email", "password" 
            FROM users WHERE "email" = $1`
    row := tx.QueryRowContext(ctx, SQL, userEmail)

    user := domain.User{}
    err := row.Scan(
        &user.ID,
        &user.FirstName,
        &user.LastName,
        &user.Email,
        &user.Password,
    )
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return user, errors.New("user is not found")
        }
        return user, err
    }
    return user, nil
}

func (repository *UserRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, userName string) (domain.User, error) {
    SQL := `SELECT "id", "first_name", "last_name", "email", "password" 
            FROM users WHERE "first_name" = $1 OR "last_name" = $1`
    row := tx.QueryRowContext(ctx, SQL, userName)

    user := domain.User{}
    err := row.Scan(
        &user.ID,
        &user.FirstName,
        &user.LastName,
        &user.Email,
        &user.Password,
    )
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return user, errors.New("user is not found")
        }
        return user, err
    }
    return user, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
    SQL := `SELECT "id", "first_name", "last_name", "email", "password" FROM users`
    rows, err := tx.QueryContext(ctx, SQL)
    helper.PanicIfError(err)
    defer rows.Close()

    var users []domain.User
    for rows.Next() {
        user := domain.User{}
        err := rows.Scan(
            &user.ID,
            &user.FirstName,
            &user.LastName,
            &user.Email,
            &user.Password,
        )
        helper.PanicIfError(err)
        users = append(users, user)
    }
    return users
}

func (repository *UserRepositoryImpl) UserExist(ctx context.Context, tx *sql.Tx, userEmail string) (bool, error) {
    SQL := `SELECT EXISTS(SELECT 1 FROM users WHERE "email" = $1)`
    var exists bool
    err := tx.QueryRowContext(ctx, SQL, userEmail).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}

