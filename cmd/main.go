package main

import (
	"context"
	"errors"
	"github.com/Ahmad-mufied/iducate-community-service/config"
	"github.com/Ahmad-mufied/iducate-community-service/data"
	"github.com/Ahmad-mufied/iducate-community-service/server"
	"github.com/Ahmad-mufied/iducate-community-service/server/handler"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	postgresDb := config.InitDB()

	dbModel := data.New(postgresDb)
	validate := validator.New()
	handler.InitHandler(dbModel, validate)

	startAndGracefullyStopServer(echo.New())

}

func startAndGracefullyStopServer(e *echo.Echo) {
	// Register routes
	server.Routes(e)

	env := config.Viper.GetString("APP_ENV")
	port := "8080"

	if env == "production" {
		log.Println("Running in production mode")
		port = config.Viper.GetString("PORT")
	} else {
		log.Println("Running in development mode")
	}

	log.Printf("Starting server on port %s...", port)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: e,
	}

	go func() {
		if err := e.StartServer(srv); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
