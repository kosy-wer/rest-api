package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

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

const version = "0.0.1"

type config struct {
	port int
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	// ========================
	// Load PORT
	// ========================
	portStr := os.Getenv("PORT")
	intPort, err := strconv.Atoi(portStr)
	if err != nil || portStr == "" {
		intPort = 4000 // fallback default sama kayak kode pertama
	}

	cfg := config{
		port: intPort,
	}

	// ========================
	// Init logger
	// ========================
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// ========================
	// Init DB, validator, repository, service, controller
	// ========================
	db, err := database.GetConnection()
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	/*
	   // Optional AuthController
	   configAuth, err := load.InitConfig()
	   if err != nil {
	       log.Fatalf("Failed to load config: %v", err)
	   }
	   authService := authService.NewAuthService(configAuth, userService)
	   authController := authController.NewAuthController(authService)
	*/

	/*
	   // Optional EmailService
	   emailCon, err := emailConfig.InitEmailConfig()
	   if err != nil {
	       log.Fatalf("Failed to load email config: %v", err)
	   }
	   emailService := emailService.NewEmailService(emailCon)

	   to := "protectorunmatched@gmail.com"
	   subject := "test subject"
	   body := "This you"
	   err = emailService.SendEmail(to, subject, body)
	   if err != nil {
	       log.Fatalf("Failed to send email: %v", err)
	   }
	*/

	// ========================
	// Router & Server
	// ========================
	app := &application{
		config: cfg,
		logger: logger,
	}

	// router := middleware.NewAuthMiddleware(api.NewRouter(userController))
	router := api.NewRouter(userController)

	srv := &http.Server{
	Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      router,
		IdleTimeout:  45 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("server started", "addr", srv.Addr)
	logger.Info("Using DATABASE_DSN", "dsn", os.Getenv("DATABASE_DSN"))

	// ========================
	// Start server
	// ========================
	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("server failed", "error", err)
		os.Exit(1)
	}
}

