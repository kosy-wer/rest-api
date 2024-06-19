package service

import (
    "context"
    "rest_api/internal/apps/register/model/web"
)

type UserService interface {
    Delete(ctx context.Context, userId int)
    FindById(ctx context.Context, userId int) web.UserResponse
    FindByName(ctx context.Context, userName string) web.UserResponse
    Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
    FindAll(ctx context.Context) []web.UserResponse
    Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
}

