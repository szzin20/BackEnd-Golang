package controllers

import (
	"fmt"
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create Doctor Transaction
func CreateDoctorTransaction(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid User ID"))
	}

	doctorID, _ := strconv.Atoi(c.QueryParam("doctor_id"))

	var doctorTransactionRequest web.CreateDoctorTransactionRequest

	if err := c.Bind(&doctorTransactionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input Doctor Transaction Data"))
	}

	if err := helper.ValidateStruct(doctorTransactionRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

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

	paymentConfirmation, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error upload image to Cloud Storage"))
	}

	doctorTransactionRequest.PaymentConfirmation = paymentConfirmation

	var doctor schema.Doctor

	if err := configs.DB.First(&doctor, doctorID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create Doctor Transaction"))
	}

	doctorTransaction := request.ConvertToCreateDTRequest(doctorTransactionRequest, uint(userID), uint(doctorID), doctor.Fullname, doctor.Specialist, doctor.Price)

	if err := configs.DB.Create(&doctorTransaction).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create Doctor Transaction"))
	}

	response := response.ConvertToCreateDTResponse(doctorTransaction, doctor)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Doctor Transaction Created Successful", response))
}

// Get Doctor Transaction by ID
func GetDoctorTransactionController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid User ID"))
	}

	transactionID, _ := strconv.Atoi(c.QueryParam("transaction_id"))

	var doctorTransaction schema.DoctorTransaction

	if err := configs.DB.First(&doctorTransaction, userID, "id = ?", transactionID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Doctor Transaction Data"))
	}
	
	var doctor schema.Doctor
	if err := configs.DB.First(&doctor, "id = ?", doctorTransaction.DoctorID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Mengambil Data Dokter"))
	}

	response := response.ConvertToCreateDTResponse(&doctorTransaction, doctor)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Doctor Transaction Data Successfully Retrieved", response))

}

// Get All Doctor Transactions
func GetAllDoctorTransactionsController(c echo.Context) error {

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid User ID"))
	}
	fmt.Println(userID)

	var doctorTransaction []schema.DoctorTransaction

	err := configs.DB.Where("deleted_at IS NULL").Find(&doctorTransaction, "user_id=?",userID).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Mengambil Data Transaksi Dokter"))
	}

	var responses []web.CreateDoctorTransactionResponse 
	for i, doctor_id := range doctorTransaction{
		fmt.Println(i)
		var doctor schema.Doctor
		err := configs.DB.Find(&doctor,"id=?" ,doctor_id.DoctorID).Error
		if err != nil{
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Mengambil Data Dokter"))
		}

		if len(doctorTransaction) == 0 {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data Transaksi Dokter Kosong"))
		}

		responses = append(responses,response.ConvertToGetAllDTResponse(doctorTransaction[i], doctor))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data Transaksi Dokter Berhasil Diambil", responses))
}

