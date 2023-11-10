package controllers

import (
	"healthcare/configs"
	"healthcare/middlewares"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterDoctor.
func RegisterDoctorController(c echo.Context) error {
	var doctor web.DoctorRegisterRequest

	// Melakukan binding data dari request ke struct doctor
	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Register Data"))
	}
	doctorRequest := request.ConvertToDoctorRegisterRequest(doctor)

	// Memeriksa apakah email dokter sudah ada dalam database
	if existingDoctor := configs.DB.Where("email = ?", doctor.Email).First(&doctorRequest).Error; existingDoctor == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))
	}
	doctorRequest.Password = helper.HashPassword(doctorRequest.Password)

	// Menyimpan data dokter ke dalam database
	if err := configs.DB.Create(&doctorRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Register"))
	}
	response := response.ConvertToDoctorRegisterResponse(doctorRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Registered Successful", response))
}



// LoginDoctor.
func LoginDoctorController(c echo.Context) error {
	var loginRequest web.DoctorLoginRequest

	// Membinding data permintaan ke variabel loginRequest.
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Login Data"))
	}

	var doctor schema.Doctor
	// Menanyakan basis data untuk menemukan dokter dengan email dan password yang diberikan.
	if err := configs.DB.Where("email = ?", loginRequest.Email).First(&doctor).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Unregistered Email"))
	}

	// Membandingkan password yang diberikan dengan hash password yang tersimpan.
	if err := helper.ComparePassword(doctor.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Incorrect Password"))
	}
	doctorLoginResponse := response.ConvertToDoctorLoginResponse(&doctor)

	// Menghasilkan token JWT untuk dokter yang terautentikasi.
	token, err := middlewares.GenerateToken(&doctorLoginResponse, doctor.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Menghasilkan JWT"))
	}
	doctorLoginResponse.Token = token

	return c.JSON(http.StatusOK, helper.SuccessResponse("Login Berhasil", doctorLoginResponse))
}

