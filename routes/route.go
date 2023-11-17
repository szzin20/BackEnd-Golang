package routes

import (
	"healthcare/controllers"
	"healthcare/middlewares"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	doctorGroup := e.Group("/doctors")
	// e.POST("/register/doctor", controllers.RegisterDoctorController)
	e.POST("/login/doctor", controllers.LoginDoctorController)
	// Doctor AUTH
	doctorGroup.GET("/profile", controllers.GetDoctorProfileController, middlewares.DoctorIDRoleAuth)
	doctorGroup.PUT("/update/profile", controllers.UpdateDoctorController, middlewares.DoctorIDRoleAuth)
	doctorGroup.DELETE("/delete/profile", controllers.DeleteDoctorController, middlewares.DoctorIDRoleAuth)
	doctorGroup.GET("", controllers.GetAllDoctorController, middlewares.DoctorIDRoleAuth)
	doctorGroup.POST("/logout-doctor", controllers.LogoutDoctorController, middlewares.DoctorIDRoleAuth)
	// Doctor Complaint
	// doctorGroup.GET("/complaints", controllers.GetAllDoctorComplaints, middlewares.DoctorIDRoleAuth)
	// doctorGroup.GET("/complaints/status", controllers.GetDoctorComplaintsByStatus, middlewares.DoctorIDRoleAuth)
	// doctorGroup.PUT("/complaints/:complaintID/status", controllers.UpdateDoctorComplaintStatus, middlewares.DoctorIDRoleAuth)
	// // Doctor Patient
	// doctorGroup.GET("/patients", controllers.GetDoctorPatientsController, middlewares.DoctorIDRoleAuth)
	// doctorGroup.GET("/patients/:status", controllers.GetDoctorPatientsByStatus, middlewares.DoctorIDRoleAuth)
	// doctorGroup.PUT("/patient-status", controllers.UpdatePatientStatusController, middlewares.DoctorIDRoleAuth)

	return

}

