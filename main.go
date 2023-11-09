package main

import (
	"healthcare/configs"
	"healthcare/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func main() {

	_, err := os.Stat(".env")
    if err == nil {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Failed to Fetch .env File")
        }
    }

	configs.Init()
	e := routes.SetupRoutes()

	port := os.Getenv("PORT")
	setPort := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(setPort))
}