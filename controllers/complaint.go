package controllers

// import (
// 	"healthcare/configs"
// 	"healthcare/models/schema"
// 	"healthcare/models/web"
// 	"healthcare/utils/helper"
// 	"healthcare/utils/response"
// 	"net/http"
// "strconv"

// 	"github.com/labstack/echo/v4"
// )

// func GetAllDataController(c echo.Context) error {
// 	dokterID, ok := c.Get("userID").(int)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Dokter ID"))
// 	}

// 	// Mendapatkan nilai dari query parameter
// 	patientStatus := c.QueryParam("patientStatus")

// 	var (
// 		doctorTransactions []schema.DoctorTransaction
// 		doctor             schema.Doctor
// 		responses          []web.ComplaintResponse
// 	)

// 	// Menerapkan filter berdasarkan patientStatus jika query parameter diberikan
// 	query := "user_id = ?"
// 	if patientStatus != "" {
// 		query += " AND patient_status = ?"
// 	}

// 	// Mendapatkan semua transaksi dokter untuk pengguna dengan atau tanpa filter patientStatus
// 	if err := configs.DB.Find(&doctorTransactions, query, dokterID, patientStatus).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to retrieve doctor transactions"))
// 	}

// 	for _, transaction := range doctorTransactions {
// 		if err := configs.DB.Find(&doctor, "id=?", transaction.DoctorID).Error; err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to retrieve doctor data"))
// 		}

// 		// Mengambil pengguna yang mengajukan keluhan berdasarkan status pasien
// 		if transaction.PatientStatus != "" {
// 			var complainingUser schema.User

// 			if err := configs.DB.Where("id = ?", transaction.UserID).First(&complainingUser).Error; err != nil {
// 				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to retrieve complaining user"))
// 			}

// 			response := response.ConvertToComplaintResponse(complainingUser, transaction, doctor)
// 			responses = append(responses, response)
// 		}
// 	}
// 	return c.JSON(http.StatusOK, helper.SuccessResponse("Data successfully retrieved", responses))
// }

// func UpdateComplaintDataController(c echo.Context) error {
// 	dokterID, ok := c.Get("userID").(int)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil ID Dokter"))
// 	}

// 	// Mendapatkan ID transaksi dari parameter query
// 	id, err := strconv.Atoi(c.QueryParam("id"))
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
