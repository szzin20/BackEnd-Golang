package main

import (
	"fmt"
	"healthcare/configs"
	"healthcare/middlewares"
	"healthcare/routes"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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
	ports, _ := strconv.Atoi(port)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", ports)))
}
