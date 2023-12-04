package controllers

import (
	"healthcare/configs"
	"healthcare/middlewares"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/helper/constanta"
	"healthcare/utils/response"
	"log"
	"strconv"
	"sort"
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

// Admin Update Payment Status and Send Notification to Doctor
func UpdatePaymentStatusByAdminController(c echo.Context) error {
	// Parse transaction ID from the request parameters
	transaction_id, err := strconv.Atoi(c.Param("transaction_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid transaction id"))
	}

	var existingData schema.DoctorTransaction
	results := configs.DB.First(&existingData, transaction_id)
	if results.Error != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound))
	}

	// Retrieve the existing transaction from the database
	var existingTransaction schema.DoctorTransaction
	result := configs.DB.First(&existingTransaction, transaction_id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve transaction"))
	}

	var updateRequest web.UpdatePaymentRequest
	if err := c.Bind(&updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidBody))
	}

	// Validate the updated payment status
	if err := helper.ValidateStruct(updateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	result = configs.DB.First(&existingTransaction, transaction_id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve transaction"))
	}

	configs.DB.Model(&existingTransaction).Updates(updateRequest)

	result = configs.DB.Model(&existingTransaction).Updates(updateRequest)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionUpdated+"payment status"))
	}


	if updateRequest.PaymentStatus == "success" {
		
		var doctor schema.Doctor
		result := configs.DB.First(&doctor, "id = ?", existingTransaction.DoctorID)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor data"))
		}

		// Send an email to the doctor
		err = helper.SendNotificationEmail(doctor.Email, doctor.Fullname, "complaints", "","","")
		log.Printf(doctor.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send verification email"))
		}
	
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionUpdated+"payment status", nil))
}

func GetAdminProfileController(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrInvalidIDParam))
	}

	var admin schema.Admin
	if err := configs.DB.First(&admin, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound))

	}

	response := response.ConvertToGetProfileAdminResponse(&admin)

	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionGet+"admin profile", response))
}

// Get all transactions by doctors
func GetAllDoctorsPaymentsByAdminsController(c echo.Context) error {
    var doctorTransactions []schema.DoctorTransaction

    // "success" 
    if err := configs.DB.Where("payment_status = ?", "success").Find(&doctorTransactions).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionGet + "doctor transactions"))
    }

    // "pending" 
    var pending []schema.DoctorTransaction
    if err := configs.DB.Where("payment_status = ?", "pending").Find(&pending).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionGet + "pending doctor transactions"))
    }

    // "cancelled" 
    var cancelled []schema.DoctorTransaction
    if err := configs.DB.Where("payment_status = ?", "cancelled").Find(&cancelled).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionGet + "cancelled doctor transactions"))
    }

    // Concatenate the results
    doctorTransactions = append(doctorTransactions, pending...)
    doctorTransactions = append(doctorTransactions, cancelled...)

    // Sort by custom order: pending, success, cancelled
    sort.Slice(doctorTransactions, func(i, j int) bool {
        order := map[string]int{"pending": 0, "success": 1, "cancelled": 2}
        return order[doctorTransactions[i].PaymentStatus] < order[doctorTransactions[j].PaymentStatus]
    })

    if len(doctorTransactions) == 0 {
        return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.ErrNotFound + "doctor transactions", nil))
    }

    Responses := response.ConvertToAdminTransactionUsersResponse(doctorTransactions)

    return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionCreated + "doctor transactions", Responses))
}




