package repository

import (
	"context"
	"database/sql"
	"errors"
	"rest_api/internal/apps/register/exception"
	"rest_api/internal/apps/register/helper"
	"rest_api/internal/apps/register/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.Student) domain.Student {
	SQL := `INSERT INTO students("Name", "Email") VALUES ($1, $2)`

	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.Student) domain.Student {
	SQL := `UPDATE students SET "Name" = $1 WHERE "Email" = $2`
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email)
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.Student) {
	SQL := `DELETE FROM students WHERE "Email"= $1`
	_, err := tx.ExecContext(ctx, SQL, user.Email)
	helper.PanicIfError(err)
}


func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, userEmail string) (domain.Student, error) {
	SQL := `SELECT "Name", "Email" FROM students WHERE "Email" = $1`
	row := tx.QueryRowContext(ctx, SQL, userEmail)
	user := domain.Student{}
	err := row.Scan(&user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user is not found")
		}
		return user, err
	}
	return user, nil
}
func (repository *UserRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, userName string) (domain.Student, error) {
    SQL := `SELECT "Name", "Email" FROM students WHERE "Name" = $1`
    row := tx.QueryRowContext(ctx, SQL, userName)
    user := domain.Student{}
    err := row.Scan(&user.Name, &user.Email)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return user, errors.New("user is not found")
        }
        return user, err
    }
    return user, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Student {
    SQL := `SELECT "Name", "Email" FROM students`
    rows, err := tx.QueryContext(ctx, SQL)
    helper.PanicIfError(err)
    defer rows.Close()

    var students []domain.Student
    for rows.Next() {
        user := domain.Student{}
        err := rows.Scan(&user.Name, &user.Email)
        helper.PanicIfError(err)
        students = append(students, user)
    }
    return students
}

func (repository *UserRepositoryImpl) UserExist(ctx context.Context, tx *sql.Tx, userEmail string) (bool, error) {
    SQL := `SELECT EXISTS(SELECT 1 FROM students WHERE "Email" = $1)`
    var exists bool
    err := tx.QueryRowContext(ctx, SQL, userEmail).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}

