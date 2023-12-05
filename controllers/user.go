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
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// User Register
func RegisterUserController(c echo.Context) error {
	var user web.UserRegisterRequest

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input register data"))
	}

	if err := helper.ValidateStruct(user); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	userRequest := request.ConvertToUserRegisterRequest(user)

	var existingUserEmail schema.User
	if existingEmail := configs.DB.Where("email = ? AND deleted_at IS NULL", userRequest.Email).First(&existingUserEmail).Error; existingEmail == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("email already exist"))
	}

	userRequest.Password = helper.HashPassword(userRequest.Password)

	if err := configs.DB.Create(&userRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to register"))
	}

	// send register notification email
	if user.Email != "" {
		notificationType := "register"
		if err := helper.SendNotificationEmail(user.Email, user.Fullname, notificationType, "Login", "", ""); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send notification email: "+err.Error()))
		}
	}

	response := response.ConvertToUserRegisterResponse(userRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("registered successful", response))
}

// User Login
func LoginUserController(c echo.Context) error {
	var loginRequest web.UserLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input login data"))
	}

	if err := helper.ValidateStruct(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	var user schema.User
	if err := configs.DB.Where("email = ? AND deleted_at IS NULL", loginRequest.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("email not registered"))
	}

	if err := helper.ComparePassword(user.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("incorrect email or password"))
	}

	userLoginResponse := response.ConvertToUserLoginResponse(user)

	token, err := middlewares.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to generate jwt"))
	}

	userLoginResponse.Token = token

	// send login notification email
	if user.Email != "" {
		notificationType := "login"
		if err := helper.SendNotificationEmail(user.Email, user.Fullname, notificationType, "Login", "", ""); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send notification email: "+err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("login successful", userLoginResponse))
}

// Get All Doctors by Admin
func GetAllUserByAdminController(c echo.Context) error {
	var User []schema.User

	err := configs.DB.Find(&User).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Mengambil Data Pengguna"))
	}

	if len(User) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data Pengguna Kosong"))
	}

	response := response.ConvertToGetAllUserByAdminResponse(User)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data Pengguna Berhasil Diambil", response))
}

// Get User Profile
func GetUserController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	var user schema.User

	if err := configs.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve user data"))
	}

	response := response.ConvertToGetUserResponse(&user)

	return c.JSON(http.StatusOK, helper.SuccessResponse("users data successfully retrieved", response))
}

// Get User by ID
func GetUserIDbyAdminController(c echo.Context) error {
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid user id"))
	}
	var user schema.User
	result := configs.DB.First(&user, user_id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve user data"))
	}
	response := response.ConvertToGetUserIDbyAdminResponse(&user)

	return c.JSON(http.StatusOK, helper.SuccessResponse("users data successfully retrieved", response))
}

// Update User Profile
func UpdateUserController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	var existingUser schema.User

	result := configs.DB.Where("id = ?", userID).First(&existingUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve user"))
	}

	var userUpdated web.UserUpdateRequest

	if err := c.Bind(&userUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input update data"))
	}

	var existingUserEmail schema.User
	if existingEmail := configs.DB.Where("email = ? AND deleted_at IS NULL", userUpdated.Email).First(&existingUserEmail).Error; existingEmail == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("email already exist"))
	}

	if err := helper.ValidateStruct(userUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	err := c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	file, fileHeader, err := c.Request().FormFile("profile_picture")

	if err == nil {
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
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid image file format. supported formats: .jpg, .jpeg, .png"))
		}

		profilePicture, err := helper.UploadFilesToGCS(c, fileHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error uploading image to cloud storage"))
		}

		userUpdated.ProfilePicture = profilePicture
	}

	userUpdated.Password = helper.HashPassword(userUpdated.Password)
	gender := strings.ToLower(userUpdated.Gender)
	bloodType := strings.ToUpper(userUpdated.BloodType)
	birthdate := userUpdated.Birthdate

	if !helper.GenderIsValid(gender) {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input gender data ('male', 'female')"))
	}

	if !helper.BloodTypeIsValid(bloodType) {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input blood type data,('a', 'b', 'o', 'ab')"))
	}

	if !helper.BirthdateIsValid(birthdate) {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input birthdate data (yyyy-mm-dd)"))
	}

	configs.DB.Model(&existingUser).Updates(userUpdated)

	userResponse := response.ConvertToUserUpdateResponse(&existingUser)

	return c.JSON(http.StatusOK, helper.SuccessResponse("user updated data successful", userResponse))
}

// Delete User
func DeleteUserController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	var existingUser schema.User
	result := configs.DB.First(&existingUser, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve user"))
	}

	configs.DB.Delete(&existingUser)

	return c.JSON(http.StatusOK, helper.SuccessResponse("user deleted data successful", nil))
}

func DeleteUserByAdminController(c echo.Context) error {
	// Parse doctor ID from the request parameters
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid user id"))
	}

	// Retrieve the existing doctor from the database
	var existingUser schema.User
	result := configs.DB.First(&existingUser, user_id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve user"))
	}

	// Delete the doctor from the database
	result = configs.DB.Delete(&existingUser, user_id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to delete user"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("user deleted data successful  ", nil))
}
