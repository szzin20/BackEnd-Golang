package routes

import (
	"github.com/labstack/echo/v4"
	"healthcare/controllers"
	"healthcare/middlewares"
)

func SetupRoutes() *echo.Echo {

	e := echo.New()

	middlewares.RemoveTrailingSlash(e)
	middlewares.Logger(e)
	middlewares.RateLimiter(e)
	middlewares.Recover(e)
	middlewares.CORS(e)

	GroupAdmin := e.Group("/admins")
	GroupAdmin.POST("/medicines", controllers.CreateMedicineController)
	GroupAdmin.GET("/medicines", controllers.GetAllMedicinesAdminController)
	GroupAdmin.GET("/medicines/:id", controllers.GetMedicineController)
	GroupAdmin.GET("/medicines/?name=", controllers.GetMedicineByNameAdminController)
	GroupAdmin.PUT("/medicines/:id", controllers.UpdateMedicineController)
	GroupAdmin.DELETE("/medicines/:id", controllers.DeleteMedicineController)

	GroupPatient := e.Group("/Patients")
	GroupPatient.GET("/medicines", controllers.GetAllMedicinesPatientController, middlewares.UserIDRoleAuth)
	GroupPatient.GET("/medicines/?name=", controllers.GetMedicineByNamePatientController, middlewares.UserIDRoleAuth)

	return e

}
