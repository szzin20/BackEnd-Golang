package controllers

import (
	"healthcare/configs"
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

// User Create Complaint
func CreateComplaintController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	transactionID, _ := strconv.Atoi(c.QueryParam("transaction_id"))

	var doctorTransaction schema.DoctorTransaction

	if err := configs.DB.First(&doctorTransaction, "user_id = ? AND id = ?", userID, transactionID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
	}

	var complaintRequest web.CreateComplaintRequest

	if err := c.Bind(&complaintRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input complaint data"))
	}

	if err := helper.ValidateStruct(complaintRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	err := c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	file, fileHeader, err := c.Request().FormFile("image")

	if err == nil {
		defer file.Close()

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
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid image file format. supported formats: .jpg, .jpeg, .png"))
		}

		complaintImage, err := helper.UploadFilesToGCS(c, fileHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error uploading image to cloud storage"))
		}

		complaintRequest.Image = complaintImage
	}

	complaint := request.ConvertToComplaintRequest(complaintRequest, uint(transactionID))

	if err := configs.DB.Create(&complaint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send complaint"))
	}

	response := response.ConvertToCreateComplaintResponse(complaint)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("complaint successful", response))
}

// Get Complaint by DoctorTransaction ID
func GetComplaintsController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	transactionID, _ := strconv.Atoi(c.QueryParam("transaction_id"))

	var doctorTransaction schema.DoctorTransaction

	if err := configs.DB.Preload("Complaint").Preload("Advice").Where("user_id = ? AND id = ?", userID, transactionID).First(&doctorTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
	}

	response := response.ConvertToDoctorTransactionResponse(&doctorTransaction)

	return c.JSON(http.StatusOK, helper.SuccessResponse("complaint data successfully retrieved", response))
}

// // GetAllDataController untuk mengambil data transaksi dokter berdasarkan beberapa parameter.
// func GetAllDataController(c echo.Context) error {
// 	userID, ok := c.Get("userID").(int)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Complaint ID"))
// 	}

// 	// Mengambil nilai transaction_id dari query parameter
// 	transactionID, err := strconv.Atoi(c.QueryParam("transaction_id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Transaction ID"))
// 	}

// 	patientStatus := c.QueryParam("patient_status")

// 	// Jika transactionID dan patientStatus kosong, ambil semua transaksi untuk dokter tersebut
// 	if transactionID == 0 && patientStatus == "" {
// 		// Mengambil semua transaksi dokter yang belum dihapus
// 		var doctorTransactions []schema.DoctorTransaction
// 		if err := configs.DB.Where("deleted_at IS NULL AND user_id = ?", userID).Find(&doctorTransactions).Error; err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Fetch Doctor Transactions"))
// 		}

// 		// Membuat response berisi data transaksi dokter dan informasi terkait
// 		var responses []web.ComplaintsResponse
// 		for _, doctorTransaction := range doctorTransactions {
// 			var doctor schema.Doctor
// 			if err := configs.DB.Find(&doctor, "id=?", doctorTransaction.DoctorID).Error; err != nil {
// 				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Fetch Doctor Data"))
// 			}
// 			var userComplaint schema.User

// 			responses = append(responses, response.ConvertToUserComplaintResponse(userComplaint, doctorTransaction, doctor))
// 		}
// 		return c.JSON(http.StatusOK, helper.SuccessResponse("Doctor Transactions Data Successfully Retrieved", responses))
// 	}

// 	// Jika patientStatus tidak kosong, filter transaksi berdasarkan status pasien
// 	if patientStatus != "" {
// 		// Filter transaksi dokter berdasarkan ID dokter dan status pasien
// 		var doctorTransactions []schema.DoctorTransaction
// 		if err := configs.DB.Where("user_id = ? AND patient_status = ?", userID, patientStatus).Find(&doctorTransactions).Error; err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Fetch Doctor Transactions"))
// 		}

// 		// Membuat response berisi data transaksi dokter dan informasi terkait
// 		var responses []web.ComplaintsResponse
// 		for _, doctorTransaction := range doctorTransactions {
// 			var doctor schema.Doctor
// 			if err := configs.DB.Find(&doctor, "id=?", doctorTransaction.DoctorID).Error; err != nil {
// 				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Fetch Doctor Data"))
// 			}

// 			// Mendapatkan data (user complaint)
// 			var userComplaint schema.User
		
// 			responses = append(responses, response.ConvertToUserComplaintResponse(userComplaint, doctorTransaction, doctor))
// 		}
// 		return c.JSON(http.StatusOK, helper.SuccessResponse("Doctor Transactions Data Successfully Retrieved", responses))
// 	}
// 	return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Parameters"))
// }

// // Update Complaint or User
// func UpdateComplaintDataController(c echo.Context) error {
// 	dokterID, ok := c.Get("userID").(int)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil ID Dokter"))
// 	}

// 	// Mendapatkan ID transaksi dari parameter query
// 	id, err := strconv.Atoi(c.QueryParam("transaction_id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Gagal mendapatkan ID Transaksi Dokter"))
// 	}

// 	// Membind data permintaan ke dalam struct UpdateComplaintRequest
// 	var complaintRequest web.UpdateComplaintRequest
// 	if err := c.Bind(&complaintRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Input tidak valid untuk pembaruan data pasien"))
// 	}

// 	// Validasi data permintaan
// 	if err := helper.ValidateStruct(complaintRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
// 	}

// 	// Mengambil transaksi dokter dari database berdasarkan ID
// 	var existingDoctorTransaction schema.DoctorTransaction
// 	if err := configs.DB.First(&existingDoctorTransaction, id).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data transaksi dokter tidak ditemukan"))
// 	}

// 	// Memastikan dokter yang masuk memiliki hak kepemilikan transaksi
// 	if uint(dokterID) != existingDoctorTransaction.DoctorID {
// 		return c.JSON(http.StatusForbidden, helper.ErrorResponse("Anda tidak memiliki izin untuk memperbarui data transaksi ini"))
// 	}

// 	// Memperbarui status dan rincian kesehatan
// 	existingDoctorTransaction.PatientStatus = complaintRequest.PatientStatus
// 	existingDoctorTransaction.HealthDetails = complaintRequest.HealthDetails

// 	// Menyimpan perubahan ke database
// 	if err := configs.DB.Save(&existingDoctorTransaction).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal menyimpan transaksi dokter ke database"))
// 	}

// 	// Mengambil data pengguna yang terkait dengan transaksi
// 	var user schema.User
// 	if err := configs.DB.First(&user, existingDoctorTransaction.UserID).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data pengguna tidak ditemukan"))
// 	}

// 	response := response.ConvertToComplaintsResponse(user, existingDoctorTransaction)

// 	return c.JSON(http.StatusOK, helper.SuccessResponse("Data transaksi dokter berhasil diperbarui", response))
// }
