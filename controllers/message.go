package controllers

import (
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// User Create Complaint Message
func CreateComplaintMessageController(c echo.Context) error {

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

	if existingRoomchat.ExpirationTime != nil && time.Now().After(*existingRoomchat.ExpirationTime) {
		return c.JSON(http.StatusForbidden, helper.ErrorResponse("roomchat expired"))
	}

	var doctortransaction schema.DoctorTransaction
	if err := configs.DB.Where("user_id = ? AND id = ?", userID, existingRoomchat.TransactionID).First(&doctortransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
	}

	var complaintMessageRequest web.CreateMessageRequest

	if err := c.Bind(&complaintMessageRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input complaint message data"))
	}

	if err := helper.ValidateStruct(complaintMessageRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	if (complaintMessageRequest.Image != "" && complaintMessageRequest.Audio != "") || (complaintMessageRequest.Image != "" && complaintMessageRequest.Message != "") || (complaintMessageRequest.Audio != "" && complaintMessageRequest.Message != "") {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("please choose only one type of message: text, image, or audio"))
	}

	if complaintMessageRequest.Image != "" {

		err = c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("file size exceeds the maximum allowed size"))
		}

		file, fileHeader, err := c.Request().FormFile("image")

		if err == nil {
			defer file.Close()

			if fileHeader.Size > 10*1024*1024 { // 10 MB limit
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("image file size exceeds the limit (10 MB)"))
			}

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
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid image file format. supported formats: .jpg, .jpeg, .png"))
			}

			imageURL, err := helper.UploadFilesToGCS(c, fileHeader)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error uploading image to cloud storage"))
			}

			complaintMessageRequest.Image = imageURL
		}
	}

	if complaintMessageRequest.Audio != "" {

		file, fileHeader, err := c.Request().FormFile("audio")

		if err == nil {
			defer file.Close()

			maxFileSize := int64(10 * 1024 * 1024) // 10 MB

			if fileHeader.Size > maxFileSize {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("file size exceeds the maximum allowed size"))
			}

			allowedAudioExtensions := []string{".mp3", ".wav", ".flac"}
			ext := filepath.Ext(fileHeader.Filename)
			allowed := false
			for _, validExt := range allowedAudioExtensions {
				if ext == validExt {
					allowed = true
					break
				}
			}
			if !allowed {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid audio file format. supported formats: .mp3, .wav, .flac"))
			}

			audioURL, err := helper.UploadFilesToGCS(c, fileHeader)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error uploading audio to cloud storage"))
			}

			complaintMessageRequest.Audio = audioURL
		}
	}

	complaint := request.ConvertToCreateComplaintMessageRequest(complaintMessageRequest, uint(roomchatID), uint(userID))

	if err := configs.DB.Create(&complaint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send complaint message"))
	}

	response := response.ConvertToCreateMessageResponse(complaint)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("complaint message successful send", response))
}

// Doctor Create Advice Message
func CreateAdviceMessageController(c echo.Context) error {

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

	if existingRoomchat.ExpirationTime != nil && time.Now().After(*existingRoomchat.ExpirationTime) {
		return c.JSON(http.StatusForbidden, helper.ErrorResponse("roomchat expired"))
	}

	var doctortransaction schema.DoctorTransaction
	if err := configs.DB.Where("doctor_id = ? AND id = ?", doctorID, existingRoomchat.TransactionID).First(&doctortransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
	}

	var adviceMessageRequest web.CreateMessageRequest

	if err := c.Bind(&adviceMessageRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input advice message data"))
	}

	if err := helper.ValidateStruct(adviceMessageRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	if (adviceMessageRequest.Image != "" && adviceMessageRequest.Audio != "") || (adviceMessageRequest.Image != "" && adviceMessageRequest.Message != "") || (adviceMessageRequest.Audio != "" && adviceMessageRequest.Message != "") {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("please choose only one type of message: text, image, or audio"))
	}

	if adviceMessageRequest.Image != "" {

		err = c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		file, fileHeader, err := c.Request().FormFile("image")

		if err == nil {
			defer file.Close()

			if fileHeader.Size > 10*1024*1024 { // 10 MB limit
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("image file size exceeds the limit (10 MB)"))
			}

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
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid image file format. supported formats: .jpg, .jpeg, .png"))
			}

			imageURL, err := helper.UploadFilesToGCS(c, fileHeader)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error uploading image to cloud storage"))
			}

			adviceMessageRequest.Image = imageURL
		}

	}

	if adviceMessageRequest.Audio != "" {

		file, fileHeader, err := c.Request().FormFile("audio")

		if err == nil {
			defer file.Close()

			maxFileSize := int64(10 * 1024 * 1024) // 10 MB

			if fileHeader.Size > maxFileSize {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("file size exceeds the maximum allowed size"))
			}

			allowedAudioExtensions := []string{".mp3", ".wav", ".flac"}
			ext := filepath.Ext(fileHeader.Filename)
			allowed := false
			for _, validExt := range allowedAudioExtensions {
				if ext == validExt {
					allowed = true
					break
				}
			}
			if !allowed {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid audio file format. supported formats: .mp3, .wav, .flac"))
			}

			audioURL, err := helper.UploadFilesToGCS(c, fileHeader)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error uploading audio to cloud storage"))
			}

			adviceMessageRequest.Audio = audioURL
		}

	}

	advice := request.ConvertToCreateAdviceMessageRequest(adviceMessageRequest, uint(roomchatID), uint(doctorID))

	if err := configs.DB.Create(&advice).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send advice message"))
	}

	response := response.ConvertToCreateMessageResponse(advice)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("advice message successful send", response))
}
