package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tamadamas/magic_stream/server/go/internal/config"
	"github.com/tamadamas/magic_stream/server/go/internal/handlers"
	"github.com/tamadamas/magic_stream/server/go/internal/repositories"
)

func main() {
	loadEnv()
	db, err := config.ConnectToDatabase()

	if err != nil {
		log.Fatal(err)
	}

	moviesHandler := handlers.NewMoviesHandler(repositories.NewMoviesRepository(db))

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.GET("/movies", moviesHandler.GetAll())
	router.GET("/movies/:id", moviesHandler.GetByID())

	if err := router.Run(":3000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Can't find .env file")
	}
}
