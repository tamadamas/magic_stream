package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tamadamas/magic_stream/server/go/internal/handlers"
	"github.com/tamadamas/magic_stream/server/go/internal/middlewares"
	"github.com/tamadamas/magic_stream/server/go/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupProtectedRoutes(r *gin.Engine, db *mongo.Database) {
	r.Use(middlewares.AuthMiddleware())

	moviesHandler := handlers.NewMoviesHandler(repositories.NewMoviesRepository(db))

	r.GET("/movies/:imdb_id", moviesHandler.GetByID())
	r.POST("/movies", moviesHandler.AddMovie())
}
