package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamadamas/magic_stream/server/go/internal/app_errors"
	"github.com/tamadamas/magic_stream/server/go/internal/models"
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
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, movies)
	}
}

func (h *MoviesHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("imdb_id")
		ctx := c.Request.Context()

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
			return
		}

		movie, err := h.repo.Find(ctx, id)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.Error(app_errors.NewNotFoundError(err, "Movie not found"))
				return
			}

			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, movie)
	}
}

func (h *MoviesHandler) AddMovie() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var movie models.Movie

		if err := c.ShouldBindJSON(movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}

		if err := h.repo.Create(ctx, &movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "OK"})
	}
}

func (h *MoviesHandler) GenresList() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx := c.Request.Context()
	}
}
