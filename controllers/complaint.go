package controllers

import (
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// User Create Complaint
func CreateComplaintController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Transaction ID"))
	}

	var existingTransactionID schema.DoctorTransaction

	result := configs.DB.First(&existingTransactionID, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Transaction ID"))
	}

	var complaint web.ComplaintRequest

	if err := c.Bind(&complaint); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Complaint Data"))
	}

	complaintRequest := request.ConvertToComplaintRequest(complaint, existingTransactionID.ID)

	if err := configs.DB.Create(&complaintRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Send Complaint"))
	}

	response := response.ConvertToComplaintResponse(complaintRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Complaint Successful", response))
}

// User Get Complaint by ID
func GetComplaintsController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Complaint ID"))
	}

	var complaint schema.Complaint

	if err := configs.DB.First(&complaint, userID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Complaint Data"))
	}

	response := response.ConvertToComplaintResponse(&complaint)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Complaint Data Successfully Retrieved", response))
}

// GetAllDataController untuk mengambil data transaksi dokter berdasarkan beberapa parameter.
func GetAllDataController(c echo.Context) error {
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Doctor ID"))
	}

	// Mengambil nilai transaction_id dari query parameter
	transactionID, err := strconv.Atoi(c.QueryParam("transaction_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Transaction ID"))
	}

	patientStatus := c.QueryParam("patient_status")

	// Jika transactionID dan patientStatus kosong, ambil semua transaksi untuk dokter tersebut
	if transactionID == 0 && patientStatus == "" {
		// Mengambil semua transaksi dokter yang belum dihapus
		var doctorTransactions []schema.DoctorTransaction
		if err := configs.DB.Where("deleted_at IS NULL AND user_id = ?", userID).Find(&doctorTransactions).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Fetch Doctor Transactions"))
		}

		// Membuat response berisi data transaksi dokter dan informasi terkait
		var responses []web.ComplaintsResponse
		for _, doctorTransaction := range doctorTransactions {
			var doctor schema.Doctor
			if err := configs.DB.Find(&doctor, "id=?", doctorTransaction.DoctorID).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Fetch Doctor Data"))
			}
			var userComplaint schema.User

			responses = append(responses, response.ConvertToUserComplaintResponse(userComplaint, doctorTransaction, doctor))
		}
		return c.JSON(http.StatusOK, helper.SuccessResponse("Doctor Transactions Data Successfully Retrieved", responses))
	}

	// Jika patientStatus tidak kosong, filter transaksi berdasarkan status pasien
	if patientStatus != "" {
		// Filter transaksi dokter berdasarkan ID dokter dan status pasien
		var doctorTransactions []schema.DoctorTransaction
		if err := configs.DB.Where("user_id = ? AND patient_status = ?", userID, patientStatus).Find(&doctorTransactions).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Fetch Doctor Transactions"))
		}

		// Membuat response berisi data transaksi dokter dan informasi terkait
		var responses []web.ComplaintsResponse
		for _, doctorTransaction := range doctorTransactions {
			var doctor schema.Doctor
			if err := configs.DB.Find(&doctor, "id=?", doctorTransaction.DoctorID).Error; err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Fetch Doctor Data"))
			}

			// Mendapatkan data (user complaint)
			var userComplaint schema.User
		
			responses = append(responses, response.ConvertToUserComplaintResponse(userComplaint, doctorTransaction, doctor))
		}
		return c.JSON(http.StatusOK, helper.SuccessResponse("Doctor Transactions Data Successfully Retrieved", responses))
	}
	return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Parameters"))
}

// Update Complaint or User
func UpdateComplaintDataController(c echo.Context) error {
	dokterID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil ID Dokter"))
	}

	// Mendapatkan ID transaksi dari parameter query
	id, err := strconv.Atoi(c.QueryParam("transaction_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Gagal mendapatkan ID Transaksi Dokter"))
	}

	// Membind data permintaan ke dalam struct UpdateComplaintRequest
	var complaintRequest web.UpdateComplaintRequest
	if err := c.Bind(&complaintRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Input tidak valid untuk pembaruan data pasien"))
	}

	// Validasi data permintaan
	if err := helper.ValidateStruct(complaintRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Mengambil transaksi dokter dari database berdasarkan ID
	var existingDoctorTransaction schema.DoctorTransaction
	if err := configs.DB.First(&existingDoctorTransaction, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data transaksi dokter tidak ditemukan"))
	}

	// Memastikan dokter yang masuk memiliki hak kepemilikan transaksi
	if uint(dokterID) != existingDoctorTransaction.DoctorID {
		return c.JSON(http.StatusForbidden, helper.ErrorResponse("Anda tidak memiliki izin untuk memperbarui data transaksi ini"))
	}

	// Memperbarui status dan rincian kesehatan
	existingDoctorTransaction.PatientStatus = complaintRequest.PatientStatus
	existingDoctorTransaction.HealthDetails = complaintRequest.HealthDetails

	// Menyimpan perubahan ke database
	if err := configs.DB.Save(&existingDoctorTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal menyimpan transaksi dokter ke database"))
	}

	// Mengambil data pengguna yang terkait dengan transaksi
	var user schema.User
	if err := configs.DB.First(&user, existingDoctorTransaction.UserID).Error; err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data pengguna tidak ditemukan"))
	}

	response := response.ConvertToComplaintsResponse(user, existingDoctorTransaction)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data transaksi dokter berhasil diperbarui", response))
}
