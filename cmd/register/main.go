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
    db, err := database.GetConnection()
    if err != nil {
        panic(err)
    }

    validate := validator.New()

    // Initialize User repository, service, and controller
    userRepository := repository.NewUserRepository()
    userService := service.NewUserService(userRepository, db, validate)
    userController := controller.NewUserController(userService)

    // Initialize router with userController
    router := api.NewRouter(userController)

    server := http.Server{
        Addr:    "localhost:3000",
        Handler: router,
        //Handler: middleware.NewAuthMiddleware(router),
    }

    err = server.ListenAndServe()
    helper.PanicIfError(err)
}

