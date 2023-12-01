package controllers

import (
	"github.com/labstack/echo/v4"
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
)

//func GetAll(offset int, limit int, queryInput []schema.Medicine) ([]schema.Medicine, int64, error) {
//
//	if offset < 0 || limit < 0 {
//		return nil, 0, nil
//	}
//
//	queryAll := queryInput
//	var total int64
//
//	query := configs.DB.Model(&queryAll)
//
//	query.Find(&queryAll).Count(&total)
//
//	query = query.Limit(limit).Offset(offset)
//
//	result := query.Find(&queryAll)
//
//	if result.Error != nil {
//		return nil, 0, result.Error
//	}
//
//	if offset >= int(total) {
//		return nil, 0, fmt.Errorf("not found")
//	}
//
//	return queryAll, total, nil
//}

// Create Medicine
func CreateMedicineController(c echo.Context) error {

	var medicine web.MedicineRequest

	if err := c.Bind(&medicine); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input medicine data"))
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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("image file is required"))
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

	imageURL, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error upload image to cloud storage"))
	}

	medicine.Image = imageURL

	medicineRequest := request.ConvertToMedicineRequest(medicine)

	if err := configs.DB.Create(&medicineRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to create medicine"))
	}

	response := response.ConvertToAdminMedicineResponse(medicineRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("medicine created successfully", response))
}

// Update Medicine by ID
func UpdateMedicineController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("medicines_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid medicine id"))
	}

	var existingMedicine schema.Medicine

	result := configs.DB.First(&existingMedicine, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve medicine id"))
	}

	var updatedMedicineRequest web.MedicineUpdateRequest

	if err := c.Bind(&updatedMedicineRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input medicine data"))
	}

	if err := helper.ValidateStruct(updatedMedicineRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Update data obat di database
	result = configs.DB.Model(&existingMedicine).Updates(updatedMedicineRequest)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to update medicine"))
	}

	response := response.ConvertToAdminMedicineUpdateResponse(&existingMedicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("medicine updated successfully", response))
}

// Update Image Medicine by ID
func UpdateImageMedicineController(c echo.Context) error {
	medicineID, err := strconv.Atoi(c.Param("medicines_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid medicine id"))
	}

	var existingMedicine schema.Medicine
	if err := configs.DB.First(&existingMedicine, medicineID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve medicine id"))
	}

	err = c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	file, fileHeader, err := c.Request().FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("image file is required"))
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

	newImage, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to upload image to cloud storage"))
	}

	if existingMedicine.Image != "" {
		oldFilename := path.Base(existingMedicine.Image)

		if err := helper.DeleteFilesFromGCS(oldFilename); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to delete old image"))
		}
	}

	existingMedicine.Image = newImage

	if err := configs.DB.Save(&existingMedicine).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to update medicine"))
	}

	response := response.ConvertToAdminMedicineImageResponse(&existingMedicine)
	return c.JSON(http.StatusOK, helper.SuccessResponse("medicine image updated successfully", response))
}

// Delete Medicine by ID
func DeleteMedicineController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("medicines_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid medicine id"))
	}

	var medicine schema.Medicine

	result := configs.DB.First(&medicine, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve medicine id"))
	}

	result = configs.DB.Delete(&medicine, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to delete medicine"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("medicine deleted successfully", nil))
}

// Delete Image Medicine by ID
func DeleteImageMedicineController(c echo.Context) error {
	medicineID, err := strconv.Atoi(c.Param("medicines_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid medicine id"))
	}

	var medicine schema.Medicine
	if err := configs.DB.First(&medicine, medicineID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve medicine id"))
	}

	if medicine.Image != "" {
		filename := path.Base(medicine.Image)
		if err := helper.DeleteFilesFromGCS(filename); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to delete image"))
		}
	}

	medicine.Image = ""

	if err := configs.DB.Save(&medicine).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to update medicine"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("medicine image deleted successfully", nil))
}

// Get Image Medicine by ID
func GetImageMedicineController(c echo.Context) error {
	medicineID, err := strconv.Atoi(c.Param("medicines_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid medicine id"))
	}

	var medicine schema.Medicine
	if err := configs.DB.First(&medicine, medicineID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve medicine data"))
	}

	response := response.ConvertToAdminMedicineImageResponse(&medicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("medicine image data successfully retrieved", response))
}

// Admin Get All Medicines Pagination
//
//	func GetAllMedicinesAdminController(c echo.Context) error {
//		params := c.QueryParams()
//		limit, err := strconv.Atoi(params.Get("limit"))
//
//		if err != nil {
//			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("params limit not valid"))
//		}
//
//		offset, err := strconv.Atoi(params.Get("offset"))
//
//		if err != nil {
//			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("params offset not valid"))
//		}
//
//		var medicines []schema.Medicine
//
//		medicine, total, err := GetAll(offset, limit, medicines)
//
//		if err != nil {
//			if strings.Contains(err.Error(), "not found") {
//				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Medicines Not Found"))
//			}
//			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
//		}
//
//		pagination := helper.Pagination(offset, limit, total)
//
//		response := response.ConvertToAdminGetAllMedicinesResponse(medicine)
//
//		return c.JSON(http.StatusOK, helper.PaginationResponse("Medicines Data Successfully Retrieved", response, pagination))
//	}

// Admin Get Medicine by ID
func GetMedicineAdminByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("medicines_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid medicine id"))
	}

	var medicine schema.Medicine

	if err := configs.DB.First(&medicine, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve medicine data"))
	}

	response := response.ConvertToAdminMedicineResponse(&medicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("medicine data successfully retrieved", response))
}

// Admin Get Medicines
func GetMedicineAdminController(c echo.Context) error {
	idStr := c.QueryParam("medicines_id")
	name := c.QueryParam("name")

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid medicine id"))
		}

		var medicine schema.Medicine

		if err := configs.DB.First(&medicine, id).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve medicine data"))
		}

		response := response.ConvertToAdminMedicineResponse(&medicine)

		return c.JSON(http.StatusOK, helper.SuccessResponse("medicine data successfully retrieved", response))

	} else if name != "" {
		var medicine schema.Medicine

		result := configs.DB.Where("name LIKE ?", "%"+name+"%").First(&medicine)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("medicine not found"))
		}

		response := response.ConvertToAdminMedicineResponse(&medicine)

		return c.JSON(http.StatusOK, helper.SuccessResponse("medicine data successfully retrieved", response))

	} else {
		var medicines []schema.Medicine

		err := configs.DB.Find(&medicines).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve medicines data"))
		}

		if len(medicines) == 0 {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("empty medicines data"))
		}

		response := response.ConvertToAdminGetAllMedicinesResponse(medicines)

		return c.JSON(http.StatusOK, helper.SuccessResponse("medicines data successfully retrieved", response))
	}
}

// User Get All Medicines Pagination
//func GetAllMedicinesUserController(c echo.Context) error {
//	params := c.QueryParams()
//	limit, err := strconv.Atoi(params.Get("limit"))
//
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("params limit not valid"))
//	}
//
//	offset, err := strconv.Atoi(params.Get("offset"))
//
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("params offset not valid"))
//	}
//
//	var medicines []schema.Medicine
//
//	medicine, total, err := GetAll(offset, limit, medicines)
//
//	if err != nil {
//		if strings.Contains(err.Error(), "not found") {
//			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Medicines Not Found"))
//		}
//		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
//	}
//
//	pagination := helper.Pagination(offset, limit, total)
//
//	response := response.ConvertToUserGetAllMedicinesResponse(medicine)
//
//	return c.JSON(http.StatusOK, helper.PaginationResponse("medicines data successfully retrieved", response, pagination))
//}

// User Get Medicine
func GetMedicineUserController(c echo.Context) error {
	idStr := c.QueryParam("medicines_id")
	name := c.QueryParam("name")

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid medicine id"))
		}

		var medicine schema.Medicine

		if err := configs.DB.First(&medicine, id).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve medicine data"))
		}

		response := response.ConvertToUserMedicineResponse(&medicine)

		return c.JSON(http.StatusOK, helper.SuccessResponse("medicine data successfully retrieved", response))

	} else if name != "" {
		var medicine schema.Medicine

		result := configs.DB.Where("name LIKE ?", "%"+name+"%").First(&medicine)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("medicine not found"))
		}

		response := response.ConvertToUserMedicineResponse(&medicine)

		return c.JSON(http.StatusOK, helper.SuccessResponse("medicine data successfully retrieved", response))

	} else {
		var medicines []schema.Medicine

		err := configs.DB.Find(&medicines).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve medicines data"))
		}

		if len(medicines) == 0 {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("empty medicines data"))
		}

		response := response.ConvertToUserGetAllMedicinesResponse(medicines)

		return c.JSON(http.StatusOK, helper.SuccessResponse("medicines data successfully retrieved", response))
	}
}

// User Get Medicine by ID
func GetMedicineUserByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("medicines_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid medicine id"))
	}

	var medicine schema.Medicine

	if err := configs.DB.First(&medicine, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve medicine data"))
	}

	response := response.ConvertToUserMedicineResponse(&medicine)

	return c.JSON(http.StatusOK, helper.SuccessResponse("medicine data successfully retrieved", response))
}
