package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamadamas/magic_stream/server/go/internal/handlers"
	"github.com/tamadamas/magic_stream/server/go/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupUnProtectedRoutes(r *gin.Engine, db *mongo.Database) {
	moviesHandler := handlers.NewMoviesHandler(repositories.NewMoviesRepository(db))
	usersHandler := handlers.NewUsersHandler(repositories.NewUsersRepository(db))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	r.GET("/movies", moviesHandler.GetAll())

	r.POST("/register", usersHandler.Register())
	r.POST("/login", usersHandler.Login())
	r.POST("/logout", usersHandler.Logout())
	r.GET("/genres", moviesHandler.GenresList())
	r.POST("/refresh", usersHandler.RefreshToken())
}
