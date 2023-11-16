package controllers

import (
	"healthcare/configs"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateDoctorTransactionController membuat transaksi dokter baru.
func CreateDoctorTransaction(c echo.Context) error {
	// Bind request data ke struct CreateDoctorTransactionRequest
	var dtRequest web.CreateDoctorTransactionRequest
	if err := c.Bind(&dtRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Data"))
	}

	// Convert request data ke struct DoctorTransaction
	doctorTransaction := request.ConvertToCreateDTRequest(dtRequest)

	// Simpan transaksi dokter ke dalam database
	if err := configs.DB.Create(doctorTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create Doctor Transaction"))
	}

	// Convert transaksi dokter ke dalam format respons yang diinginkan
	dtResponse := response.ConvertToCreateDTResponse(doctorTransaction)

	// Return respons JSON dengan status HTTP 201 (Created) dan data transaksi dokter yang baru dibuat
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Doctor Transaction Created Successfully", dtResponse))
}
