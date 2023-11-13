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

// Create Complaint
func CreateComplaintController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Transaction ID"))
	}

	var existingTransactionID schema.DoctorTransaction

	result := configs.DB.First(&existingTransactionID, id)
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

// Get All Complaint
func GetAllComplaintsController(c echo.Context) error {
	var complaints []schema.Complaint

	err := configs.DB.Find(&complaints).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Complaints Data"))
	}

	if len(complaints) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Complaints Data"))
	}

	response := response.ConvertToGetAllComplaintsResponse(complaints)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Complaints Data Successfully Retrieved", response))
}


// Get Complaint by ID
func GetComplaintController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Complaint ID"))
	}

	var complaint schema.Complaint

	if err := configs.DB.First(&complaint, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Complaint Data"))
	}

	response := response.ConvertToComplaintResponse(&complaint)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Complaint Data Successfully Retrieved", response))
}
