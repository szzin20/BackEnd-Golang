package controllers

import (
	"fmt"
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// User Create Roomchat
func CreateRoomchatController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	transactionID, err := strconv.Atoi(c.Param("transaction_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid transaction id"))
	}

	var existingRoomchat schema.Roomchat

	if err := configs.DB.First(&existingRoomchat, "transaction_id = ?", transactionID).Error; err == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("roomchat for this transaction id already exists"))
	}

	var doctorTransaction schema.DoctorTransaction

	if err := configs.DB.First(&doctorTransaction, "user_id = ? AND id = ? AND payment_status = 'success'", userID, transactionID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
	}

	roomchat := request.CreateRoomchatRequest(uint(transactionID))

	if err := configs.DB.Create(&roomchat).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to create roomchat"))
	}

	response := response.ConvertToCreateRoomchatResponse(&roomchat)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("roomchat created successful", response))
}

// User Get Roomchat by ID
func GetUserRoomchatController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	roomchatID, err := strconv.Atoi(c.Param("roomchat_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid roomchat id"))
	}

	var existingRoomchat schema.Roomchat
	if err := configs.DB.Debug().First(&existingRoomchat, "id = ?", roomchatID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve roomchat data"))
	}

	var doctortransaction schema.DoctorTransaction
	if err := configs.DB.Where("user_id = ? AND id = ?", userID, existingRoomchat.TransactionID).First(&doctortransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
	}

	var roomchat schema.Roomchat
	if err := configs.DB.Where("id = ?", roomchatID).Preload("Message").First(&roomchat).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve message data"))
	}

	var doctor schema.Doctor
	if err := configs.DB.Where("id = ?", doctortransaction.DoctorID).First(&doctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor data"))
	}

	response := response.ConvertToRoomchatResponse(&roomchat, &doctor)
	fmt.Println(response)

	return c.JSON(http.StatusOK, response)
}


















// // User  Get All Roomchats or Get Roomchat Details by Transaction ID
// func GetUserRoomchatsController(c echo.Context) error {

// 	userID, ok := c.Get("userID").(int)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
// 	}

// 	roomchatID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid roomchat id"))
// 	}

// 	var existingRoomchat schema.Roomchat
// 	if err := configs.DB.First(&existingRoomchat, "id = ?", roomchatID).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve roomchat id"))
// 	}

// 	var doctorTransaction schema.DoctorTransaction
// 	if err := configs.DB.First(&doctorTransaction, "id = ?", existingRoomchat.TransactionID).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
// 	}

// 	if doctorTransaction.UserID != uint(userID) {
// 		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("you are not authorized to access this roomchat"))
// 	}

// 	if roomchatID == 0 {

// 		var doctorTransaction []schema.DoctorTransaction

// 		err := configs.DB.Find(&doctorTransaction, "user_id=?", userID).Error
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
// 		}

// 		var responses []web.DoctorTransactionDetailsResponse
// 		for i, doctor_id := range doctorTransaction {

// 			var doctor schema.Doctor
// 			err := configs.DB.Find(&doctor, "id=?", doctor_id.DoctorID).Error
// 			if err != nil {
// 				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor data"))
// 			}

// 			if len(doctorTransaction) == 0 {
// 				return c.JSON(http.StatusNotFound, helper.ErrorResponse("empty doctor transaction details data"))
// 			}

// 			responses = append(responses, response.ConvertToGetAllDoctorTransactionDetailsResponse(doctorTransaction[i], doctor))
// 		}

// 		return c.JSON(http.StatusOK, helper.SuccessResponse("doctor transaction details data successfully retrieved", responses))

// 	}

// 	var doctortransaction schema.DoctorTransaction
// 	if err := configs.DB.Where("user_id = ? AND id = ?", userID, transactionID).First(&doctortransaction).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
// 	}

// 	var roomchat schema.Roomchat
// 	if err := configs.DB.Where("transaction_id = ?", transactionID).Preload("Message").First(&roomchat).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve message data"))
// 	}

// 	response := response.ConvertToRoomchatResponse(&roomchat)

// 	return c.JSON(http.StatusOK, helper.SuccessResponse("roomchat data successfully retrieved", response))
// }

// // User Get Roomchat by ID
// func GetUserRoomchatsController(c echo.Context) error {

// 	userID, ok := c.Get("userID").(int)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
// 	}

// 	roomchatID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid roomchat id"))
// 	}

// 	var existingRoomchat schema.Roomchat
// 	if err := configs.DB.First(&existingRoomchat, "id = ?", roomchatID).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve roomchat id"))
// 	}

// 	var existingRoomchat schema.Roomchat

// 	if err := configs.DB.First(&existingRoomchat, "transaction_id = ?", transactionID).Error; err == nil {
// 		return c.JSON(http.StatusConflict, helper.ErrorResponse("roomchat for this transaction id already exists"))
// 	}

// 	var doctorTransaction schema.DoctorTransaction

// 	if err := configs.DB.First(&doctorTransaction, "user_id = ? AND id = ? AND payment_status = 'success'", userID, transactionID).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
// 	}

// 	roomchat := request.CreateRoomchatRequest(uint(transactionID))

// 	if err := configs.DB.Create(&roomchat).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to create roomchat"))
// 	}

// 	response := response.ConvertToCreateRoomchatResponse(&roomchat)

// 	return c.JSON(http.StatusCreated, helper.SuccessResponse("roomchat created successful", response))

// var existingRoomchat schema.Roomchat
// if err := configs.DB.First(&existingRoomchat, "id = ?", roomchatID).Error; err != nil {
// 	return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve roomchat id"))
// }

// var doctorTransaction schema.DoctorTransaction
// if err := configs.DB.First(&doctorTransaction, "id = ?", existingRoomchat.TransactionID).Error; err != nil {
// 	return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
// }

// if doctorTransaction.UserID != uint(userID) {
// 	return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("you are not authorized to access this roomchat"))
// }

// var doctortransaction schema.DoctorTransaction
// if err := configs.DB.Where("user_id = ? AND id = ?", userID, transactionID).First(&doctortransaction).Error; err != nil {
// 	return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
// }

// var roomchat schema.Roomchat
// if err := configs.DB.Where("transaction_id = ?", transactionID).Preload("Message").First(&roomchat).Error; err != nil {
// 	return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve message data"))
// }

// response := response.ConvertToRoomchatResponse(&roomchat)

// 	// return c.JSON(http.StatusOK, helper.SuccessResponse("roomchat data successfully retrieved", response))
// }

// // User  Get All Consultation or Get Roomchat Details by Transaction ID
// func GetUserRoomchatController(c echo.Context) error {

// 	userID, ok := c.Get("userID").(int)
// 	if !ok {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
// 	}

// 	transactionID, _ := strconv.Atoi(c.QueryParam("transaction_id"))

// 	if transactionID == 0 {

// 		var doctorTransaction []schema.DoctorTransaction

// 		err := configs.DB.Find(&doctorTransaction, "user_id=?", userID).Error
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
// 		}

// 		var responses []web.DoctorTransactionDetailsResponse
// 		for i, doctor_id := range doctorTransaction {

// 			var doctor schema.Doctor
// 			err := configs.DB.Find(&doctor, "id=?", doctor_id.DoctorID).Error
// 			if err != nil {
// 				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor data"))
// 			}

// 			if len(doctorTransaction) == 0 {
// 				return c.JSON(http.StatusNotFound, helper.ErrorResponse("empty doctor transaction details data"))
// 			}

// 			responses = append(responses, response.ConvertToGetAllDoctorTransactionDetailsResponse(doctorTransaction[i], doctor))
// 		}

// 		return c.JSON(http.StatusOK, helper.SuccessResponse("doctor transaction details data successfully retrieved", responses))

// 	}

// 	var doctortransaction schema.DoctorTransaction
// 	if err := configs.DB.Where("user_id = ? AND id = ?", userID, transactionID).First(&doctortransaction).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
// 	}

// 	var roomchat schema.Roomchat
// 	if err := configs.DB.Where("transaction_id = ?", transactionID).Preload("Message").First(&roomchat).Error; err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve message data"))
// 	}

// 	response := response.ConvertToRoomchatResponse(&roomchat)

// 	return c.JSON(http.StatusOK, helper.SuccessResponse("roomchat data successfully retrieved", response))
// }
