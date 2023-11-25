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

// Doctor Create Advice
func CreateAdviceController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	transactionID, _ := strconv.Atoi(c.QueryParam("transaction_id"))

	var doctorTransaction schema.DoctorTransaction

	if err := configs.DB.First(&doctorTransaction, "user_id = ? AND id = ?", userID, transactionID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
	}

	var adviceRequest web.AdviceRequest

	if err := c.Bind(&adviceRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input advice data"))
	}

	if err := helper.ValidateStruct(adviceRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	err := c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	file, fileHeader, err := c.Request().FormFile("image")

	if err == nil {
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
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid image file format. supported formats: .jpg, .jpeg, .png"))
		}

		adviceImage, err := helper.UploadFilesToGCS(c, fileHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error uploading image to cloud storage"))
		}

		adviceRequest.Image = adviceImage
	}


	advice := request.ConvertToAdviceRequest(adviceRequest, uint(transactionID))

	if err := configs.DB.Create(&advice).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send advice"))
	}

	response := response.ConvertToAdviceResponse(advice)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("advice successful", response))
}

// Get Advice by DoctorTransaction ID
func GetAdvicesController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	transactionID, _ := strconv.Atoi(c.QueryParam("transaction_id"))

	var doctorTransaction schema.DoctorTransaction

	if err := configs.DB.
	Preload("Complaint").
	Preload("Advice").
    Where("user_id = ? AND id = ?", userID, transactionID).
    Order("created_at ASC").
    First(&doctorTransaction).Error; err != nil {
    return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
}
	response := response.ConvertToAdviceResponse(&doctorTransaction)

	return c.JSON(http.StatusOK, helper.SuccessResponse("advice data successfully retrieved", response))
}
