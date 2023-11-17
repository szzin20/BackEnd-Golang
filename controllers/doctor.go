package controllers

import (
	"healthcare/configs"
	"healthcare/middlewares"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/response"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterDoctorController
// func RegisterDoctorController(c echo.Context) error {
// 	var doctor web.DoctorRegisterRequest

// 	if err := c.Bind(&doctor); err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Input Data Registrasi Tidak Valid"))
// 	}

// 	doctorRequest := request.ConvertToDoctorRegisterRequest(doctor)

// 	// Periksa apakah email sudah ada
// 	if existingDoctor := configs.DB.Where("email = ?", doctorRequest.Email).First(&doctorRequest).Error; existingDoctor == nil {
// 		return c.JSON(http.StatusConflict, helper.ErrorResponse("Email Sudah Ada"))
// 	}

// 	// Hash kata sandi
// 	doctorRequest.Password = helper.HashPassword(doctorRequest.Password)

// 	// Simpan data dokter ke database
// 	if err := configs.DB.Create(&doctorRequest).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Registrasi"))
// 	}

// 	// Mengirim email pemberitahuan
// 	err := helper.SendNotificationEmail(doctorRequest.Email, doctorRequest.Fullname, "register", "drg")
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengirim email verifikasi"))
// 	}

// 	response := response.ConvertToDoctorRegisterResponse(doctorRequest)

// 	return c.JSON(http.StatusCreated, helper.SuccessResponse("Selamat Pendaftaran sukses", response))
// }

// LoginDoctorController
func LoginDoctorController(c echo.Context) error {
	var loginRequest web.DoctorLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Login Data"))
	}

	var doctor schema.Doctor
	if err := configs.DB.Where("email = ?", loginRequest.Email).First(&doctor).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Email Not Registered"))
	}

	if err := helper.ComparePassword(doctor.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Incorrect Password"))
	}

	// Set the status to "Online" when the doctor logs in
	doctor.Status = "Online"
	if err := configs.DB.Save(&doctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to update doctor status"))
	}

	doctorLoginResponse := response.ConvertToDoctorLoginResponse(&doctor)

	token, err := middlewares.GenerateToken(doctor.ID, doctor.Email, doctor.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Generate JWT"))
	}

	doctorLoginResponse.Token = token

	// Send login notification email
	if doctor.Email != "" {
		notificationType := "login"
		if err := helper.SendNotificationEmail(doctor.Email, doctor.Fullname, notificationType, "drg"); err != nil {
			log.Println("Failed to send notification email:", err)
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Login Successful", doctorLoginResponse))
}

// LogoutDoctorController mengatur status dokter menjadi "Offline" saat logout
func LogoutDoctorController(c echo.Context) error {
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

	// Update status menjadi "Offline"
	existingDoctor.Status = "Offline"
	if err := configs.DB.Save(&existingDoctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal memperbarui status dokter"))
	}
	ResponseLogout := response.ConvertToDoctorLogoutResponse(&existingDoctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Logout berhasil", ResponseLogout))
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
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil Profil Dokter"))
	}

	response := response.ConvertToGetDoctorResponse(&doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Profil Dokter berhasil diambil", response))
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
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data dokter"))
	}

	// Hapus data dokter
	if err := configs.DB.Delete(&existingDoctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal menghapus dokter"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Akun dokter berhasil dihapus", nil))
}




// patients
// func GetDoctorPatientsController(c echo.Context) error {
// 	// Mendapatkan ID dokter dari konteks
// 	doctorID, ok := c.Get("userID").(uint)
// 	if !ok {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("ID dokter tidak valid"))
// 	}

// 	// Mengambil data dokter dari database berdasarkan ID
// 	var doctor schema.Doctor
// 	if err := configs.DB.Preload("DoctorTransaction").First(&doctor, doctorID).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.JSON(http.StatusNotFound, helper.ErrorResponse("Dokter tidak ditemukan"))
// 		}
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Kesalahan internal server"))
// 	}
// 	response := response.ConvertToDoctorPatientsResponse(doctor)
// 	return c.JSON(http.StatusOK, helper.SuccessResponse("Data pasien dokter berhasil diambil", response))
// }

// // dokter mendapatkan daftar pasien melalui status
// func GetDoctorPatientsByStatusController(c echo.Context) error {
// 	doctorID, ok := c.Get("userID").(uint)
// 	if !ok {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid doctor ID"))
// 	}

// 	// Ambil status dari parameter URL, misalnya "/doctor/patients/recovered"
// 	status := c.Param("status")

// 	var doctor schema.Doctor
// 	if err := configs.DB.Preload("DoctorTransaction").First(&doctor, doctorID).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.JSON(http.StatusNotFound, helper.ErrorResponse("Doctor not found"))
// 		}
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Internal server error"))
// 	}

// 	var response = map[string]interface{}{
// 		"doctor_id":       doctor.ID,
// 		"transaction_id":  0,
// 		"doctor_fullname": doctor.Fullname,
// 		"patients":        []map[string]interface{}{},
// 	}

// 	for _, transaction := range doctor.DoctorTransaction {
// 		// Hanya tambahkan pengguna dengan status yang sesuai
// 		if transaction.Status.PatientStatus == schema.PatientStatus(status) {
// 			var user schema.User
// 			if err := configs.DB.First(&user, transaction.UserID).Error; err != nil {
// 				if err == gorm.ErrRecordNotFound {
// 					return c.JSON(http.StatusNotFound, helper.ErrorResponse("User not found"))
// 				}
// 				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Internal server error"))
// 			}

// 			patientInfo := map[string]interface{}{
// 				"id":     user.ID,
// 				"name":   user.Fullname,
// 				"email":  user.Email,
// 				"gender": user.Gender,
// 				"status": transaction.Status,
// 			}

// 			response["patients"] = append(response["patients"].([]map[string]interface{}), patientInfo)
// 		}
// 	}
// 	return c.JSON(http.StatusOK, response)
// }

// // dokter mengubah status pasien
// func UpdatePatientStatusController(c echo.Context) error {
// 	doctorID, ok := c.Get("userID").(int)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to get doctor ID"))
// 	}

// 	var request schema.DoctorTransaction
// 	if err := c.Bind(&request); err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Failed to read request data"))
// 	}

// 	// Validate request parameters
// 	if err := helper.ValidateStruct(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
// 	}

// 	// Ensure UserID is included in the request
// 	if request.UserID == 0 {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("UserID not found in the request"))
// 	}

// 	// Get existing data from the database using doctorID and request.UserID
// 	var existingTransaction schema.DoctorTransaction
// 	result := configs.DB.
// 		Where("doctor_id = ? AND user_id = ?", doctorID, request.UserID).
// 		First(&existingTransaction)
// 	if result.Error != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to fetch doctor transaction data"))
// 	}

// 	// Check if the record exists
// 	if result.RowsAffected == 0 {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Record not found for the given doctor and user"))
// 	}

// 	// Fetch associated Doctor information using DoctorID
// 	var doctor schema.Doctor
// 	if err := configs.DB.First(&doctor, existingTransaction.DoctorID).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to fetch associated Doctor information"))
// 	}

// 	// Update status pasien dalam dokter_transaction
// 	existingTransaction.Status.PatientStatus = request.Status.PatientStatus

// 	// Simpan perubahan ke database
// 	if err := configs.DB.Save(&existingTransaction).Error; err != nil {
// 		log.Println("Gagal memperbarui status pasien di database:", err)
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal memperbarui status pasien"))
// 	}

// 	// Konversi objek DokterTransaksi ke objek respons yang diinginkan
// 	response := response.ConvertToDoctorTransactionResponse(&existingTransaction)

// 	// Respon berhasil dengan objek respons yang dihasilkan
// 	return c.JSON(http.StatusOK, helper.SuccessResponse("Status Pasien Diperbarui", response))

// }

// func GetDoctorPatientsController(c echo.Context) error {
// 	doctorID := c.Get("userID")
// 	// doctorID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid doctor ID"))
// 	}

// 	// Fetch doctor from the database
// 	var doctor schema.Doctor
// 	if err := configs.DB.Preload("DoctorTransaction").First(&doctor, doctorID).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.JSON(http.StatusNotFound, helper.ErrorResponse("Doctor not found"))
// 		}
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Internal server error"))
// 	}

// 	// Fetch associated user details based on UserID in DoctorTransaction
// 	var user schema.User
// 	if len(doctor.DoctorTransaction) > 0 {
// 		userID := doctor.DoctorTransaction[0].UserID
// 		if err := configs.DB.First(&user, userID).Error; err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				return c.JSON(http.StatusNotFound, helper.ErrorResponse("User not found"))
// 			}
// 			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Internal server error"))
// 		}
// 	}

// 	// Prepare the response
// 	response := response.ConvertToGetAllDoctorTransactionResponse(user)
// 	response := map[string]interface{}{
// 		"doctor_id":           doctor.ID,
// 		"fullname":            doctor.Fullname,
// 		"email":               doctor.Email,
// 		"status":              doctor.Status,
// 		"price":               doctor.Price,
// 		"tag":                 doctor.Tag,
// 		"profile_picture":     doctor.ProfilePicture,
// 		"registration_letter": doctor.RegistrationLetter
// 	}

// 	if len(doctor.DoctorTransaction) > 0 {
// 		response["user_fullname"] = user.Fullname
// 	}

// 	return c.JSON(http.StatusOK, helper.SuccessResponse("Data Pengguna Berhasil Diambil", response))
// }
