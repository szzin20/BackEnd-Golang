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

	// userID, ok := c.Get("userID").(int)
	// if !ok {
	// 	return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid User ID"))
	// }
	

	var dtRequest web.CreateDoctorTransactionRequest

	if err := c.Bind(&dtRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input transaction Data"))
	}

	dtr := request.ConvertToCreateDTRequest(dtRequest)

	if err := configs.DB.Create(&dtr).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create Transaction"))
	}

	dtResponse := response.ConvertToCreateDTResponse(dtr)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Transaction Created Successfully", dtResponse))
}
