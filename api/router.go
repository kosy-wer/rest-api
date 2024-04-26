package api

import (
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

	router.PanicHandler = exception.ErrorHandler

	return router
}
