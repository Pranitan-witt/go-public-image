package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-public-image/service"
)

type movieHandler struct {
	Svc service.MovieService
}

func NewMovieHandler(msvc service.MovieService) *movieHandler {
	return &movieHandler{
		Svc: msvc,
	}
}

func (h *movieHandler) GetMovies(c *gin.Context) {
	result, err := h.Svc.GetMovies()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}
