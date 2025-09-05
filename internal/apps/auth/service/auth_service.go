package service

import (
	"context"

	//"rest_api/internal/apps/register/model/domain"
	"rest_api/internal/apps/register/model/web"

	"golang.org/x/oauth2"
)

type AuthService interface {
	AuthCodeURL(state string) string
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	GetUserInfo(ctx context.Context, token *oauth2.Token) (*web.UserLoginRequest, error)
	RegisterUser(ctx context.Context, token *oauth2.Token) (web.UserResponse, error)
}
