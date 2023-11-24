package routes

import (
	"healthcare/controllers"
	"healthcare/middlewares"
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

	AdminJWT := middlewares.AdminRoleAuth
	UserJWT := middlewares.UserIDRoleAuth
	DoctorJWT := middlewares.DoctorIDRoleAuth

	gAdmins := e.Group("/admins")
	gAdmins.POST("/login", controllers.LoginAdminController)
	gAdmins.POST("/register/doctor", controllers.RegisterDoctorByAdminController, AdminJWT)
	gAdmins.GET("/list/doctors", controllers.GetAllDoctorByAdminController, AdminJWT)
	gAdmins.PUT("/update/doctor/:id", controllers.UpdateDoctorByAdminController, AdminJWT)
	gAdmins.PUT("/update/payment/:id", controllers.UpdatePaymentStatusByAdminController, AdminJWT)
	gAdmins.DELETE("/delete/doctor/:id", controllers.DeleteDoctorByAdminController, AdminJWT)
	gAdmins.PUT("/:id", controllers.UpdateAdminController, AdminJWT)
	gAdmins.POST("/medicines", controllers.CreateMedicineController, AdminJWT)
	gAdmins.GET("/medicines", controllers.GetAllMedicinesAdminController)
	gAdmins.GET("/medicines/:id", controllers.GetMedicineController)
	gAdmins.GET("/medicine", controllers.GetMedicineByNameAdminController)
	gAdmins.PUT("/medicines/:id", controllers.UpdateMedicineController, AdminJWT)
	gAdmins.DELETE("/medicines/:id", controllers.DeleteMedicineController, AdminJWT)
	gAdmins.GET("/medicines/:id/image", controllers.GetImageMedicineController)
	gAdmins.PUT("/medicines/:id/image", controllers.UpdateImageMedicineController, AdminJWT)
	gAdmins.DELETE("/medicines/:id/image", controllers.DeleteImageMedicineController, AdminJWT)

	gUsers := e.Group("/users")
	gUsers.POST("/register", controllers.RegisterUserController)
	gUsers.POST("/login", controllers.LoginUserController)
	gUsers.GET("/profile", controllers.GetUserController, UserJWT)
	gUsers.PUT("/profile", controllers.UpdateUserController, UserJWT)
	gUsers.DELETE("", controllers.DeleteUserController, UserJWT)
	gUsers.GET("/medicines", controllers.GetAllMedicinesUserController)
	gUsers.GET("/medicines/:id", controllers.GetMedicineUserController)
	gUsers.GET("/medicine", controllers.GetMedicineByNameUserController)
	gUsers.GET("/doctors/available", controllers.GetAvailableDoctor)
	gUsers.GET("/doctors", controllers.GetSpecializeDoctor)
	gUsers.GET("/articles", controllers.GetAllArticles)
	gUsers.GET("/articles/:id", controllers.GetArticleByID)
	gUsers.GET("/article", controllers.GetAllArticlesByTitle)
	gUsers.POST("/doctor-payments", controllers.CreateDoctorTransactionController, UserJWT)
	gUsers.GET("/doctor-payments", controllers.GetAllDoctorTransactionsController, UserJWT) 
	gUsers.GET("/doctor-payment", controllers.GetDoctorTransactionsController, UserJWT)


	gDoctors := e.Group("/doctors")
	gDoctors.POST("/login", controllers.LoginDoctorController)
	gDoctors.GET("/profile", controllers.GetDoctorProfileController, DoctorJWT)
	gDoctors.GET("/:id", controllers.GetDoctorByIDController)
	gDoctors.PUT("/profile", controllers.UpdateDoctorController, DoctorJWT)
	gDoctors.DELETE("", controllers.DeleteDoctorController, DoctorJWT)
	gDoctors.GET("", controllers.GetAllDoctorController)
	gDoctors.GET("/articles", controllers.DoctorGetAllArticles, DoctorJWT)
	gDoctors.GET("/articles/:id", controllers.DoctorGetArticleByID, DoctorJWT)
	gDoctors.POST("/articles", controllers.CreateArticle, DoctorJWT)
	gDoctors.PUT("/articles/:id", controllers.UpdateArticleById, DoctorJWT)
	gDoctors.DELETE("/articles/:id", controllers.DeleteArticleById, DoctorJWT)

	// gDoctors.GET("manage/patitient", controllers.GetAllPatientsController, DoctorJWT)
	// gDoctors.GET("manage/patitient/:status", controllers.GetPatientsByStatusController, DoctorJWT)
	// gDoctors.PUT("manage/patitient/:idTransaksi", controllers.UpdatePatientController, DoctorJWT)

}
