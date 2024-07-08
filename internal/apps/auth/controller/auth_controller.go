// internal/apps/auth/controller/auth_controller.go
package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	GoogleLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GoogleCallback(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
