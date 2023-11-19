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

// Create Medicine
func CreateMedicineController(c echo.Context) error {
	var medicine web.MedicineRequest

	if err := c.Bind(&medicine); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Medicine Data"))
	}

	if err := helper.ValidateStruct(medicine); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	medicineRequest := request.ConvertToMedicineRequest(medicine)

	if err := configs.DB.Create(&medicineRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create Medicine"))
	}

	response := response.ConvertToAdminMedicineResponse(medicineRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Medicine Created Successfully", response))
}

// Update Medicine by ID
func UpdateMedicineController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Medicine ID"))
	}

	var existingMedicine schema.Medicine

	result := configs.DB.First(&existingMedicine, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicine ID"))
	}

	var updatedMedicineRequest web.MedicineRequest

	if err := c.Bind(&updatedMedicineRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Medicine Data"))
	}

	if err := helper.ValidateStruct(updatedMedicineRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	updatedMedicine := request.ConvertToMedicineRequest(updatedMedicineRequest)

	result = configs.DB.Model(&existingMedicine).Updates(updatedMedicine)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Update Medicine"))
	}

	response := response.ConvertToAdminMedicineResponse(&existingMedicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Updated Successfully", response))
}

// Delete Medicine by ID
func DeleteMedicineController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Medicine ID"))
	}

	var medicine schema.Medicine

	result := configs.DB.First(&medicine, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicine ID"))
	}

	result = configs.DB.Delete(&medicine, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Delete Medicine"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Deleted Successfully", nil))
}

// Get Medicine by ID
func GetMedicineController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Medicine ID"))
	}

	var medicine schema.Medicine

	if err := configs.DB.First(&medicine, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicine Data"))
	}

	response := response.ConvertToAdminMedicineResponse(&medicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Data Successfully Retrieved", response))
}

// Get All Medicines (Admin)
func GetAllMedicinesAdminController(c echo.Context) error {
	var medicines []schema.Medicine

	err := configs.DB.Find(&medicines).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicines Data"))
	}

	if len(medicines) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Medicines Data"))
	}

	response := response.ConvertToAdminGetAllMedicinesResponse(medicines)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicines Data Successfully Retrieved", response))
}

// Get Medicine by Name (Admin)
func GetMedicineByNameAdminController(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Name parameter is required"))
	}

	var medicine schema.Medicine

	result := configs.DB.Where("name = ?", name).First(&medicine)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Medicine not found"))
	}

	response := response.ConvertToAdminMedicineResponse(&medicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Data Successfully Retrieved", response))
}

// Get All Medicines (Patient)
func GetAllMedicinesPatientController(c echo.Context) error {
	var medicines []schema.Medicine

	err := configs.DB.Find(&medicines).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicines Data"))
	}

	if len(medicines) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Medicines Data"))
	}

	response := response.ConvertToPatientGetAllMedicinesResponse(medicines)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicines Data Successfully Retrieved", response))
}

// Get Medicine by Name (Patient)
func GetMedicineByNamePatientController(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Name parameter is required"))
	}

	var medicine schema.Medicine

	result := configs.DB.Where("name = ?", name).First(&medicine)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Medicine not found"))
	}

	response := response.ConvertToPatientMedicineResponse(&medicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Data Successfully Retrieved", response))
}
