package handler

import (
	"net/http"
	"shorten-service/internal/model"
	"github.com/gin-gonic/gin"
)

func RegisterShorten(router *gin.Engine, baseURL string) {
	router.POST("/shorten", func (c *gin.Context) {
		var request model.ShortenRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		code := "123"

		response := model.ShortenResponse{
			OriginalURL: request.OriginalURL,
			ShortenURL:  baseURL + "/" + code,
		}
		c.JSON(http.StatusOK, response)
	})
}