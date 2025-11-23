package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tamadamas/magic_stream/server/go/internal/config"
	"github.com/tamadamas/magic_stream/server/go/internal/handlers"
	"github.com/tamadamas/magic_stream/server/go/internal/middlewares"
	"github.com/tamadamas/magic_stream/server/go/internal/repositories"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	loadEnv()
	db := config.ConnectToDatabase()

	moviesHandler := handlers.NewMoviesHandler(repositories.NewMoviesRepository(db))

	r := gin.Default()
	r.Use(middlewares.RequestIDMiddleware())
	r.Use(middlewares.TimeoutMiddleware(30 * time.Second))
	r.Use(middlewares.ErrorMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	r.GET("/movies", moviesHandler.GetAll())
	r.GET("/movies/:id", moviesHandler.GetByID())

	if err := r.Run(":3000"); err != nil {
		slog.Error(err.Error())
		return
	}
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		slog.Error("Can't find .env file")
	}
}
