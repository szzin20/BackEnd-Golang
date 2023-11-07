package main

import (
	"healthcare/middlewares"
	"healthcare/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("can't access .env files!")
	}

	e := echo.New()

	routes.SetupRoutes(e)

	// load middlewares
	middlewares.CORS(e)
	middlewares.Logger(e)
	middlewares.RateLimiter(e)
	middlewares.Recover(e)
	middlewares.RemoveTrailingSlash(e)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
