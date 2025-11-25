package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tamadamas/magic_stream/server/go/internal/config"
	"github.com/tamadamas/magic_stream/server/go/internal/middlewares"
	"github.com/tamadamas/magic_stream/server/go/internal/routes"
)

func main() {
	loadEnv()
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	db := config.ConnectToDatabase()

	r := gin.Default()
	r.Use(middlewares.RequestIDMiddleware())
	r.Use(middlewares.TimeoutMiddleware(30 * time.Second))
	r.Use(middlewares.ErrorMiddleware())

	routes.SetupUnProtectedRoutes(r, db)
	routes.SetupProtectedRoutes(r, db)

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
