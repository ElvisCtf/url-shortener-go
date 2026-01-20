package router

import (
	"shorten-service/internal/handler"
    "github.com/gin-gonic/gin"
)

func SetupRouter(baseURL string) *gin.Engine {
	router := gin.Default()

	// register POST /shorten
	handler.RegisterShorten(router, baseURL)

	return router
}