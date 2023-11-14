package routes

import (
	"healthcare/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Rute untuk admin login
	e.POST("/admin/login", controllers.LoginAdminController)
	e.PUT("/admin/:id", controllers.UpdateAdminController)
}
