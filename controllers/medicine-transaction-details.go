package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"
	"path/filepath"
	"strconv"
)

func CreateCheckoutController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid user ID"))
	}

	medicinetransactionID, _ := strconv.Atoi(c.QueryParam("medicine_transaction_id"))

	var checkout web.CheckoutRequest

	if err := c.Bind(&checkout); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Checkout Data"))
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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Image File is Required"))
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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid image file format. Supported formats: jpg, jpeg, png"))
	}

	imageURL, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error upload image to Cloud Storage"))
	}

	checkout.PaymentConfirmation = imageURL

	var medicineTransaction schema.MedicineTransaction
	if err := configs.DB.Where("id = ? AND user_id = ?", medicinetransactionID, userID).First(&medicineTransaction).Error; err != nil {
		return c.JSON(http.StatusForbidden, helper.ErrorResponse("You do not have permission to create a checkout for this MedicineTransaction"))
	}

	var existingCheckout schema.Checkout
	if err := configs.DB.Where("medicine_transaction_id = ?", medicinetransactionID).First(&existingCheckout).Error; err == nil {
		if existingCheckout.PaymentStatus == "pending" || existingCheckout.PaymentStatus == "success" {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Checkout already exists with status: "+existingCheckout.PaymentStatus))
		}
	}

	checkoutRequest := request.ConvertToCheckoutRequest(checkout)
	if err := configs.DB.Create(&checkoutRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create Checkout"))
	}

	if err := configs.DB.Model(&medicineTransaction).Update("status_transaction", "sudah dibayar").Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to update MedicineTransaction status"))
	}

	response := response.ConvertToCheckoutResponse(checkoutRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Checkout Created Successfully", response))
}

// Get Checkout
func GetCheckoutController(c echo.Context) error {
	// Extract user ID from the context
	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid user ID"))
	}

	checkoutIDStr := c.QueryParam("id")
	paymentStatus := c.QueryParam("payment_status")

	var checkouts []schema.Checkout
	var err error

	// Join MedicineTransaction to get the user ID
	query := configs.DB.
		Joins("JOIN medicine_transactions ON checkouts.medicine_transaction_id = medicine_transactions.id").
		Where("medicine_transactions.user_id = ?", userID)

	if checkoutIDStr != "" {
		checkoutID, err := strconv.Atoi(checkoutIDStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Checkout ID"))
		}

		query = query.Where("checkouts.id = ?", checkoutID)
	} else if paymentStatus != "" {
		query = query.Where("checkouts.payment_status = ?", paymentStatus)
	}

	err = query.Find(&checkouts).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Checkout Data"))
	}

	if len(checkouts) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Checkout Data"))
	}

	response := response.ConvertToCheckoutListResponse(checkouts)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Checkout Data Successfully Retrieved", response))
}

// UpdateCheckoutController function
func UpdateCheckoutController(c echo.Context) error {
	checkoutID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Checkout ID"))
	}

	var existingCheckout schema.Checkout

	result := configs.DB.Preload("MedicineTransaction.MedicineDetails").First(&existingCheckout, checkoutID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Checkout ID"))
	}

	medicineTransaction := existingCheckout.MedicineTransaction
	if err := reduceStock(medicineTransaction.MedicineDetails); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	var updatedCheckoutRequest web.CheckoutUpdate
	if err := c.Bind(&updatedCheckoutRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Checkout Data"))
	}

	updatedCheckout := request.ConvertToCheckoutUpdate(updatedCheckoutRequest)

	// Update the existing Checkout
	result = configs.DB.Table("checkouts").Where("id = ?", checkoutID).Updates(updatedCheckout)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Update Checkout"))
	}

	var hasil schema.Checkout

	hasil = configs.DB.Preload("MedicineTransaction.MedicineDetails").First(&existingCheckout, checkoutID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Checkout ID"))
	}

	response := response.ConvertToCheckoutResponse(hasil)
	return c.JSON(http.StatusOK, helper.SuccessResponse("Checkout Updated Successfully", response))
}

func reduceStock(medicineDetails []schema.MedicineDetails) error {
	for _, md := range medicineDetails {
		medicine := schema.Medicine{}
		if err := configs.DB.First(&medicine, md.MedicineID).Error; err != nil {
			return errors.New("Invalid Medicine ID")
		}

		if medicine.Stock < md.Quantity {
			return errors.New("Insufficient stock")
		}

		// Reduce stock in the Medicine table
		newStock := medicine.Stock - md.Quantity
		if err := configs.DB.Model(&medicine).Update("stock", newStock).Error; err != nil {
			return errors.New("Failed to update Medicine stock")
		}
	}
	return nil
}
