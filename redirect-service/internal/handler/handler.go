package handler

import (
	"redirect-service/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRedirect(router *gin.Engine, service *service.Redirect) {
	router.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")
		c.JSON(200, gin.H{
			"received_code": code,
		})
	})
}