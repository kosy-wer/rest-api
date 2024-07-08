package main

import (
	"log"
	"net/http"
	"rest_api/api"
	authController "rest_api/internal/apps/auth/controller"
	"rest_api/internal/apps/auth/load"
	authService "rest_api/internal/apps/auth/service"
	"rest_api/internal/apps/database"

	//"rest_api/internal/apps/register/middleware"
	"rest_api/internal/apps/register/controller"
	"rest_api/internal/apps/register/helper"
	"rest_api/internal/apps/register/repository"
	"rest_api/internal/apps/register/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
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

	load.InitConfig()

	authService := authService.NewAuthService(load.AppConfig, userService)
	authController := authController.NewAuthController(authService)

	router := api.NewRouter(userController, authController)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
		//Handler: middleware.NewAuthMiddleware(router),
	}

	// Inisialisasi konfigurasi

	// Verifikasi bahwa konfigurasi telah diinisialisasi dengan benar
	log.Printf("Google Client ID: %s", load.AppConfig.GoogleClientID)
	log.Printf("Google Client Secret: %s", load.AppConfig.GoogleClientSecret)

	//
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
