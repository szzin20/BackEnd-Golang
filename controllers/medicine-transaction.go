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
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	var medicineTransactionRequest web.MedicineTransactionRequest

	if err := c.Bind(&medicineTransactionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Medicine Transaction Data"))
	}

	if err := helper.ValidateStruct(medicineTransactionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	medicineTransaction := request.ConvertToMedicineTransactionRequest(medicineTransactionRequest, uint(userID))

	totalPrice := 0
	for i, md := range medicineTransaction.MedicineDetails {
		medicine := schema.Medicine{}
		if err := configs.DB.First(&medicine, md.MedicineID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Medicine ID"))
		}

		if medicine.Stock < md.Quantity {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Insufficient stock"))
		}

		medicineTransaction.MedicineDetails[i].TotalPriceMedicine = md.Quantity * medicine.Price

		totalPrice += medicineTransaction.MedicineDetails[i].TotalPriceMedicine
	}

	medicineTransaction.TotalPrice = totalPrice

	if err := configs.DB.Create(&medicineTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create Medicine Transaction"))
	}

	response := response.ConvertToMedicineTransactionResponse(medicineTransaction)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Medicine Transaction Created Successfully", response))
}

// Get Medicine Transaction
//func GetMedicineTransactionController(c echo.Context) error {
//	userID, ok := c.Get("userID").(int)
//	if !ok {
//		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
//	}
//
//	idStr := c.QueryParam("id")
//	status := c.QueryParam("status_transaction")
//
//	var medicineTransactions []schema.MedicineTransaction
//	var err error
//
//	if idStr != "" {
//		id, err := strconv.Atoi(idStr)
//		if err != nil {
//			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Medicine Transaction ID"))
//		}
//
//		err = configs.DB.Preload("MedicineDetails").Where("user_id = ? AND id = ?", userID, id).Find(&medicineTransactions).Error
//	} else if status != "" {
//		err = configs.DB.Preload("MedicineDetails").Where("user_id = ? AND status_transaction = ?", userID, status).Find(&medicineTransactions).Error
//	} else {
//		err = configs.DB.Preload("MedicineDetails").Where("user_id = ?", userID).Find(&medicineTransactions).Error
//	}
//
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicine Transactions Data"))
//	}
//
//	if len(medicineTransactions) == 0 {
//		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Medicine Transactions Data"))
//	}
//
//	response := response.ConvertToMedicineTransactionListResponse(medicineTransactions)
//
//	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Transactions Data Successfully Retrieved", response))
//}

// Get Medicine Transaction by ID
func GetMedicineTransactionByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Medicine Transaction ID"))
	}

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	var medicineTransaction schema.MedicineTransaction

	if err := configs.DB.Preload("MedicineDetails").Where("user_id = ?", userID).First(&medicineTransaction, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicine Transaction Data"))
	}

	response := response.ConvertToMedicineTransactionResponse(&medicineTransaction)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Transaction Data Successfully Retrieved", response))
}

func DeleteMedicineTransactionController(c echo.Context) error {

	userID, ok := c.Get("userID").(uint)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid user ID"))
	}

	medicineTransactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid MedicineTransaction ID"))
	}

	var medicineTransaction schema.MedicineTransaction
	if err := configs.DB.Where("id = ? AND user_id = ?", medicineTransactionID, userID).First(&medicineTransaction).Error; err != nil {
		return c.JSON(http.StatusForbidden, helper.ErrorResponse("You do not have permission to delete this MedicineTransaction"))
	}

	if err := configs.DB.Delete(&medicineTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to delete MedicineTransaction"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("MedicineTransaction Deleted Successfully", nil))
}

func GetMedicineTransactionController(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	params := c.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("limit"+constanta.ErrQueryParamRequired))
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("offset"+constanta.ErrQueryParamRequired))
	}

	status := params.Get("status_transaction")

	var medicinesTransaction []schema.MedicineTransaction

	medicinesTransaction, total, err := GetMedicineTransactionPagination(userID, offset, limit, status)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("medicines "+constanta.ErrNotFound))
		}
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	pagination := helper.Pagination(offset, limit, total)

	response := response.ConvertToMedicineTransactionListResponse(medicinesTransaction)

	return c.JSON(http.StatusOK, helper.PaginationResponse("Medicine Transactions Data Successfully Retrieved", response, pagination))
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
