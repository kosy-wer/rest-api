package main

import (
	"log"
	"net/http"
	"rest_api/api"
	/*
	authController "rest_api/internal/apps/auth/controller"
	"rest_api/internal/apps/auth/load"
	authService "rest_api/internal/apps/auth/service"
*/
	//emailConfig "rest_api/internal/apps/email/config"
	//emailService "rest_api/internal/apps/email/service"

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

	/*config, err := load.InitConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	authService := authService.NewAuthService(config, userService)
	authController := authController.NewAuthController(authService)
*/
	/*emailCon, err := emailConfig.InitEmailConfig()
	if err != nil {
		log.Fatalf("Failed to load email config: %v", err)
	}

	// Inisialisasi EmailService
	emailService := emailService.NewEmailService(emailCon)

	// Contoh penggunaan pengiriman email
	to := "protectorunmatched@gmail.com"
	subject := "test subject"
	body := "This you"

	err = emailService.SendEmail(to, subject, body)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}
	*/
	router := api.NewRouter(userController, /*authController*/)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
		//Handler: middleware.NewAuthMiddleware(router),
	}

	// Inisialisasi konfigurasi

	// Verifikasi bahwa konfigurasi telah diinisialisasi dengan benar
	log.Printf("start server")

	//
	err = server.ListenAndServe()
	helper.PanicIfError(err)
	log.Printf("start server")
}
