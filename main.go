package main

import (
	"fmt"
	"healthcare/configs"
	"healthcare/middlewares"
	"healthcare/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {

	_ = godotenv.Load() // ignore error to anticipate server not run

	configs.Init()
	e := echo.New()

	// load middlewares
	middlewares.RemoveTrailingSlash(e)
	middlewares.Logger(e)
	middlewares.RateLimiter(e)
	middlewares.Recover(e)
	middlewares.CORS(e)

	// load router
	routes.SetupRoutes(e)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	setPort := fmt.Sprintf(":%s", port)
	e.Logger.Fatal(e.Start(setPort))
}
