
package main

import (
    "fmt"
    "os"

    "shorten-service/internal/repository"
    "shorten-service/internal/router"
    "shorten-service/internal/service"
)

func main() {
    addr := env("ADDR", ":8080")
    baseURL := env("BASE_URL", "http://localhost:8080")

    repo := repository.NewDevRepo()
    service := service.NewShorten(baseURL, repo)
    router := router.SetupRouter(service)

    if err := router.Run(addr); err != nil {
		fmt.Printf("startup service failed, err: %v\n", err)
	}
}

func env(key, defaultVal string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return defaultVal
}