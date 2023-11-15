package controllers

import (
	"errors"
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

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// RegisterDoctorController
func RegisterDoctorController(c echo.Context) error {
	var doctor web.DoctorRegisterRequest

	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Input Data Registrasi Tidak Valid"))
	}

	doctorRequest := request.ConvertToDoctorRegisterRequest(doctor)

	// Periksa apakah email sudah ada
	if existingDoctor := configs.DB.Where("email = ?", doctorRequest.Email).First(&doctorRequest).Error; existingDoctor == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("Email Sudah Ada"))
	}

	// Hash kata sandi
	doctorRequest.Password = helper.HashPassword(doctorRequest.Password)

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

// LoginDoctorController
func LoginDoctorController(c echo.Context) error {
	var loginRequest web.DoctorLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Data Login Tidak Valid"))
	}

	var doctor schema.Doctor
	if err := configs.DB.Where("email = ?", loginRequest.Email).First(&doctor).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Email Tidak Terdaftar"))
	}

	if err := helper.ComparePassword(doctor.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Kata Sandi Salah"))
	}

	doctorLoginResponse := response.ConvertToDoctorLoginResponse(&doctor)

	token, err := middlewares.GenerateToken(int(doctor.ID), doctor.Email, doctor.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Menghasilkan JWT"))
	}

	doctorLoginResponse.Token = token

	// Mengirim notifikasi melalui email
	if doctor.Email != "" {
		jenisNotifikasi := "login"
		err := helper.SendNotificationEmail(doctor.Email, doctor.Fullname, jenisNotifikasi, "drg")
		if err != nil {
			fmt.Println("Gagal mengirim notifikasi melalui email:", err)
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Login Berhasil", doctorLoginResponse))
}

// UpdateDoctorController updates a doctor's information
func UpdateDoctorController(c echo.Context) error {
	// Ekstrak ID dokter dari token
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mendapatkan ID Dokter"))
	}

	// Ambil dokter yang sudah ada dari database menggunakan userID
	var existingDoctor schema.Doctor
	result := configs.DB.First(&existingDoctor, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data dokter"))
	}

	// Bind data permintaan pembaruan
	var doctorUpdated web.DoctorUpdateRequest
	if err := c.Bind(&doctorUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Input tidak valid untuk pembaruan data dokter"))
	}

	// Enkripsi kata sandi hanya jika kata sandi baru disediakan
	if doctorUpdated.Password != "" {
		doctorUpdated.Password = helper.HashPassword(doctorUpdated.Password)
	}

	// Perbarui data dokter dan periksa kesalahan selama pembaruan
	if err := configs.DB.Model(&existingDoctor).Updates(doctorUpdated).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal memperbarui data dokter"))
	}
	response := response.ConvertToDoctorUpdateResponse(&existingDoctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data dokter berhasil diperbarui", response))
}

// DoctorProfile
func GetDoctorProfileController(c echo.Context) error {
	// Ekstrak ID dokter dari token
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil ID Dokter"))
	}

	var doctor schema.Doctor
	if err := configs.DB.First(&doctor, userID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusConflict, helper.ErrorResponse("Profil Dokter tidak ditemukan"))
		}
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil Profil Dokter"))
	}

	response := response.ConvertToGetDoctorResponse(&doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Profil Dokter berhasil diambil", response))
}

// DeleteDoctor
func DeleteDoctorController(c echo.Context) error {
	// Ekstrak ID dokter dari token
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mendapatkan ID Dokter"))
	}

	// Ambil dokter yang sudah ada dari database menggunakan userID
	var existingDoctor schema.Doctor
	result := configs.DB.First(&existingDoctor, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("Dokter tidak ditemukan"))
		}
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data dokter"))
	}

	// Hapus data dokter
	if err := configs.DB.Delete(&existingDoctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal menghapus dokter"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Akun dokter berhasil dihapus", nil))
}

// GetAllDoctorsController retrieves a list of all doctors
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

// patients

// dokter mendapatkan daftar pasien
func GetDoctorPatientsController(c echo.Context) error {
	doctorID := c.Get("DoctorID")

	// Ambil semua pasien yang terkait dengan dokter melalui doctor_transaction
	var patients []schema.User
	if err := configs.DB.
		Joins("JOIN doctor_transactions ON users.id = doctor_transactions.user_id").
		Where("doctor_transactions.doctor_id = ?", doctorID).
		Find(&patients).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data pasien"))
	}

	// Ambil semua keluhan yang terkait dengan pasien-pasien tersebut
	var complaints []schema.Complaint
	for _, patient := range patients {
		var patientComplaints []schema.Complaint
		if err := configs.DB.Where("patient_id = ?", patient.ID).Find(&patientComplaints).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data keluhan"))
		}
		complaints = append(complaints, patientComplaints...)
	}

	response := response.ConvertToGetAllComplaintsResponse(complaints)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Daftar Pasien", response))
}

// dokter mendapatkan daftar pasien melalui status
func GetDoctorPatientsByStatus(c echo.Context) error {
	doctorID := c.Get("DoctorID")

	// Ambil semua pasien yang terkait dengan dokter melalui doctor_transaction
	var patients []schema.User
	if err := configs.DB.
		Joins("JOIN doctor_transactions ON users.id = doctor_transactions.user_id").
		Where("doctor_transactions.doctor_id = ?", doctorID).
		Find(&patients).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data pasien"))
	}

	// Dapatkan parameter status dari URL sebagai boolean
	status := c.Param("status")
	bStatus, err := strconv.ParseBool(status)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Parameter status tidak valid"))
	}

	// Validasi status
	if err := helper.ValidateStruct(bStatus); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Dapatkan semua keluhan terkait pasien berdasarkan status
	var complaints []schema.Complaint
	for _, patient := range patients {
		var patientComplaints []schema.Complaint
		if err := configs.DB.
			Where("patient_id = ? AND status = ?", patient.ID, bStatus).
			Find(&patientComplaints).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data keluhan"))
		}
		complaints = append(complaints, patientComplaints...)
	}

	responseData := response.ConvertToGetAllComplaintsResponse(complaints)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Daftar Keluhan Pasien Berdasarkan Status", responseData))
}

// dokter mengubah status pasien
func UpdatePatientStatusController(c echo.Context) error {
	doctorID, ok := c.Get("DoctorID").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mendapatkan ID dokter"))
	}

	// Dapatkan parameter dari request
	var request schema.DoctorTransaction
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Gagal membaca data permintaan"))
	}

	// Validasi parameter request
	if err := helper.ValidateStruct(request); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Perbarui status pasien dalam doctor_transaction
	if err := configs.DB.
		Model(&schema.DoctorTransaction{}).
		Where("doctor_id = ? AND user_id = ?", doctorID, request.UserID).
		Updates(map[string]interface{}{"status": request.Status}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal memperbarui status pasien"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Status Pasien Diperbarui", nil))
}
