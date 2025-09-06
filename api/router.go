package api

import (
	//"net/http"
//	auth "rest_api/internal/apps/auth/controller"
	"rest_api/internal/apps/register/controller"
	"rest_api/internal/apps/register/exception"

	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"os"
)


func writeJSON(w http.ResponseWriter, status int, data interface{}) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    return json.NewEncoder(w).Encode(data)
}

func NewRouter(userController controller.UserController,/*authController auth.AuthController*/) *httprouter.Router {
	router := httprouter.New()

	// User routes
	router.GET("/api/users", userController.FindAll)
	router.POST("/api/users", userController.Create)
	router.PUT("/api/users/:userEmail", userController.Update)
	router.GET("/api/users/:userEmail", userController.FindByEmail)
	router.DELETE("/api/users/:userEmail", userController.Delete)
	router.GET("/api/env-check", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        dsn := os.Getenv("DATABASE_DSN")
        resp := map[string]string{
            "status": "ok",
            "dsn":    dsn,
        }

        _ = writeJSON(w, http.StatusOK, resp)
    })
/*
	router.POST("/api/login", userController.LoginHandler)
	 //Login and Logout routes
	   router.POST("/api/google/login", userController.LoginHandler)
	   router.POST("/api/google/callback", userController.LoginHandler)
	   router.POST("/api/logout", controller.LogoutHandler)
	*/

	// Serve Swagger UI
/*
fileServer := http.FileServer(http.Dir("/data/data/com.termux/files/home/go/swagger-ui/dist"))
router.Handler(http.MethodGet, "/swagger-ui/*filepath", http.StripPrefix("/swagger-ui", fileServer))

// Serve the swagger.json file
router.GET("/swagger/swagger.json", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    http.ServeFile(w, r, "/data/data/com.termux/files/home/go/swagger-ui/dist/swagger.json")
})

	router.GET("/google/login", authController.GoogleLogin)
	router.GET("/google_callback", authController.GoogleCallback)
*/
	router.PanicHandler = exception.ErrorHandler

	// Apply the middleware
	return router
}
