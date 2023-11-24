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

	"github.com/jinzhu/gorm"
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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Data Login Tidak Valid"))
	}

	if err := helper.ValidateStruct(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	var doctor schema.Doctor
	if err := configs.DB.Where("email = ?", loginRequest.Email).First(&doctor).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse( "Email Tidak Terdaftar."))
	}

	if err := helper.ComparePassword(doctor.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Kata Sandi Salah" ))
	}

	// The rest of your code for generating a token and handling the successful login
	token, err := middlewares.GenerateToken(doctor.ID, doctor.Email, doctor.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse( "Gagal Menghasilkan JWT: "+err.Error()))
	}

	doctorLoginResponse := response.ConvertToDoctorLoginResponse(&doctor)
	doctorLoginResponse.Token = token

	// Send login notification email
	if doctor.Email != "" {
		notificationType := "login"
		if err := helper.SendNotificationEmail(doctor.Email, doctor.Fullname, notificationType, "drg"); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengirim email notifikasi: "+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse( "Login Berhasil", doctorLoginResponse))
}

func GetAvailableDoctor(c echo.Context) error {

	var Doctor []schema.Doctor

	err := configs.DB.Where("status = ?", true).Find(&Doctor).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Mengambil Data Dokter"))
	}

	if len(Doctor) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data tidak ditemukan"))
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
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data tidak ditemukan"))
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
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("Dokter tidak ditemukan"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil Profil Dokter"))
		}
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
	file, fileHeader, err := c.Request().FormFile("profile_picture")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("File Gambar Diperlukan"))
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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse( "Format file gambar tidak valid,Format yang didukung: jpg, jpeg, png"))
	}

	// Upload the image to Cloud Storage
	ProfilePicture, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error mengunggah gambar ke Cloud Storage"))
	}

	// Update the doctor details
	existingDoctor.ProfilePicture = ProfilePicture
	existingDoctor.Status = doctorUpdated.Status
	if err := configs.DB.Model(&existingDoctor).Updates(doctorUpdated).Error; err != nil {
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

// Manage patient

// GetAllPatientsController
// func GetAllPatientsController(c echo.Context) error {
//     dokterID, ok := c.Get("userID").(int)
//     if !ok {
//         return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil ID Dokter"))
//     }

//     // Ambil transaksi dokter dari database
//     var doctorTransactions []schema.DoctorTransaction
//     if err := configs.DB.Where("doctor_id = ?", dokterID).Find(&doctorTransactions).Error; err != nil {
//         return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data transaksi dokter"))
//     }

//     // Periksa jika tidak ada transaksi yang ditemukan
//     if len(doctorTransactions) == 0 {
//         return c.JSON(http.StatusNotFound, helper.ErrorResponse("Tidak ada data transaksi dokter"))
//     }

//     patientResponses := response.ConvertToDoctorPatientResponses(doctorTransactions)

//     return c.JSON(http.StatusOK, helper.SuccessResponse("Data pasien berhasil diambil", patientResponses))
// }

// func GetPatientsByStatusController(c echo.Context) error {
// 	dokterID, ok := c.Get("userID").(int)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil ID Dokter"))
// 	}
// 	status := c.Param("status")

// 	// Validasi bahwa status tidak boleh kosong
// 	if status == "" {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Status tidak boleh kosong"))
// 	}

// 	// Ambil transaksi dokter dari database berdasarkan ID dokter dan status
// 	var doctorTransactions []schema.DoctorTransaction
// 	if err := configs.DB.Where("doctor_id = ? AND patient_status = ?", dokterID, status).Find(&doctorTransactions).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data transaksi dokter"))
// 	}

// 	// Jika tidak ada transaksi dokter yang ditemukan, kembalikan respons Not Found
// 	if len(doctorTransactions) == 0 {
// 		return c.JSON(http.StatusNotFound, helper.ErrorResponse(fmt.Sprintf("Tidak ada data transaksi dokter dengan status %s", status)))
// 	}

// 	patientResponses := response.ConvertToDoctorPatientResponses(doctorTransactions)

// 	// Bangun pesan keberhasilan
// 	successMessage := fmt.Sprintf("Data pasien dengan status %s berhasil diambil", status)
// 	return c.JSON(http.StatusOK, helper.SuccessResponse(successMessage, patientResponses))
// }

// func UpdatePatientController(c echo.Context) error {
// 	dokterID, ok := c.Get("userID").(int)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil ID Dokter"))
// 	}

// 	// Mendapatkan ID transaksi dokter
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Gagal mendapatkan ID Transaksi Dokter"))
// 	}

// 	// Membanding data permintaan ke dalam struktur DoctorPatientRequest
// 	var patientRequest web.DoctorPatientRequest
// 	if err := c.Bind(&patientRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Input tidak valid untuk pembaruan data pasien"))
// 	}
// 	if err := helper.ValidateStruct(patientRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
// 	}

// 	// Mengambil data transaksi dokter dari database berdasarkan ID
// 	var existingDoctorTransaction schema.DoctorTransaction
// 	if err := configs.DB.First(&existingDoctorTransaction, id).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data transaksi dokter tidak ditemukan"))
// 	}

// 	// Memastikan transaksi dokter milik dokter yang sedang login
// 	if uint(dokterID) != existingDoctorTransaction.DoctorID {
// 		return c.JSON(http.StatusForbidden, helper.ErrorResponse("Anda tidak memiliki izin untuk memperbarui data transaksi ini"))
// 	}

// 	// Memperbarui status dan Health Details
// 	existingDoctorTransaction.PatientStatus = patientRequest.PatientStatus
// 	existingDoctorTransaction.HealthDetails = patientRequest.HealthDetails

// 	// Menyimpan perubahan ke dalam database
// 	if err := configs.DB.Save(&existingDoctorTransaction).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal menyimpan transaksi dokter ke database"))
// 	}

// 	var patientUser schema.User
// 	if err := configs.DB.First(&patientUser, existingDoctorTransaction.UserID).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data pengguna tidak ditemukan"))
// 	}

// 	response := response.ConvertTopatientDoctorTransaksiResponse(patientUser, existingDoctorTransaction)

// 	return c.JSON(http.StatusOK, helper.SuccessResponse("Data transaksi dokter berhasil diperbarui", response))
// }
