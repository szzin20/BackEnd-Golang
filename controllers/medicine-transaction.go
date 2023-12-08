package controllers

import (
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
	"strconv"
	"strings"
)

func CreateMedicineTransaction(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid user id"))
	}

	var medicineTransactionRequest web.MedicineTransactionRequest

	if err := c.Bind(&medicineTransactionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidBody))
	}

	for _, md := range medicineTransactionRequest.MedicineDetails {
		if md.MedicineID == 0 {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("medicine id cannot be empty"))
		}
		if md.Quantity == 0 {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("quantity cannot be empty"))
		}
	}

	if err := helper.ValidateStruct(medicineTransactionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	medicineTransaction := request.ConvertToMedicineTransactionRequest(medicineTransactionRequest, uint(userID))

	totalPrice := 0
	for i, md := range medicineTransaction.MedicineDetails {
		medicine := schema.Medicine{}
		if err := configs.DB.First(&medicine, md.MedicineID).Error; err != nil {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("medicine id "+constanta.ErrNotFound))
		}

		if medicine.Stock < md.Quantity {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("insufficient stock"))
		}

		medicineTransaction.MedicineDetails[i].TotalPriceMedicine = md.Quantity * medicine.Price

		totalPrice += medicineTransaction.MedicineDetails[i].TotalPriceMedicine
	}

	medicineTransaction.TotalPrice = totalPrice

	if err := configs.DB.Create(&medicineTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionCreated+"medicine transaction"))
	}

	response := response.ConvertToMedicineTransactionResponse(medicineTransaction)

	return c.JSON(http.StatusCreated, helper.SuccessResponse(constanta.SuccessActionCreated+"medicine transaction", response))
}

// Get Medicine Transaction by ID
func GetMedicineTransactionByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("medtrans_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid medicine transaction id"))
	}

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid user id"))
	}

	var medicineTransaction schema.MedicineTransaction

	if err := configs.DB.Preload("MedicineDetails").Where("user_id = ?", userID).First(&medicineTransaction, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound))
	}

	response := response.ConvertToMedicineTransactionResponse(&medicineTransaction)

	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionGet+"medicine transaction", response))
}

func DeleteMedicineTransactionController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid user id"))
	}

	medicineTransactionID, err := strconv.Atoi(c.Param("medtrans_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidIDParam))
	}

	var medicineTransaction schema.MedicineTransaction
	if err := configs.DB.Where("id = ? AND user_id = ?", medicineTransactionID, userID).First(&medicineTransaction).Error; err != nil {
		return c.JSON(http.StatusForbidden, helper.ErrorResponse("permission denied"))
	}

	if err := configs.DB.Delete(&medicineTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionDeleted+"medicine transaction"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionDeleted+"medicine transaction", nil))
}

func GetMedicineTransactionController(c echo.Context) error {
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

	status := params.Get("status_transaction")

	var medicinesTransaction []schema.MedicineTransaction

	medicinesTransaction, total, err := GetMedicineTransactionPagination(userID, offset, limit, status)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("medicines transaction "+constanta.ErrNotFound))
		}
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
	}

	pagination := helper.Pagination(offset, limit, total)

	response := response.ConvertToMedicineTransactionListResponse(medicinesTransaction)

	return c.JSON(http.StatusOK, helper.PaginationResponse(constanta.SuccessActionGet+"medicine transactions", response, pagination))
}

func GetMedicineTransactionPagination(userID int, offset int, limit int, status string) ([]schema.MedicineTransaction, int64, error) {
	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	var medicineTransactions []schema.MedicineTransaction
	var total int64
	query := configs.DB.Model(&medicineTransactions).Where("user_id = ?", userID)

	if status != "" {
		query = query.Where("status_transaction = ?", status)
	}

	query.Preload("MedicineDetails").
		Order("created_at DESC").
		Find(&medicineTransactions).Count(&total)

	query = query.Limit(limit).Offset(offset)

	result := query.Find(&medicineTransactions)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if offset >= int(total) {
		return nil, 0, fmt.Errorf("not found")
	}

	return medicineTransactions, total, nil
}
