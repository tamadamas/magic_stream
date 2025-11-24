package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamadamas/magic_stream/server/go/internal/models"
	"github.com/tamadamas/magic_stream/server/go/internal/repositories"
)

type UsersHandler struct {
	repo *repositories.UsersRepository
}

func NewUsersHandler(repo *repositories.UsersRepository) *UsersHandler {
	return &UsersHandler{repo}
}

func (h *UsersHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var user models.User

		if err := c.ShouldBindJSON(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}

		if err := h.repo.Create(ctx, &user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "OK"})
	}
}
