package controllers

import (
	"fmt"
	"healthcare/configs"
	"healthcare/middlewares"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/helper/constanta"
	"healthcare/utils/response"
	"net/http"
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllAdminsPagination(offset int, limit int, queryInput []schema.Admin) ([]schema.Admin, int64, error) {

	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	queryAll := queryInput
	var total int64

	query := configs.DB.Model(&queryAll)

	query.Find(&queryAll).Count(&total)

	query = query.Limit(limit).Offset(offset)

	result := query.Find(&queryAll)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if offset >= int(total) {
		return nil, 0, fmt.Errorf("not found")
	}

	return queryAll, total, nil
}

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

// Admin Update Payment Status 
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
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("limit"+constanta.ErrQueryParamRequired))
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("offset"+constanta.ErrQueryParamRequired))
	}

	var doctorTransactions []schema.DoctorTransaction
	var total int64

	// Fetch all transactions with payment_status IN ('success', 'pending', 'cancelled')
	query := configs.DB.Where("payment_status IN (?)", []string{"success", "pending", "cancelled"}).Find(&doctorTransactions)
	if query.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionGet+"doctor transactions"))
	}

	sort.Slice(doctorTransactions, func(i, j int) bool {
		order := map[string]int{"pending": 0, "success": 1, "cancelled": 2}
		return order[doctorTransactions[i].PaymentStatus] < order[doctorTransactions[j].PaymentStatus]
	})

	// Count total number of records
	total = int64(len(doctorTransactions))

	// Apply limit and offset to the result
	start := offset
	end := offset + limit
	if start > len(doctorTransactions) {
		start = len(doctorTransactions)
	}
	if end > len(doctorTransactions) {
		end = len(doctorTransactions)
	}
	doctorTransactions = doctorTransactions[start:end]

	if len(doctorTransactions) == 0 {
		return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.ErrNotFound+"doctor transactions", nil))
	}

	pagination := helper.Pagination(offset, limit, total)
	Responses := response.ConvertToAdminTransactionUsersResponse(doctorTransactions)

	return c.JSON(http.StatusOK, helper.PaginationResponse(constanta.SuccessActionCreated+"doctor transactions", Responses, pagination))
}

func GetDoctorTransactionByIDController(c echo.Context) error {

	var doctorTransaction schema.DoctorTransaction

	transactionID := c.QueryParam("transaction_id")

	if transactionID == "" {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Transaction ID is required"))
	}
	if err := configs.DB.Where("id = ?", transactionID).First(&doctorTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionGet+"doctor transaction by ID"))
	}

	if doctorTransaction.ID == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound+"doctor transaction by ID"))
	}

	response := response.ConvertToAdminTransactionUsersResponse([]schema.DoctorTransaction{doctorTransaction})
	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionCreated+"doctor transaction by ID", response))
}
