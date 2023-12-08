package controllers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/helper/constanta"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

// Create Checkout By User
func CreateCheckoutController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid user id"))
	}

	medicinetransactionID, _ := strconv.Atoi(c.QueryParam("medicine_transaction_id"))

	var checkout web.CheckoutRequest

	if err := c.Bind(&checkout); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidBody))
	}

	if err := helper.ValidateStruct(checkout); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Upload files
	err := c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	file, fileHeader, err := c.Request().FormFile("payment_confirmation")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrImageFileRequired))
	}
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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidImageFormat))
	}

	imageURL, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error upload image to cloud storage"))
	}

	checkout.PaymentConfirmation = imageURL

	var medicineTransaction schema.MedicineTransaction
	if err := configs.DB.Where("id = ? AND user_id = ?", medicinetransactionID, userID).First(&medicineTransaction).Error; err != nil {
		return c.JSON(http.StatusForbidden, helper.ErrorResponse("permission denied"))
	}

	var existingCheckout schema.Checkout
	if err := configs.DB.Where("medicine_transaction_id = ?", medicinetransactionID).First(&existingCheckout).Error; err == nil {
		if existingCheckout.PaymentStatus == "pending" {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("checkout is still pending"))
		}
		if existingCheckout.PaymentStatus == "success" {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("checkout is success"))
		}
	}

	checkoutRequest := request.ConvertToCheckoutRequest(checkout)
	if err := configs.DB.Create(&checkoutRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionCreated+"checkout"))
	}

	if err := configs.DB.Model(&medicineTransaction).Update("status_transaction", "sudah dibayar").Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionUpdated+"medicine transaction status"))
	}

	var created schema.Checkout
	if err := configs.DB.Preload("MedicineTransaction.MedicineDetails").First(&created, checkoutRequest.ID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionGet+"created checkout"))
	}

	response := response.ConvertToGetCheckoutResponse(&created)

	return c.JSON(http.StatusCreated, helper.SuccessResponse(constanta.SuccessActionCreated+"checkout", response))
}

// Get Checkout By User
func GetUserCheckoutController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid user id"))
	}

	params := c.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("limit"+constanta.ErrQueryParamRequired))
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("offset"+constanta.ErrQueryParamRequired))
	}

	paymentStatus := params.Get("payment_status")

	var checkouts []schema.Checkout

	checkouts, total, err := GetUserAllCheckoutPagination(offset, limit, paymentStatus, userID, []schema.Checkout{})

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("checkouts "+constanta.ErrNotFound))
		}
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
	}

	pagination := helper.Pagination(offset, limit, total)

	response := response.ConvertToGetAllCheckoutResponse(checkouts)

	return c.JSON(http.StatusOK, helper.PaginationResponse(constanta.SuccessActionGet+"checkouts", response, pagination))
}

func GetUserAllCheckoutPagination(offset int, limit int, paymentStatus string, userID int, queryInput []schema.Checkout) ([]schema.Checkout, int64, error) {
	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	queryAll := queryInput
	var total int64

	query := configs.DB.Model(&queryAll).
		Joins("JOIN medicine_transactions ON checkouts.medicine_transaction_id = medicine_transactions.id").
		Where("medicine_transactions.user_id = ?", uint(userID))

	if paymentStatus != "" {
		query = query.Where("checkouts.payment_status = ?", paymentStatus)
	}

	query = query.Preload("MedicineTransaction.MedicineDetails").
		Order("checkouts.created_at DESC")

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

func GetUserCheckoutByIDController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid user id"))
	}

	checkoutID, err := strconv.Atoi(c.Param("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid checkout id"))
	}

	var checkout schema.Checkout
	result := configs.DB.
		Joins("JOIN medicine_transactions ON checkouts.medicine_transaction_id = medicine_transactions.id").
		Preload("MedicineTransaction.MedicineDetails").
		Where("medicine_transactions.user_id = ? AND checkouts.id = ?", userID, checkoutID).
		First(&checkout)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound+" checkout"))
	}

	response := response.ConvertToGetCheckoutResponse(&checkout)

	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionGet+"checkout", response))
}

// UpdateCheckoutController By Admin
func UpdateCheckoutController(c echo.Context) error {
	checkoutID, err := strconv.Atoi(c.Param("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid checkout id"))
	}

	var existingCheckout schema.Checkout

	result := configs.DB.Preload("MedicineTransaction.MedicineDetails").First(&existingCheckout, checkoutID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionGet+"checkout id"))
	}

	var updatedCheckoutRequest web.CheckoutUpdate
	if err := c.Bind(&updatedCheckoutRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidBody))
	}

	updatedCheckout := request.ConvertToCheckoutUpdate(updatedCheckoutRequest)

	if updatedCheckout.PaymentStatus == "success" {
		medicineTransaction := existingCheckout.MedicineTransaction
		if err := reduceStock(medicineTransaction.MedicineDetails); err != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
	}

	if updatedCheckout.PaymentStatus == "cancelled" {
		if err := configs.DB.Table("medicine_transactions").
			Where("id = ?", existingCheckout.MedicineTransactionID).
			Update("status_transaction", "belum dibayar").Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionUpdated+"medicine transaction status"))
		}
	}

	result = configs.DB.Table("checkouts").Where("id = ?", checkoutID).Updates(updatedCheckout)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionUpdated+"checkout"))
	}

	var updated schema.Checkout
	if err := configs.DB.Preload("MedicineTransaction.MedicineDetails").First(&updated, checkoutID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionGet+"updated checkout"))
	}

	response := response.ConvertToGetCheckoutResponse(&updated)

	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionUpdated+"checkout", response))
}

func reduceStock(medicineDetails []schema.MedicineDetails) error {
	for _, md := range medicineDetails {
		medicine := schema.Medicine{}
		if err := configs.DB.First(&medicine, md.MedicineID).Error; err != nil {
			return errors.New("invalid medicine id")
		}

		if medicine.Stock < md.Quantity {
			return errors.New("insufficient stock")
		}

		// Reduce stock in the Medicine table
		newStock := medicine.Stock - md.Quantity
		if err := configs.DB.Model(&medicine).Update("stock", newStock).Error; err != nil {
			return errors.New(constanta.ErrActionUpdated + "medicine stock")
		}
	}
	return nil
}

// Get Checkout By Admin
func GetAdminCheckoutController(c echo.Context) error {

	params := c.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("limit"+constanta.ErrQueryParamRequired))
	}

	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("offset"+constanta.ErrQueryParamRequired))
	}

	paymentStatus := params.Get("payment_status")
	userID, err := strconv.Atoi(params.Get("user_id"))

	var checkouts []schema.Checkout

	checkouts, total, err := GetAdminAllCheckoutPagination(offset, limit, userID, paymentStatus, []schema.Checkout{})

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("checkouts "+constanta.ErrNotFound))
		}
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
	}

	pagination := helper.Pagination(offset, limit, total)

	response := response.ConvertToGetAllCheckoutResponse(checkouts)

	return c.JSON(http.StatusOK, helper.PaginationResponse(constanta.SuccessActionGet+"checkouts", response, pagination))
}

func GetAdminAllCheckoutPagination(offset, limit, userID int, paymentStatus string, queryInput []schema.Checkout) ([]schema.Checkout, int64, error) {
	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	queryAll := queryInput
	var total int64

	query := configs.DB.Model(&queryAll).
		Joins("JOIN medicine_transactions ON checkouts.medicine_transaction_id = medicine_transactions.id").
		Joins("JOIN users ON medicine_transactions.user_id = users.id")

	if userID != 0 {
		query = query.Where("users.id = ?", userID)
	}

	if paymentStatus != "" {
		query = query.Where("checkouts.payment_status = ?", paymentStatus)
	}

	query = query.Preload("MedicineTransaction.MedicineDetails").
		Order("checkouts.created_at DESC")

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

func GetAdminCheckoutByIDController(c echo.Context) error {

	checkoutID, err := strconv.Atoi(c.Param("checkout_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid checkout id"))
	}

	var checkout schema.Checkout
	result := configs.DB.
		Preload("MedicineTransaction.MedicineDetails").
		Joins("JOIN medicine_transactions ON checkouts.medicine_transaction_id = medicine_transactions.id").
		Where("checkouts.id = ?", checkoutID).
		First(&checkout)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrActionGet+"checkout"))
	}

	response := response.ConvertToGetCheckoutResponse(&checkout)

	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionGet+"checkout", response))
}
