package routes

import (
	"healthcare/controllers"
	"healthcare/middlewares"

	// echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	UserJWT := middlewares.UserIDRoleAuth
	DoctorJWT := middlewares.DoctorIDRoleAuth

	gAdmins := e.Group("/admins")
	gAdmins.POST("/login", controllers.LoginAdminController)
	gAdmins.PUT("/:id", controllers.UpdateAdminController)

	gUsers := e.Group("/users")
	// gUsers.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	gUsers.POST("/register", controllers.RegisterUserController)
	gUsers.POST("/login", controllers.LoginUserController)
	gUsers.GET("/profile", controllers.GetUserController, UserJWT)
	gUsers.PUT("/profile", controllers.UpdateUserController, UserJWT)
	gUsers.DELETE("", controllers.DeleteUserController, UserJWT)
	gUsers.GET("/doctor-payments", controllers.GetAllDoctorTransactionsController, UserJWT)
	gUsers.POST("/doctor-payments", controllers.CreateDoctorTransaction, UserJWT)
	gUsers.GET("/doctor-payments", controllers.GetDoctorTransactionController, UserJWT)
	// gUsers.GET("/doctor-payments", controllers.GetDoctorTransactionByStatusController, UserJWT)
	
	gDoctors := e.Group("/doctors")
	gDoctors.POST("/login", controllers.LoginDoctorController)
	gDoctors.GET("/profile", controllers.GetDoctorProfileController, DoctorJWT)
	gDoctors.PUT("/profile", controllers.UpdateDoctorController, DoctorJWT)
	gDoctors.DELETE("", controllers.DeleteDoctorController, DoctorJWT)
	gDoctors.GET("", controllers.GetAllDoctorController, DoctorJWT)

}
