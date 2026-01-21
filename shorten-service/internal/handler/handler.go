package handler

import (
	"net/http"

	"shorten-service/internal/model"
	"shorten-service/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterShorten(router *gin.Engine, service *service.Shorten) {
	router.POST("/shorten", func (c *gin.Context) {
		var request model.ShortenRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := service.Create(request.OriginalURL)
		c.JSON(http.StatusOK, response)
	})
}