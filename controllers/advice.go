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

// Doctor Create Advice
func CreateAdviceController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Complaint ID"))
	}

	var existingComplaintID schema.Complaint

	result := configs.DB.First(&existingComplaintID, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Complaint ID"))
	}

	var advice web.AdviceRequest

	if err := c.Bind(&advice); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Advice Data"))
	}

	adviceRequest := request.ConvertToAdviceRequest(advice, existingComplaintID.ID)

	if err := configs.DB.Create(&adviceRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Send Advice"))
	}

	response := response.ConvertToAdviceResponse(adviceRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Advice Successful", response))
}

// Doctor Get Advice by ID
func GetAdvicesController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Advice ID"))
	}

	var advice schema.Advice

	if err := configs.DB.First(&advice, userID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Advice Data"))
	}

	response := response.ConvertToAdviceResponse(&advice)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Advice Data Successfully Retrieved", response))
}

