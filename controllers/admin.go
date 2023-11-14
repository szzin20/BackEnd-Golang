package controllers

import (
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

// LoginAdminController handles admin login requests.
func LoginAdminController(c echo.Context) error {
	// Bind the request body to the AdminLoginRequest struct
	loginRequest := new(web.AdminLoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid request"))
	}

	// Use the conversion function to convert AdminLoginRequest to Admin
	admin := request.ConvertToAdminLoginRequest(*loginRequest)

	// Find the admin by email
	if err := configs.DB.Where("email = ? AND password = ?", loginRequest.Email, loginRequest.Password).First(&admin).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("invalid email or password"))
	}

	// Convert the admin to login response
	loginResponse := response.ConvertToAdminLoginResponse(admin)

	// Return the success response with JWT token
	return c.JSON(http.StatusOK, helper.SuccessResponse("login successful", loginResponse))
}

// UpdateAdminController handles admin update requests.
func UpdateAdminController(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
    }

    var updatedAdmin web.AdminUpdateRequest

    if err := c.Bind(&updatedAdmin); err != nil {
        return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid request body"))
    }

    var existingAdmin schema.Admin
    result := configs.DB.First(&existingAdmin, id)
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to retrieve user"))
    }

    configs.DB.Model(&existingAdmin).Updates(schema.Admin{
        Name:     updatedAdmin.Name,
        Email:    updatedAdmin.Email,
        Password: updatedAdmin.Password,
    })

    response := response.ConvertToAdminUpdateResponse(&existingAdmin)

    return c.JSON(http.StatusOK, helper.SuccessResponse("User data successfully updated", response))
}
