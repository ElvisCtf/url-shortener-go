package main

import (
	"log"

	"redirect-service/internal/repository"
	"redirect-service/internal/router"
	"redirect-service/internal/service"
	"redirect-service/internal/util"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
        log.Println("No .env file found")
    }
	addr := util.Env("ADDR", ":8081")

	repo := repository.NewRepo()
	service := service.NewRedirect(repo)
	router := router.SetupRouter(service)
	
	if err := router.Run(addr); err != nil {
		log.Fatalf("startup service failed, err: %v\n", err)
	}
}