package routes

import (
	"healthcare/controllers"
	"healthcare/middlewares"

	"github.com/labstack/echo/v4"
)

func SetupRoutes() *echo.Echo {
	e := echo.New()

	middlewares.RemoveTrailingSlash(e)
	middlewares.Logger(e)
	middlewares.RateLimiter(e)
	middlewares.Recover(e)
	middlewares.CORS(e)

	GroupDoctor := e.Group("/doctors")
	e.POST("/doctors/register", controllers.RegisterDoctorController)
	e.POST("/doctors/login", controllers.LoginDoctorController)
	GroupDoctor.PUT("/:id", controllers.UpdateDoctorController, middlewares.DoctorIDRoleAuth)
	GroupDoctor.DELETE("/:id", controllers.DeleteDoctorController, middlewares.DoctorIDRoleAuth)
	GroupDoctor.GET("/profile", controllers.GetDoctorProfileController, middlewares.DoctorIDRoleAuth)
	// patients
	GroupDoctor.GET("/patients", controllers.GetDoctorPatientsController, middlewares.DoctorIDRoleAuth)
	GroupDoctor.GET("/patients/:status", controllers.GetDoctorPatientsByStatus, middlewares.DoctorIDRoleAuth)
	// complaints
	GroupDoctor.GET("/complaints", controllers.GetAllDoctorComplaints, middlewares.DoctorIDRoleAuth)
	GroupDoctor.GET("/complaints/:status", controllers.GetDoctorComplaintsByStatus, middlewares.DoctorIDRoleAuth)
	GroupDoctor.PUT("/complaints/:status", controllers.UpdateDoctorComplaintStatus, middlewares.DoctorIDRoleAuth)
	return e
}
