package controllers

import (
	"fmt"
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/helper/constanta"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"

	"sort"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GetAllRoomchatPagination(doctorID int, offset int, limit int, queryInput []schema.DoctorTransaction) ([]schema.DoctorTransaction, int64, error) {

	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	queryAll := queryInput
	var total int64

	query := configs.DB.Model(&queryAll)

	query.Preload("Roomchat.Message").Where("doctor_id = ? AND payment_status = ?", doctorID, "success").Find(&queryAll).Count(&total)

	query = query.Limit(limit).Offset(offset)

	result := query.Preload("Roomchat.Message").Where("doctor_id = ? AND payment_status = ?", doctorID, "success").Find(&queryAll)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if offset >= int(total) {
		return nil, 0, fmt.Errorf("not found")
	}

	return queryAll, total, nil
}

// User Create Roomchat and Send Notification to Doctor
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

	if err == nil {

		var doctor schema.Doctor
		result := configs.DB.First(&doctor, "id = ?", doctorTransaction.DoctorID)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor data"))
		}

		err = helper.SendNotificationEmail(doctor.Email, doctor.Fullname, "complaints", "", "", "",false)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send verification email"))
		}

	}

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

	response := response.ConvertToRoomchatUserResponse(&roomchat, &doctor)

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

	var user schema.User
	if err := configs.DB.Where("id = ?", doctortransaction.UserID).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve user data"))
	}

	response := response.ConvertToRoomchatDoctorResponse(&roomchat, &user)

	return c.JSON(http.StatusOK, helper.SuccessResponse("roomchat data successfully retrieved", response))
}


// Doctor Get All Roomchats 
func GetAllDoctorRoomchatController(c echo.Context) error {

	doctorID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid doctor id"))
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("limit"+constanta.ErrQueryParamRequired))
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("offset"+constanta.ErrQueryParamRequired))
	}

	var DoctorTransactions []schema.DoctorTransaction

	existingDoctorTransactions, total, err := GetAllRoomchatPagination(doctorID, offset, limit, DoctorTransactions)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("doctor transaction data not found"))
	}

	sort.Slice(existingDoctorTransactions, func(i, j int) bool {

		getLastMessageTimestamp := func(transaction schema.DoctorTransaction) time.Time {
			if len(transaction.Roomchat.Message) > 0 {
				return transaction.Roomchat.Message[len(transaction.Roomchat.Message)-1].CreatedAt
			}

			return time.Time{}
		}

		return getLastMessageTimestamp(existingDoctorTransactions[i]).After(getLastMessageTimestamp(existingDoctorTransactions[j]))
	})

	var responses []web.RoomchatListResponse
	for _, doctorTransaction := range existingDoctorTransactions {

		if doctorTransaction.DoctorID != uint(doctorID) {
			continue
		}

		if doctorTransaction.Roomchat.ID == 0 {
			continue
		}

		if doctorTransaction.PaymentStatus != "success" {
			continue
		}

		var user schema.User
		if err := configs.DB.First(&user, "id = ?", doctorTransaction.UserID).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve user data"))
		}

		var lastMessage schema.Message
		if len(doctorTransaction.Roomchat.Message) > 0 {
			lastMessage = doctorTransaction.Roomchat.Message[len(doctorTransaction.Roomchat.Message)-1]
		}

		response := response.ConvertToGetAllRoomchats(user, doctorTransaction.Roomchat, lastMessage)
		responses = append(responses, response)
	}

	pagination := helper.Pagination(offset, limit, total)

	return c.JSON(http.StatusOK, helper.PaginationResponse("roomchat data successfully retrieved", responses, pagination))
}
