package service

import (
    "context"
    "database/sql"
    "github.com/go-playground/validator/v10"
    "rest_api/internal/apps/register/exception"
    "rest_api/internal/apps/register/helper"
    "rest_api/internal/apps/register/model/domain"
    "rest_api/internal/apps/register/model/web"
    "rest_api/internal/apps/register/repository"
)

type UserServiceImpl struct {
    UserRepository repository.UserRepository
    DB             *sql.DB
    Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
    return &UserServiceImpl{
        UserRepository: userRepository,
        DB:             DB,
        Validate:       validate,
    }
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
    err := service.Validate.Struct(request)
    helper.PanicIfError(err)

    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    user := domain.User{
        Name:  request.Name,
        Email: request.Email,
    }

    user = service.UserRepository.Save(ctx, tx, user)

    return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
    err := service.Validate.Struct(request)
    helper.PanicIfError(err)

    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
    if err != nil {
        panic(exception.NewNotFoundError(err.Error()))
    }

    user.Name = request.Name

    user = service.UserRepository.Update(ctx, tx, user)

    return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userEmail string) {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    user, err := service.UserRepository.FindByEmail(ctx, tx, userEmail)
    if err != nil {
        panic(exception.NewNotFoundError(err.Error()))
    }

    service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindByEmail(ctx context.Context, userEmail string) web.UserResponse {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    user, err := service.UserRepository.FindByEmail(ctx, tx, userEmail)
    if err != nil {
        panic(exception.NewNotFoundError(err.Error()))
    }
    return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindByName(ctx context.Context, userName string) web.UserResponse {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    user, err := service.UserRepository.FindByName(ctx, tx, userName)
    if err != nil {
        panic(exception.NewNotFoundError(err.Error()))
    }

    return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
    tx, err := service.DB.Begin()
    helper.PanicIfError(err)
    defer helper.CommitOrRollback(tx)

    users := service.UserRepository.FindAll(ctx, tx)

    return helper.ToUserResponses(users)
}

