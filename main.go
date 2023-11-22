package main

import (
	"fmt"
	"healthcare/configs"
	"healthcare/middlewares"
	"healthcare/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// requiredEnvVars := []string{
	// 	"DB_USERNAME",
	// 	"DB_PASSWORD",
	// 	"DB_HOST",
	// 	"DB_PORT",
	// 	"DB_NAME",
	// 	"JWT_SECRET",
	// 	"SERVER_PORT",
	// 	"BUCKET_NAME",
	// 	"BUCKET_SA",
	// }

	// allVarsExist := true

	// for _, envVar := range requiredEnvVars {
	// 	_, exists := os.LookupEnv(envVar)
	// 	if !exists {
	// 		log.Printf("Missing environment variable: %s\n", envVar)
	// 		allVarsExist = false
	// 	}
	// }

	// if !allVarsExist {
	// 	err := godotenv.Load()
	// 	if err != nil {
	// 		log.Println("Can't access .env files!")
	// 	}
	// }

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
