package routes

import (
	"healthcare/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	e.POST("/users/create/doctor-transaction", controllers.CreateDoctorTransaction)

}
