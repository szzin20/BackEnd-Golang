package controllers

import (
	"fmt"
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

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// RegisterDoctorController
func RegisterDoctorByAdminController(c echo.Context) error {
	var doctor web.DoctorRegisterRequest

	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input register data"))
	}
	if err := helper.ValidateStruct(doctor); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	doctorRequest := request.ConvertToDoctorRegisterRequest(doctor)

	err := c.Request().ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	file, fileHeader, err := c.Request().FormFile("profile_picture")
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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid image file format. supported formats: jpg, jpeg, png"))
	}

	imageURL, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error upload image to cloud storage"))
	}

	doctorRequest.ProfilePicture = imageURL
	// Periksa apakah email sudah ada
	if existingDoctor := configs.DB.Where("email = ?", doctorRequest.Email).First(&doctorRequest).Error; existingDoctor == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("email already exist"))
	}

	// Hash kata sandi
	doctorRequest.Password = helper.HashPassword(doctor.Password)

	// Simpan data dokter ke database
	if err := configs.DB.Create(&doctorRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to register"))
	}

	// Mengirim email pemberitahuan
	err = helper.SendNotificationEmail(doctorRequest.Email, doctorRequest.Fullname, "register", "")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send verification email"))
	}

	response := response.ConvertToDoctorRegisterResponse(doctorRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("registered successful", response))
}

// Login Doctor
func LoginDoctorController(c echo.Context) error {
	var loginRequest web.DoctorLoginRequest

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid login data"))
	}

	if err := helper.ValidateStruct(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	var doctor schema.Doctor
	if err := configs.DB.Where("email = ? AND deleted_at IS NULL", loginRequest.Email).First(&doctor).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("email not registered"))
	}

	if err := helper.ComparePassword(doctor.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, helper.ErrorResponse("incorrect password"))
	}

	// The rest of your code for generating a token and handling the successful login
	token, err := middlewares.GenerateToken(doctor.ID, doctor.Email, doctor.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to generate jwt: " + err.Error()))
	}

	doctorLoginResponse := response.ConvertToDoctorLoginResponse(&doctor)
	doctorLoginResponse.Token = token

	// Send login notification email
	if doctor.Email != "" {
		notificationType := "login"
		if err := helper.SendNotificationEmail(doctor.Email, doctor.Fullname, notificationType, ""); err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to send notification email: " + err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("login successful", doctorLoginResponse))
}

func GetAvailableDoctor(c echo.Context) error {

	var Doctor []schema.Doctor

	err := configs.DB.Where("status = ?", true).Find(&Doctor).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve data"))
	}

	if len(Doctor) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("data not found"))
	}

	response := response.ConvertToGetAllDoctorResponse(Doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("data successfully retrieved", response))
}

func GetSpecializeDoctor(c echo.Context) error {
	specialist := c.QueryParam("specialist")

	if specialist == "" {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("parameter specialist required!"))
	}

	var doctors []schema.Doctor
	err := configs.DB.Where("specialist LIKE ?", "%"+specialist+"%").Find(&doctors).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve data"))
	}

	if len(doctors) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("data not found"))
	}

	response := response.ConvertToGetAllDoctorResponse(doctors)

	return c.JSON(http.StatusOK, helper.SuccessResponse("data successfully retrieved", response))
}

// Get Doctor Profile
func GetDoctorProfileController(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor id"))
	}

	var doctor schema.Doctor
	if err := configs.DB.First(&doctor, userID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("data not found"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor profile"))
		}
	}

	response := response.ConvertToGetDoctorResponse(&doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("data successfully retrieved", response))
}

// Get All Doctors
func GetAllDoctorController(c echo.Context) error {
	var doctors []schema.Doctor

	err := configs.DB.Find(&doctors).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve user data"))
	}

	if len(doctors) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("data is empty"))
	}

	response := response.ConvertToGetAllDoctorResponse(doctors)

	return c.JSON(http.StatusOK, helper.SuccessResponse("data successfully retrieved", response))
}

// Get All Doctors by Admin
func GetAllDoctorByAdminController(c echo.Context) error {
	var Doctor []schema.Doctor

	err := configs.DB.Find(&Doctor).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve data"))
	}

	if len(Doctor) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("data is empty"))
	}

	response := response.ConvertToGetAllDoctorByAdminResponse(Doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("data successfully retrieved", response))
}

// Update Doctor
func UpdateDoctorController(c echo.Context) error {
	// Get userID from the context
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to get doctor id"))
	}

	// Fetch the existing doctor based on userID
	var existingDoctor schema.Doctor
	result := configs.DB.First(&existingDoctor, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor data"))
	}

	// Parse the request body into the DoctorUpdateRequest struct
	var doctorUpdated web.DoctorUpdateRequest
	if err := c.Bind(&doctorUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid data input"))
	}

	// Check if the email already exists for another doctor
	var existingDoctorEmail schema.Doctor
	if existingEmail := configs.DB.Where("email = ? AND deleted_at IS NULL", doctorUpdated.Email).First(&existingDoctorEmail).Error; existingEmail == nil {
		return c.JSON(http.StatusConflict, helper.ErrorResponse("email already exists"))
	}

	// Validate the request payload
	if err := helper.ValidateStruct(doctorUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Hash the password if provided
	if doctorUpdated.Password != "" {
		doctorUpdated.Password = helper.HashPassword(doctorUpdated.Password)
	}

	// Parse multipart form for file upload
	err := c.Request().ParseMultipartForm(10 << 20)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Extract the image file from the form
	file, fileHeader, err := c.Request().FormFile("profile_picture")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("image file is required"))
	}
	defer file.Close()

	// Check if the file format is allowed
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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid image file format, supported formats: jpg, jpeg, png"))
	}

	// Upload the image to Cloud Storage
	ProfilePicture, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error uploading image to cloud storage"))
	}

	// Update the doctor details
	existingDoctor.ProfilePicture = ProfilePicture
	existingDoctor.Status = doctorUpdated.Status
	if err := configs.DB.Model(&existingDoctor).Updates(doctorUpdated).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to update data"))
	}

	configs.DB.Save(&existingDoctor)

	response := response.ConvertToDoctorUpdateResponse(&existingDoctor)
	return c.JSON(http.StatusOK, helper.SuccessResponse("data successfully updated", response))
}

// Update Doctor by Admin
func UpdateDoctorByAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid doctor id"))
	}

	var doctorUpdated web.DoctorUpdateRequest

	if err := c.Bind(&doctorUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid input data"))
	}

	if err := helper.ValidateStruct(doctorUpdated); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	if doctorUpdated.Password != "" {
		doctorUpdated.Password = helper.HashPassword(doctorUpdated.Password)
	}

	var existingDoctor schema.Doctor
	result := configs.DB.First(&existingDoctor, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to update data"))
	}

	configs.DB.Model(&existingDoctor).Updates(doctorUpdated)

	response := response.ConvertToDoctorUpdateResponse(&existingDoctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("doctor updated data successful", response))
}

// Delete Doctor
func DeleteDoctorController(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to get doctor id"))
	}

	var existingDoctor schema.Doctor
	result := configs.DB.First(&existingDoctor, userID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor data"))
	}

	if err := configs.DB.Delete(&existingDoctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to delete doctor"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("doctor account successfully deleted", nil))
}

// DeleteDoctorByAdminController deletes a doctor by admin
func DeleteDoctorByAdminController(c echo.Context) error {
	// Parse doctor ID from the request parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid id"))
	}

	// Retrieve the existing doctor from the database
	var existingDoctor schema.Doctor
	result := configs.DB.First(&existingDoctor, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve data"))
	}

	// Delete the doctor from the database
	result = configs.DB.Delete(&existingDoctor, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to delete data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("data deleted successfuly  ", nil))
}

// Get Doctor by ID
func GetDoctorByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("failed to retrieve doctor id"))
	}

	var doctor schema.Doctor
	result := configs.DB.First(&doctor, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to fetch doctor data"))
	}

	response := response.ConvertToGetIDDoctorResponse(&doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("doctor details successfully retrieved", response))
}

// Manage User
func GetManagePatientController(c echo.Context) error {
	doctorID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid user ID"))
	}

	transactionID, _ := strconv.Atoi(c.QueryParam("transaction_id"))
	patientStatus := c.QueryParam("patient_status")

	var manageUser []schema.DoctorTransaction

	var err error

	if transactionID != 0 {
		// Get transaction by ID
		err = configs.DB.First(&manageUser, "doctor_id = ? AND id = ?", doctorID, transactionID).Error
	} else if patientStatus != "" {
		// Get transactions by patient status
		err = configs.DB.Find(&manageUser, "doctor_id = ? AND patient_status = ?", doctorID, patientStatus).Error
	} else {
		// Get all transactions
		err = configs.DB.Where("deleted_at IS NULL").Find(&manageUser, "doctor_id=?", doctorID).Error
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve doctor transaction data"))
	}

	if len(manageUser) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(fmt.Sprintf("no doctor transaction data found for doctorid: %d, transactionID: %d, patientStatus: %s", doctorID, transactionID, patientStatus)))
	}

	var responses []web.ManageUserResponse
	for _, doctorTransaction := range manageUser {
		var user schema.User
		err := configs.DB.First(&user, "id=?", doctorTransaction.UserID).Error
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve user data"))
		}

		response := response.ConvertToManageUserResponse(doctorTransaction, user)
		responses = append(responses, response)
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("doctor transaction data successfully retrieved", responses))
}

// Update manage user
func UpdateManagePatientController(c echo.Context) error {
	// Getting the doctor ID from the context
	doctorID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("invalid user id"))
	}

	var requestBody web.UpdateManageUserRequest

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("invalid request body"))
	}
	if err := helper.ValidateStruct(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	// Checking if the required fields have been provided
	if requestBody.HealthDetails == "" && requestBody.PatientStatus == "" {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("health details or patient status is required"))
	}

	transactionID, _ := strconv.Atoi(c.QueryParam("transaction_id"))

	var doctorTransaction schema.DoctorTransaction
	// Getting doctor transaction based on doctor ID and transaction ID
	err := configs.DB.First(&doctorTransaction, "doctor_id = ? AND id = ?", doctorID, transactionID).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("doctor transaction not found"))
	}

	// Updating health details and patient status if provided
	if requestBody.HealthDetails != "" {
		doctorTransaction.HealthDetails = requestBody.HealthDetails
	}

	if requestBody.PatientStatus != "" {
		doctorTransaction.PatientStatus = requestBody.PatientStatus
	}

	// Saving the updated doctor transaction to the database
	if err := configs.DB.Save(&doctorTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to update health details and patient status"))
	}

	// Getting user data
	var user schema.User
	err = configs.DB.First(&user, "id=?", doctorTransaction.UserID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to retrieve user data"))
	}
	response := response.ConvertToManageUserResponse(doctorTransaction, user)

	return c.JSON(http.StatusOK, helper.SuccessResponse("health details and patient status successfully updated", response))
}


