package controllers

import (
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
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
	if err := configs.DB.First(&existingRoomchat, "id = ?", roomchatID).Error; err != nil {
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

	return c.JSON(http.StatusOK, helper.SuccessResponse("roomchat data successfully retrieved", response))
}

// Doctor Get Roomchat by ID
func GetDoctorRoomchatController(c echo.Context) error {

	doctorID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid doctor id"))
	}

	roomchatID, err := strconv.Atoi(c.Param("roomchat_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid roomchat id"))
	}

	var existingRoomchat schema.Roomchat
	if err := configs.DB.First(&existingRoomchat, "id = ?", roomchatID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve roomchat data"))
	}

	var doctortransaction schema.DoctorTransaction
	if err := configs.DB.Where("doctor_id = ? AND id = ?", doctorID, existingRoomchat.TransactionID).First(&doctortransaction).Error; err != nil {
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

	return c.JSON(http.StatusOK, helper.SuccessResponse("roomchat data successfully retrieved", response))
}

// Doctor Get All Roomchats
func GetAllDoctorRoomchatController(c echo.Context) error {

	doctorID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid doctor id"))
	}

	var existingDoctorTransactions []schema.DoctorTransaction
	if err := configs.DB.Find(&existingDoctorTransactions, "doctor_id = ?", doctorID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
	}

	var responses []web.RoomchatListResponse

	for _, doctorTransaction := range existingDoctorTransactions {

		userID := doctorTransaction.UserID

		var user schema.User
		if err := configs.DB.First(&user, "id = ?", userID).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve user data"))
		}

		var existingRoomchat schema.Roomchat
		if err := configs.DB.Preload("Message").Where("transaction_id = ?", doctorTransaction.ID).Find(&existingRoomchat).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve roomchat data"))
		}

		var LastMessage schema.Message
		lenghtMessage := len(existingRoomchat.Message) - 1
		if lenghtMessage >= 0 {
			LastMessage = existingRoomchat.Message[lenghtMessage]
		}

		response := response.ConvertToGetAllRoomchats(user, existingRoomchat, LastMessage)

		responses = append(responses, response)

	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("roomchat data successfully retrieved", responses))
}
