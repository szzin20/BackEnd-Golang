package controllers

import (
	"healthcare/configs"
	"healthcare/middlewares"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Get All User
func GetAllUsersController(c echo.Context) error {
	var users []schema.User

	err := configs.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Users Data"))
	}

	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Users Data"))
	}

	response := response.ConvertToGetAllUsersResponse(users)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Users Data Successfully Retrieved", response))
}

// Get User by ID
func GetUserController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid User ID"))
	}

	var user schema.User

	if err := configs.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve User Data"))
	}

	response := response.ConvertToGetUserResponse(&user)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Users Data Successfully Retrieved", response))
}

// Register User
func RegisterUserController(c echo.Context) error {
	var user web.UserRegisterRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Register Data"))
	}

	userRequest := request.ConvertToUserRegisterRequest(user)

	if existingUser := configs.DB.Where("email = ?", userRequest.Email).First(&userRequest).Error; existingUser == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))
	}

	userRequest.Password = helper.HashPassword(userRequest.Password)

	if err := configs.DB.Create(&userRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Register"))
	}

	response := response.ConvertToUserRegisterResponse(userRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Registered Successful", response))
}

// User Login
func LoginUserController(c echo.Context) error {
	var loginRequest web.UserLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Login Data"))
	}

	var user schema.User
	if err := configs.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Email Not Registered"))
	}

	if err := helper.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("Incorrect Email or Password"))
	}

	userLoginResponse := response.ConvertToUserLoginResponse(user)

	token, err := middlewares.GenerateToken(&userLoginResponse, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Generate JWT"))
	}

	userLoginResponse.Token = token

	return c.JSON(http.StatusOK, helper.SuccessResponse("Login Successful", userLoginResponse))
}


// Update User
func UpdateUserController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid user ID"))
	}

	var existingUser schema.User

	result := configs.DB.First(&existingUser, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve User"))
	}

	var userUpdated web.UserUpdateRequest

	if err := c.Bind(&userUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Update Data"))
	}

	userUpdated.Password = helper.HashPassword(userUpdated.Password)

	configs.DB.Model(&existingUser).Updates(userUpdated)

	response := response.ConvertToUserUpdateResponse(&existingUser)

	return c.JSON(http.StatusOK, helper.SuccessResponse("User Updated Data Successful", response))
}

// Delete User
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid User ID"))
	}

	var existingUser schema.User
	result := configs.DB.First(&existingUser, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve User"))
	}

	configs.DB.Delete(&existingUser)

	return c.JSON(http.StatusOK, helper.SuccessResponse("User Deleted Data Successful", nil))
}


