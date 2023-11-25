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

	"github.com/labstack/echo/v4"
)

// User Create Complaint
func CreateComplaintController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Transaction ID"))
	}

	transactionID, _ := strconv.Atoi(c.QueryParam("transaction_id"))

	var doctorTransaction schema.DoctorTransaction

	if err := configs.DB.First(&doctorTransaction, "user_id = ? AND id = ?", userID, transactionID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Doctor Transaction Data"))
	}

	var complaintRequest web.CreateComplaintRequest

	if err := c.Bind(&complaintRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Complaint Data"))
	}

	if err := helper.ValidateStruct(complaintRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	err := c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	file, fileHeader, err := c.Request().FormFile("image")
	if err == http.ErrMissingFile {
    	complaintRequest.Image = ""
	} else if err != nil {
    	return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Image File is Required"))
	} else {
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
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid image file format. Supported formats: jpg, jpeg, png"))
		}
	
		images, err := helper.UploadFilesToGCS(c, fileHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error uploading image to Cloud Storage"))
		}
	
		complaintRequest.Image = images
	}

	complaint := request.ConvertToComplaintRequest(complaintRequest, uint(transactionID))

	if err := configs.DB.Create(&complaint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Send Complaint"))
	}

	response := response.ConvertToComplaintResponse(complaint)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Complaint Successful", response))
}


// Get Complaint by DoctorTransaction ID
func GetComplaintsController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Complaint ID"))
	}

	transactionID, _ := strconv.Atoi(c.QueryParam("transaction_id"))

	var doctorTransaction schema.DoctorTransaction

	if err := configs.DB.Preload("Complaint").Where("user_id = ? AND id = ?", userID, transactionID).First(&doctorTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Doctor Transaction Data"))
	}

	response := response.ConvertToDoctorTransactionResponse(&doctorTransaction)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Complaint Data Successfully Retrieved", response))
}
