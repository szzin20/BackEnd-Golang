package routes

import (
	"healthcare/controllers"
	"healthcare/middlewares"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

func SetupRoutes() *echo.Echo {

	e := echo.New()

	middlewares.RemoveTrailingSlash(e)
	middlewares.Logger(e)
	middlewares.RateLimiter(e)
	middlewares.Recover(e)
	middlewares.CORS(e)

	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to Fetch .env File")
		}
	}

	JWT := middleware.JWT([]byte(os.Getenv("JWT_SECRET")))

	gUsers := e.Group("/users")
	gUsers.POST("/register", controllers.RegisterUserController)
	gUsers.POST("/login", controllers.LoginUserController)
	gUsers.GET("", controllers.GetAllUsersController, JWT)
	gUsers.GET("/:id", controllers.GetUserController, JWT)
	gUsers.PUT("/:id", controllers.UpdateUserController, JWT)
	gUsers.DELETE("/:id", controllers.DeleteUserController, JWT)

	gDocter := e.Group("/doctors")
	gDocter.POST("/register", controllers.)
	gDocter.POST("/login", controllers.)
	gDocter.GET("/:id", controllers., JWT)
	gDocter.PUT("/:id", controllers., JWT)
	gDocter.DELETE("/:id", controllers., JWT)

	return e

}
