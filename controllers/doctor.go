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
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

// RegisterDoctorController
func RegisterDoctorByAdminController(c echo.Context) error {
	var doctor web.DoctorRegisterRequest

	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Input Data Registrasi Tidak Valid"))
	}
	if err := helper.ValidateStruct(doctor); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	doctorRequest := request.ConvertToDoctorRegisterRequest(doctor)

	// Periksa apakah email sudah ada
	if existingDoctor := configs.DB.Where("email = ?", doctorRequest.Email).First(&doctorRequest).Error; existingDoctor == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("Email Sudah Ada"))
	}

	// Hash kata sandi
	doctorRequest.Password = helper.HashPassword(doctor.Password)

	// Simpan data dokter ke database
	if err := configs.DB.Create(&doctorRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Registrasi"))
	}

	// Mengirim email pemberitahuan
	err := helper.SendNotificationEmail(doctorRequest.Email, doctorRequest.Fullname, "register", "drg")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengirim email verifikasi"))
	}

	response := response.ConvertToDoctorRegisterResponse(doctorRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Selamat Pendaftaran sukses", response))
}

// Login Doctor
func LoginDoctorController(c echo.Context) error {
	var loginRequest web.DoctorLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Login Data"))
	}

	if err := helper.ValidateStruct(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	var doctor schema.Doctor
	if err := configs.DB.Where("email = ?", loginRequest.Email).First(&doctor).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Email Not Registered"))
	}

	if err := helper.ComparePassword(doctor.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Incorrect Password"))
	}

	// The rest of your code for generating a token and handling the successful login
	token, err := middlewares.GenerateToken(doctor.ID, doctor.Email, doctor.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Generate JWT: "+err.Error()))
	}

	doctorLoginResponse := response.ConvertToDoctorLoginResponse(&doctor)
	doctorLoginResponse.Token = token

	// Send login notification email
	if doctor.Email != "" {
		notificationType := "login"
		if err := helper.SendNotificationEmail(doctor.Email, doctor.Fullname, notificationType, "drg"); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to send notification email: "+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Login Successful", doctorLoginResponse))
}

func GetAvailableDoctor(c echo.Context) error {

	var Doctor []schema.Doctor

	err := configs.DB.Where("status = ?", "1").Find(&Doctor).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Mengambil Data Dokter"))
	}

	if len(Doctor) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data Dokter Kosong"))
	}

	response := response.ConvertToGetAllDoctorResponse(Doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data Dokter Berhasil Diambil", response))

}

func GetSpecializeDoctor(c echo.Context) error {
	specialist := c.QueryParam("specialist")

	if specialist == "" {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Parameter specialist required!"))
	}

	var doctors []schema.Doctor
	err := configs.DB.Where("specialist LIKE ?", "%"+specialist+"%").Find(&doctors).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Mengambil Data Dokter"))
	}

	if len(doctors) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data Dokter Kosong"))
	}

	response := response.ConvertToGetAllDoctorResponse(doctors)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data Dokter Berhasil Diambil", response))

}

// Get Doctor Profile
func GetDoctorProfileController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil ID Dokter"))
	}

	var doctor schema.Doctor
	if err := configs.DB.First(&doctor, userID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil Profil Dokter"))
	}

	response := response.ConvertToGetDoctorResponse(&doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Profil Dokter berhasil diambil", response))
}

// Get All Doctors
func GetAllDoctorController(c echo.Context) error {
	var Doctor []schema.Doctor

	err := configs.DB.Find(&Doctor).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Mengambil Data Pengguna"))
	}

	if len(Doctor) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data Pengguna Kosong"))
	}

	response := response.ConvertToGetAllDoctorResponse(Doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data Pengguna Berhasil Diambil", response))
}

// Get All Doctors by Admin
func GetAllDoctorByAdminController(c echo.Context) error {
	var Doctor []schema.Doctor

	err := configs.DB.Find(&Doctor).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Mengambil Data Pengguna"))
	}

	if len(Doctor) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data Pengguna Kosong"))
	}

	response := response.ConvertToGetAllDoctorByAdminResponse(Doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data Pengguna Berhasil Diambil", response))
}

// Update Doctor
func UpdateDoctorController(c echo.Context) error {
	// Get userID from the context
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mendapatkan ID Dokter"))
	}

	// Fetch the existing doctor based on userID
	var existingDoctor schema.Doctor
	result := configs.DB.First(&existingDoctor, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data dokter"))
	}

	// Parse the request body into the DoctorUpdateRequest struct
	var doctorUpdated web.DoctorUpdateRequest
	if err := c.Bind(&doctorUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Input tidak valid untuk pembaruan data dokter"))
	}

	// Validate the request payload
	if err := helper.ValidateStruct(doctorUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Hash the password if provided
	if doctorUpdated.Password != "" {
		doctorUpdated.Password = helper.HashPassword(doctorUpdated.Password)
	}

	// Parse multipart form for file upload
	err := c.Request().ParseMultipartForm(10 << 20)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Extract the image file from the form
	file, fileHeader, err := c.Request().FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Image File is Required"))
	}
	defer file.Close()

	// Check if the file format is allowed
	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	ext := filepath.Ext(fileHeader.Filename)
	allowed := false
	for _, validExt := range allowedExtensions {
		if ext == validExt {
			allowed = true
			break
		}
	}
	if !allowed {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid image file format. Supported formats: jpg, jpeg, png"))
	}

	// Upload the image to Cloud Storage
	ProfilePicture, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error upload image to Cloud Storage"))
	}

	// Update the doctor details
	existingDoctor.ProfilePicture = ProfilePicture
	existingDoctor.Status = doctorUpdated.Status
	if err := configs.DB.Model(&existingDoctor).Updates(existingDoctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal memperbarui data dokter"))
	}

	configs.DB.Save(&existingDoctor)

	response := response.ConvertToDoctorUpdateResponse(&existingDoctor)
	return c.JSON(http.StatusOK, helper.SuccessResponse("Data dokter berhasil diperbarui", response))
}

// Update Doctor by Admin
func UpdateDoctorByAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Gagal mendapatkan ID Dokter"))
	}

	var doctorUpdated web.DoctorUpdateRequest

	if err := c.Bind(&doctorUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Input tidak valid untuk pembaruan data dokter"))
	}

	if err := helper.ValidateStruct(doctorUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	if doctorUpdated.Password != "" {
		doctorUpdated.Password = helper.HashPassword(doctorUpdated.Password)
	}

	var existingDoctor schema.Doctor
	result := configs.DB.First(&existingDoctor, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal memperbarui data dokter"))
	}

	configs.DB.Model(&existingDoctor).Updates(doctorUpdated)

	response := response.ConvertToDoctorUpdateResponse(&existingDoctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data dokter berhasil diperbarui oleh admin", response))
}

// Delete Doctor
func DeleteDoctorController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mendapatkan ID Dokter"))
	}

	var existingDoctor schema.Doctor
	result := configs.DB.First(&existingDoctor, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data dokter"))
	}

	if err := configs.DB.Delete(&existingDoctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal menghapus dokter"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Akun dokter berhasil dihapus", nil))
}

// DeleteDoctorByAdminController deletes a doctor by admin
func DeleteDoctorByAdminController(c echo.Context) error {
	// Parse doctor ID from the request parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Gagal mendapatkan ID Dokter"))
	}

	// Retrieve the existing doctor from the database
	var existingDoctor schema.Doctor
	result := configs.DB.First(&existingDoctor, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data dokter"))
	}

	// Delete the doctor from the database
	result = configs.DB.Delete(&existingDoctor, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal menghapus dokter"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Akun dokter berhasil dihapus oleh admin  ", nil))
}

// Get Doctor by ID
func GetDoctorByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Gagal mendapatkan ID Dokter"))
	}
	var doctor schema.Doctor
	result := configs.DB.First(&doctor, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data dokter"))
	}
	response := response.ConvertToGetIDDoctorResponse(&doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Detail Dokter berhasil diambil", response))
}
