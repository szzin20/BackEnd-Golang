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

// Create Advice
func CreateAdviceController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Complaint ID"))
	}

	var existingComplaintID schema.Complaint

	result := configs.DB.First(&existingComplaintID, id)
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

// Get All Advices
func GetAllAdvicesController(c echo.Context) error {
	var advices []schema.Advice

	err := configs.DB.Find(&advices).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Advices Data"))
	}

	if len(advices) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Advices Data"))
	}

	response := response.ConvertToGetAllAdvicesResponse(advices)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Advices Data Successfully Retrieved", response))
}

// Get Advice by ID
func GetAdviceController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Advice ID"))
	}

	var advice schema.Advice

	if err := configs.DB.First(&advice, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Advice Data"))
	}

	response := response.ConvertToAdviceResponse(&advice)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Advice Data Successfully Retrieved", response))
}

