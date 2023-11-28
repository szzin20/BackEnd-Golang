package controllers

import (
	"healthcare/configs"
	"healthcare/middlewares"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/response"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

// Admin Login
func LoginAdminController(c echo.Context) error {
	var loginRequest web.AdminLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input login data"))
	}

	if err := helper.ValidateStruct(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	var admin schema.Admin
	if err := configs.DB.Where("email = ?", loginRequest.Email).First(&admin).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("email not registered"))
	}

	if err := configs.DB.Where("email = ? AND password = ?", loginRequest.Email, loginRequest.Password).First(&admin).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("incorrect email or password"))
	}

	token, err := middlewares.GenerateToken(admin.ID, admin.Email, admin.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to generate jwt"))
	}
	adminLoginResponse := response.ConvertToAdminLoginResponse(admin)
	adminLoginResponse.Token = token

	return c.JSON(http.StatusOK, helper.SuccessResponse("login successful", adminLoginResponse))
}

// Update Admin
func UpdateAdminController(c echo.Context) error {
	userID := c.Get("userID")

	var updatedAdmin web.AdminUpdateRequest

	if err := c.Bind(&updatedAdmin); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input update data"))
	}

	var existingAdmin schema.Admin
	result := configs.DB.First(&existingAdmin, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve admin"))
	}

	configs.DB.Model(&existingAdmin).Updates(updatedAdmin)

	response := response.ConvertToAdminUpdateResponse(&existingAdmin)

	return c.JSON(http.StatusOK, helper.SuccessResponse("admin updated data successful", response))
}

// UpdatePaymentStatusByAdminController updates payment status by admin
func UpdatePaymentStatusByAdminController(c echo.Context) error {
	// Parse transaction ID from the request parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid transaction id"))
	}

	// Retrieve the existing transaction from the database
	var existingTransaction schema.DoctorTransaction
	result := configs.DB.First(&existingTransaction, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve transaction"))
	}

	// Bind the updated payment status from the request body
	var updateRequest struct {
		PaymentStatus string `json:"payment_status" validate:"required,oneof=pending success cancelled"`
	}

	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input update data"))
	}

	// Validate the updated payment status
	if err := helper.ValidateStruct(updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("payment status updated successfully", nil))
}
