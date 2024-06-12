package api

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "rest_api/internal/apps/register/controller"
    "rest_api/internal/apps/register/exception"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
    router := httprouter.New()

    // User routes
    router.GET("/api/users", userController.FindAll)
    router.POST("/api/users", userController.Create)
    router.PUT("/api/users/:userId", userController.Update)
    router.GET("/api/users/:userId", userController.FindById)
    router.DELETE("/api/users/:userId", userController.Delete)

    // Serve Swagger UI
    fileServer := http.FileServer(http.Dir("/storage/emulated/0/rest_api/swagger-ui/dist"))
    router.Handler(http.MethodGet, "/swagger-ui/*filepath", http.StripPrefix("/swagger-ui", fileServer))

    // Serve the swagger.json file
    router.GET("/swagger/swagger.json", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        http.ServeFile(w, r, "/storage/emulated/0/rest_api/swagger-ui/dist/swagger.json")
    })

    router.PanicHandler = exception.ErrorHandler

    return router
}

