package controllers

import (
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"
	"path"
	"path/filepath"
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

	// Upload files
	err := c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	file, fileHeader, err := c.Request().FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Image File is Required"))
	}
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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid image file format. Supported formats: jpg, jpeg, png"))
	}

	image, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error upload image to Cloud Storage"))
	}

	medicine.Image = image

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

	var updatedMedicineRequest web.MedicineUpdateRequest

	if err := c.Bind(&updatedMedicineRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Medicine Data"))
	}

	if err := helper.ValidateStruct(updatedMedicineRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Update data obat di database
	result = configs.DB.Model(&existingMedicine).Updates(updatedMedicineRequest)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Update Medicine"))
	}

	response := response.ConvertToAdminMedicineUpdateResponse(&existingMedicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Updated Successfully", response))
}

// Update Image Medicine by ID
func UpdateImageMedicineController(c echo.Context) error {
	medicineID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Medicine ID"))
	}

	var existingMedicine schema.Medicine
	if err := configs.DB.First(&existingMedicine, medicineID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicine ID"))
	}

	err = c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	file, fileHeader, err := c.Request().FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Image File is Required"))
	}
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

	newImage, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to upload image to Cloud Storage"))
	}

	if existingMedicine.Image != "" {
		oldFilename := path.Base(existingMedicine.Image)

		if err := helper.DeleteFilesFromGCS(oldFilename); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to delete old image"))
		}
	}

	existingMedicine.Image = newImage

	if err := configs.DB.Save(&existingMedicine).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to update Medicine"))
	}

	response := response.ConvertToAdminMedicineImageResponse(&existingMedicine)
	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Image Updated Successfully", response))
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

// Delete Image Medicine by ID
func DeleteImageMedicineController(c echo.Context) error {
	medicineID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Medicine ID"))
	}

	var medicine schema.Medicine
	if err := configs.DB.First(&medicine, medicineID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicine ID"))
	}

	if medicine.Image != "" {
		filename := path.Base(medicine.Image)
		if err := helper.DeleteFilesFromGCS(filename); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Delete Image"))
		}
	}

	medicine.Image = ""

	if err := configs.DB.Save(&medicine).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Update Medicine"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Image Deleted Successfully", nil))
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

// Get Image Medicine by ID
func GetImageMedicineController(c echo.Context) error {
	medicineID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Medicine ID"))
	}

	var medicine schema.Medicine
	if err := configs.DB.First(&medicine, medicineID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicine Data"))
	}

	response := response.ConvertToAdminMedicineImageResponse(&medicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Image Data Successfully Retrieved", response))
}

// Admin Get All Medicines
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

// Admin Get Medicine by Name
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

// User Get All Medicines
func GetAllMedicinesUserController(c echo.Context) error {
	var medicines []schema.Medicine

	err := configs.DB.Find(&medicines).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Medicines Data"))
	}

	if len(medicines) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Medicines Data"))
	}

	response := response.ConvertToUserGetAllMedicinesResponse(medicines)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicines Data Successfully Retrieved", response))
}

// User Get Medicine by Name
func GetMedicineByNameUserController(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Name parameter is required"))
	}

	var medicine schema.Medicine

	result := configs.DB.Where("name = ?", name).First(&medicine)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Medicine not found"))
	}

	response := response.ConvertToUserMedicineResponse(&medicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Medicine Data Successfully Retrieved", response))
}
