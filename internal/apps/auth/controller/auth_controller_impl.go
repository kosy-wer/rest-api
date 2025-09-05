// internal/apps/auth/controller/auth_controller.go
package controller

import (
	"context"
	//"encoding/json"
	"net/http"
	"rest_api/internal/apps/auth/service"
	"rest_api/internal/apps/register/helper"
	"rest_api/internal/apps/register/model/web"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (a *AuthControllerImpl) GoogleLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	state := "randomstate"
	url := a.AuthService.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *AuthControllerImpl) GoogleCallback(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	state := r.URL.Query().Get("state")
	if state != "randomstate" {
		http.Error(w, "State did not match", http.StatusBadRequest)
		return
	}

	code := r.URL.Query().Get("code")

	token, err := a.AuthService.Exchange(context.Background(), code)
	helper.PanicIfError(err)

	userResponse, err := a.AuthService.RegisterUser(context.Background(), token)
	helper.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse}

	helper.WriteToResponseBody(w, webResponse)

	w.Write([]byte("User email: " + userResponse.Email + "\nUser name: " + userResponse.Name))
}
