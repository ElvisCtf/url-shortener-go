package main

import (
	"log"
	"redirect-service/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
        log.Println("No .env file found")
    }

    addr := util.Env("ADDR", ":8081")

	router := gin.Default()
	if err := router.Run(addr); err != nil {
		log.Fatalf("startup service failed, err: %v\n", err)
	}
}