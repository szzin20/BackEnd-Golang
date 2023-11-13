package controllers

import (
	"fmt"
	"healthcare/configs"
	"healthcare/middlewares"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// RegisterDoctor
func RegisterDoctorController(c echo.Context) error {
	var doctor web.DoctorRegisterRequest

	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Register Data"))
	}

	// Generate a verification code
	verificationCode := helper.GenerateVerificationCode()
	doctorRequest := request.ConvertToDoctorRegisterRequest(doctor)

	// Check if email already exists
	if existingDoctor := configs.DB.Where("email = ?", doctorRequest.Email).First(&doctorRequest).Error; existingDoctor == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))
	}

	// Hash the password
	doctorRequest.Password = helper.HashPassword(doctorRequest.Password)

	// Save the user to the database
	if err := configs.DB.Create(&doctorRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Register"))
	}

	// Send verification email
	err := helper.SendVerificationEmail(doctorRequest.Email, verificationCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to send verification email"))
	}

	response := response.ConvertToDoctorRegisterResponse(doctorRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Registered Successfully. Please check your email for verification instructions.", response))
}

// DoctorLogin
func LoginDoctorController(c echo.Context) error {
	var loginRequest web.DoctorLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Login Data"))
	}

	var doctor schema.Doctor
	// query := "SELECT id, email, password, role FROM doctors WHERE email = ?"
	// err := configs.DB.QueryRow(query, loginRequest.Email).Scan(&doctor.ID, &doctor.Email, &doctor.Password, &doctor.Role)


	if err := configs.DB.Where("email = ?", loginRequest.Email).First(&doctor).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Email Not Registered"))
	}

	if err := helper.ComparePassword(doctor.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Incorrect Email or Password"))
	}

	doctorLoginResponse := response.ConvertToDoctorLoginResponse(&doctor)

	token, err := middlewares.GenerateToken(doctor.ID, doctor.Email, doctor.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Generate JWT"))
	}

	doctorLoginResponse.Token = token

	// Sending email notification
	if doctor.Email != "" {
		body := "Hello, " + doctor.Fullname + "! You have successfully logged in."

		if err := helper.SendLoginNotificationEmail(doctor.Email, doctor.Fullname, body); err != nil {
			fmt.Println("Failed to send email notification:", err)
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Login Successful", doctorLoginResponse))
}

// UpdateDoctor
func UpdateDoctorController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid doctor ID"))
	}

	var existingDoctor schema.Doctor

	result := configs.DB.First(&existingDoctor, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Doctor"))
	}

	var doctorUpdated web.DoctorUpdateRequest

	if err := c.Bind(&doctorUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Update Data"))
	}

	// Enkripsi kata sandi hanya jika ada kata sandi yang diberikan
	if doctorUpdated.Password != "" {
		doctorUpdated.Password = helper.HashPassword(doctorUpdated.Password)
	}

	// Perbarui data dokter dan periksa apakah ada kesalahan selama pembaruan
	if err := configs.DB.Model(&existingDoctor).Updates(doctorUpdated).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to update doctor data"))
	}

	response := response.ConvertToDoctorUpdateResponse(&existingDoctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Doctor Updated Data Successful", response))
}

// DoctorProfile
func GetDoctorProfileController(c echo.Context) error {
	// Ekstrak ID dokter dari token
	doctorID:= c.Get("DoctorID")

	var doctor schema.Doctor
	if err := configs.DB.Preload("Doctor").Where("id = ?", doctorID).First(&doctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil Profil Dokter"))
	}

	response := response.ConvertToGetDoctorResponse(&doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Profil Dokter berhasil diambil", response))
}

// DeleteDoctor
func DeleteDoctorController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Doctor ID"))
	}

	var existingDoctor schema.Doctor
	result := configs.DB.First(&existingDoctor, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Doctor"))
	}

	configs.DB.Delete(&existingDoctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Doctor Deleted Data Successful", nil))
}

// dokter mendapatkan daftar pasien
func GetDoctorPatientsController(c echo.Context) error {
	doctorID:= c.Get("DoctorID")
	

	// Ambil data dokter dari database
	var doctor schema.Doctor
	if err := configs.DB.First(&doctor, doctorID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to retrieve Doctor Profile"))
	}

	// Ambil semua pasien
	var patients []schema.User
	if err := configs.DB.Find(&patients).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to retrieve patient data"))
	}

	// Ambil semua keluhan yang terkait dengan pasien-pasien tersebut
	var complaints []schema.Complaint
	for _, patient := range patients {
		var patientComplaints []schema.Complaint
		if err := configs.DB.Where("patient_id = ?", patient.ID).Find(&patientComplaints).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to retrieve complaint data"))
		}
		complaints = append(complaints, patientComplaints...)
	}

	response := response.ConvertToGetAllComplaintsResponse(complaints)

	return c.JSON(http.StatusOK, helper.SuccessResponse("List of Patient", response))
}

// dokter mendapatkan daftar pasien melalui status
func GetDoctorPatientsByStatus(c echo.Context) error {
	doctorID:= c.Get("DoctorID")

	// Ambil data dokter dari database
	var doctor schema.Doctor
	if err := configs.DB.First(&doctor, doctorID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to retrieve Doctor Profile"))
	}

	// Ambil semua pasien
	var patients []schema.User
	if err := configs.DB.Model(&doctor).Association("Patients").Find(&patients); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to retrieve patient data"))
	}

	// Get the status parameter from the URL as boolean
	status := c.Param("status")
	bStatus, err := strconv.ParseBool(status)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid status parameter"))
	}

	// Validasi status
	if err := helper.ValidateStruct(bStatus); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Dapatkan semua keluhan terkait pasien berdasarkan status
	var complaints []schema.Complaint
	for _, patient := range patients {
		var patientComplaints []schema.Complaint
		if err := configs.DB.Where("patient_id = ? AND status = ?", patient.ID, bStatus).Find(&patientComplaints).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to retrieve complaint data"))
		}
		complaints = append(complaints, patientComplaints...)
	}

	responseData := response.ConvertToGetAllComplaintsResponse(complaints)

	return c.JSON(http.StatusOK, helper.SuccessResponse("List of Patient Complaints Based on Status", responseData))
}

// dokter memperbarui status pasien
// func UpdateDoctorPatientStatus(c echo.Context) error {
// 	doctorID:= c.Get("DoctorID")

// 	// Ambil data dokter dari database
// 	var doctor schema.Doctor
// 	if err := configs.DB.First(&doctor, doctorID).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to retrieve Doctor Profile"))
// 	}

// 	// Ambil ID pasien dari path parameter
// 	patientID := c.Param("patientID")

// 	// Ambil data request dari body
// 	var updateRequest web.UserUpdateRequest
// 	if err := c.Bind(&updateRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Failed to process request"))
// 	}

// 	// Validasi data request
// 	if err := c.Validate(&updateRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
// 	}

// 	// Perbarui status pasien dalam database
// 	var patient schema.User
// 	if err := configs.DB.First(&patient, patientID).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Patient not found"))
// 	}

// 	// Update status pasien
// 	patient.Status = updateRequest.Status

// 	if err := configs.DB.Save(&patient).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to update patient status"))
// 	}

// 	// Mengembalikan respons JSON sukses dengan data pasien yang diperbarui
// 	return c.JSON(http.StatusOK, helper.SuccessResponse("Patient Status Updated", map[string]interface{}{"patient": &patient}))
// }
