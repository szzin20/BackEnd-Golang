package controllers

import (
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// User Create Complaint
func CreateComplaintController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Transaction ID"))
	}

	var existingTransactionID schema.DoctorTransaction

	result := configs.DB.First(&existingTransactionID, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Transaction ID"))
	}

	var complaint web.ComplaintRequest

	if err := c.Bind(&complaint); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Complaint Data"))
	}

	complaintRequest := request.ConvertToComplaintRequest(complaint, existingTransactionID.ID)

	if err := configs.DB.Create(&complaintRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Send Complaint"))
	}

	response := response.ConvertToComplaintResponse(complaintRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Complaint Successful", response))
}

// User Get Complaint by ID
func GetComplaintsController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Complaint ID"))
	}

	var complaint schema.Complaint

	if err := configs.DB.First(&complaint, userID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Complaint Data"))
	}

	response := response.ConvertToComplaintResponse(&complaint)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Complaint Data Successfully Retrieved", response))
}
