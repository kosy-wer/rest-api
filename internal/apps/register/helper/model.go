package helper

import (
    "rest_api/internal/apps/register/model/domain"
    "rest_api/internal/apps/register/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
    return web.UserResponse{
        Name: user.Name,
        Email: user.Email,
    }
}

func ToUserResponses(users []domain.User) []web.UserResponse {
    var userResponses []web.UserResponse
    for _, user := range users {
        userResponses = append(userResponses, ToUserResponse(user))
    }
    return userResponses
}

