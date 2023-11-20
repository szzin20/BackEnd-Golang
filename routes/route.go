package routes

import (
	"github.com/labstack/echo/v4"
	"healthcare/controllers"
	"healthcare/middlewares"
)

func SetupRoutes(e *echo.Echo) {

	UserJWT := middlewares.UserIDRoleAuth
	DoctorJWT := middlewares.DoctorIDRoleAuth

	gAdmins := e.Group("/admins")
	gAdmins.POST("/login", controllers.LoginAdminController)
	gAdmins.PUT("/:id", controllers.UpdateAdminController)
	gAdmins.GET("/list/users", controllers.GetAllUserByAdminController)
	gAdmins.POST("/register/doctor", controllers.RegisterDoctorByAdminController) 
	gAdmins.GET("/list/doctors", controllers.GetAllDoctorByAdminController) 
	gAdmins.PUT("update/doctor/:id", controllers.UpdateDoctorByAdminController)
	gAdmins.PUT("update/payment/:id", controllers.UpdatePaymentStatusByAdminController)
	gAdmins.DELETE("delete/doctor/:id", controllers.DeleteDoctorByAdminController)
	gAdmins.POST("/medicines", controllers.CreateMedicineController)
	gAdmins.GET("/medicines", controllers.GetAllMedicinesAdminController)
	gAdmins.GET("/medicines/:id", controllers.GetMedicineController)
	gAdmins.GET("/medicines/?name=", controllers.GetMedicineByNameAdminController)
	gAdmins.PUT("/medicines/:id", controllers.UpdateMedicineController)
	gAdmins.DELETE("/medicines/:id", controllers.DeleteMedicineController)

	gUsers := e.Group("/users")
	gUsers.POST("/register", controllers.RegisterUserController)
	gUsers.POST("/login", controllers.LoginUserController)
	gUsers.GET("/profile", controllers.GetUserController, UserJWT)
	gUsers.PUT("/profile", controllers.UpdateUserController, UserJWT)
	gUsers.DELETE("", controllers.DeleteUserController, UserJWT)
	gUsers.GET("/medicines", controllers.GetAllMedicinesUserController, UserJWT)
	gUsers.GET("/medicines/?name=", controllers.GetMedicineByNameUserController, UserJWT)

	gDoctors := e.Group("/doctors")
	gDoctors.POST("/login", controllers.LoginDoctorController)
	gDoctors.GET("/profile", controllers.GetDoctorProfileController, DoctorJWT)
	gDoctors.PUT("/profile", controllers.UpdateDoctorController, DoctorJWT)
	gDoctors.DELETE("", controllers.DeleteDoctorController, DoctorJWT)
	gDoctors.GET("", controllers.GetAllDoctorController)
}
