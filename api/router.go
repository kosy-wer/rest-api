package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"rest_api/internal/apps/register/controller"
	"rest_api/internal/apps/register/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

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
