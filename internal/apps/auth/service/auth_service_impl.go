package service

import (
	"context"
	"encoding/json"
	"rest_api/internal/apps/auth/load"

	"rest_api/internal/apps/register/model/domain"
	user "rest_api/internal/apps/register/service"

	"golang.org/x/oauth2"
)

type AuthServiceImpl struct {
	Config      *load.Config
	UserService user.UserService
}

func NewAuthService(config *load.Config, userService user.UserService) AuthService {
	return &AuthServiceImpl{
		Config:      config,
		UserService: userService,
	}
}

func (a *AuthServiceImpl) AuthCodeURL(state string) string {
	return a.Config.GoogleLoginConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (a *AuthServiceImpl) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return a.Config.GoogleLoginConfig.Exchange(ctx, code)
}

func (a *AuthServiceImpl) GetUserInfo(ctx context.Context, token *oauth2.Token) (*domain.User, error) {
	client := a.Config.GoogleLoginConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var user domain.User

	//var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil

}
