package controllers

import (
	"errors"
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllDoctorComplaints(c echo.Context) error {
	// Ekstrak DoctorID dari konteks
	doctorID, ok := c.Get("DoctorID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil DoctorID"))
	}

	// Ambil semua keluhan yang terkait dengan pasien-pasien yang ditangani oleh dokter tersebut
	var doctorComplaint []schema.Complaint
	if err := configs.DB.
		Preload("DoctorTransaction").
		Where("doctor_id = ?", doctorID).
		Find(&doctorComplaint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data keluhan: "+err.Error()))
	}

	responseData := response.ConvertToGetAllComplaintsResponse(doctorComplaint)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Daftar Keluhan Dokter", responseData))
}

func GetDoctorComplaintsByStatus(c echo.Context) error {
	doctorID, ok := c.Get("DoctorID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil DoctorID"))
	}

	// Ambil status dari parameter query
	status := c.QueryParam("status")

	var doctor schema.Doctor
	if err := configs.DB.First(&doctor, doctorID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil Profil Dokter"))
	}

	// Ambil semua pasien untuk dokter tersebut
	var patients []schema.User
	if err := configs.DB.Model(&doctor).Association("Patients").Find(&patients); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data pasien"))
	}

	// Dapatkan parameter status dari URL sebagai boolean
	bStatus, err := strconv.ParseBool(status)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Parameter status tidak valid"))
	}

	// Validasi status menggunakan fungsi bantuan
	if err := helper.ValidateStruct(bStatus); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Dapatkan semua keluhan terkait pasien-pasien tersebut berdasarkan status yang diberikan
	var complaints []schema.Complaint
	for _, patient := range patients {
		var patientComplaints []schema.Complaint
		if err := configs.DB.Where("patient_id = ? AND status = ?", patient.ID, bStatus).Find(&patientComplaints).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data keluhan"))
		}
		complaints = append(complaints, patientComplaints...)
	}

	// Periksa jika tidak ada keluhan
	if len(complaints) == 0 {
		return c.JSON(http.StatusOK, helper.SuccessResponse("Tidak ada keluhan ditemukan berdasarkan status yang diberikan", nil))
	}

	// Konversi data keluhan ke format respons yang diinginkan
	var convertedData []web.ComplaintsUpdateResponse
	for _, complaint := range complaints {
		converted := response.ConvertToGetComplaintsResponse(&complaint)
		convertedData = append(convertedData, converted)
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Daftar Keluhan Dokter Berdasarkan Status", convertedData))
}

func UpdateDoctorComplaintStatus(c echo.Context) error {
	// Ekstrak DoctorID dari konteks
	doctorID, ok := c.Get("DoctorID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil DoctorID"))
	}

	// Ambil complaintID dari parameter 
	complaintID := c.Param("complaintID")

	// Bind data permintaan pembaruan dari body permintaan
	var updateRequest web.ComplaintsUpdateRequest
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Gagal memproses permintaan"))
	}

	// Validasi data permintaan
	if err := c.Validate(&updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Ambil keluhan dari database berdasarkan complaintID dan doctorID
	var complaint schema.Complaint
	if err := configs.DB.Preload("DoctorTransaction").Where("id = ? AND doctor_transaction.doctor_id = ?", complaintID, doctorID).First(&complaint).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("Keluhan tidak ditemukan"))
		}
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data keluhan"))
	}

	// Perbarui status keluhan
	complaint.Status = updateRequest.Status

	// Simpan keluhan yang diperbarui ke database
	if err := configs.DB.Save(&complaint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal memperbarui status keluhan"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Status Keluhan Diperbarui", map[string]interface{}{"keluhan": &complaint}))
}

