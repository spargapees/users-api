package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	users_api "github.com/spargapees/users-api/config"
	"github.com/spargapees/users-api/handler"
	repository "github.com/spargapees/users-api/repository"
	"github.com/spargapees/users-api/server"
	service "github.com/spargapees/users-api/service"
	"os"
	"os/signal"
	"syscall"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logger.SetOutput(os.Stdout)

	if err := godotenv.Load(); err != nil {
		logger.WithError(err).Warn("No .env file found, proceeding with environment variables")
	} else {
		logger.Info("Loaded .env file successfully")
	}

	cfg, err := users_api.NewConfig()
	if err != nil {
		logger.WithError(err).Fatal("Failed to load configuration")
	}
	logger.Info("Configuration loaded successfully")

	db, err := sqlx.Open("postgres", cfg.DSN())
	if err != nil {
		logger.WithError(err).Fatal("Failed to connect to PostgreSQL")
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.WithError(err).Warn("Failed to close database connection")
		} else {
			logger.Info("Database connection closed successfully")
		}
	}()

	if err := db.Ping(); err != nil {
		logger.WithError(err).Fatal("Database connection test failed")
	}
	logger.Info("Successfully connected to PostgreSQL!")

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	srv := new(server.Server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	go func() {
		if err := srv.Run(port, handler.InitRoutes()); err != nil {
			logger.Fatalf("Error occurred while running HTTP server: %s", err.Error())
		}
	}()

	logger.Info("Users-API Started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("Server shutting down...")
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
