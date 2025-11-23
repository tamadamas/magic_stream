package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamadamas/magic_stream/server/go/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MoviesHandler struct {
	repo *repositories.MoviesRepository
}

func NewMoviesHandler(repo *repositories.MoviesRepository) *MoviesHandler {
	return &MoviesHandler{repo}
}

func (h *MoviesHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		movies, err := h.repo.All(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, movies)
	}
}

func (h *MoviesHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ctx := c.Request.Context()

		movie, err := h.repo.Find(ctx, id)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, movie)
	}

}
