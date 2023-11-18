package routes

import (
	"healthcare/controllers"
	"healthcare/middlewares"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	doctorGroup := e.Group("/doctors")
	e.POST("/login/doctor", controllers.LoginDoctorController)
	// Doctor AUTH
	doctorGroup.GET("/profile", controllers.GetDoctorProfileController, middlewares.DoctorIDRoleAuth)
	doctorGroup.PUT("/update/profile", controllers.UpdateDoctorController, middlewares.DoctorIDRoleAuth)
	doctorGroup.DELETE("/delete/profile", controllers.DeleteDoctorController, middlewares.DoctorIDRoleAuth)
	doctorGroup.GET("", controllers.GetAllDoctorController, middlewares.DoctorIDRoleAuth)

	return
}

