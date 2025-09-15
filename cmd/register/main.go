package main

import (
	"log/slog"
	"net/http"
	"rest_api/api"
	"os"
	"time"
	"fmt"
	"strconv"
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
	//"rest_api/internal/apps/register/helper"
	"rest_api/internal/apps/register/repository"
	"rest_api/internal/apps/register/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)
const version = "0.0.1"

type config struct {
    port int
}


func main() {
    var cfg config

    // PORT dari env, default 4000
    port := os.Getenv("PORT")
    intPort, err := strconv.Atoi(port)
    if err != nil || intPort == 0 {
        intPort = 4000
    }
    cfg.port = intPort

    // Logger
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    db, err := database.GetConnection()
    if err != nil {
        logger.Error("failed to connect database", "error", err)
        os.Exit(1)
    }
    defer db.Close()

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
    //authcontroller is non active should commented this code main for local
    router := api.NewRouter(userController, authController)
    server := http.Server{
        Addr:    "localhost:3000",
        Handler: router,
        //Handler: middleware.NewAuthMiddleware(router),
    }
    */

    // Router aktif
    router := api.NewRouter(userController)

    // Server
    srv := &http.Server{
        Addr:         fmt.Sprintf(":%d", cfg.port),
        Handler:      router,
        IdleTimeout:  45 * time.Second,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
    }

    logger.Info("server started", "addr", srv.Addr, "version", version)

    if err := srv.ListenAndServe(); err != nil {
        logger.Error("server stopped", "error", err)
        os.Exit(1)
    }
}

