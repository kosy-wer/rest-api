package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	_ "github.com/lib/pq"
	"rest_api/api"
	"rest_api/internal/apps/database"
	"rest_api/internal/apps/register/controller"
	"rest_api/internal/apps/register/helper"
	"rest_api/internal/apps/register/repository"
	"rest_api/internal/apps/register/service"
	//"rest_api/internal/apps/register/middleware"
)

func main() {

	db , err := database.GetConnection()                          
        if err != nil {                                               
           panic(err)
                                                                              }
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := api.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
		//Handler: middleware.NewAuthMiddleware(router),
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
