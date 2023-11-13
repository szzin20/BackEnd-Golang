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

	gComplaints := e.Group("/complaints")
	gComplaints.POST("/:transaction_id", controllers.CreateComplaintController, JWT)
	gComplaints.GET("", controllers.GetAllComplaintsController, JWT)
	gComplaints.GET("/:transaction_id", controllers.GetComplaintController, JWT)

	gAdvices := e.Group("/advices")
	gAdvices.POST("/:complaint_id", controllers.CreateAdviceController, JWT)
	gAdvices.GET("", controllers.GetAllAdvicesController, JWT)
	gAdvices.GET("/:complaint_id", controllers.GetAdviceController, JWT)

	return e

}
