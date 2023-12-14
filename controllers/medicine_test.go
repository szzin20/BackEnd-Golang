package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"healthcare/configs"
	"healthcare/models/web"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func InitTestDB() *echo.Echo {
	e := echo.New()
	configs.ConnectDBTest()
	return e
}

func TestGetMedicineControllerValid(t *testing.T) {
	e := InitTestDB()
	offset := 0
	limit := 10
	url := fmt.Sprintf("/users/medicines?offset=%d&limit=%d", offset, limit)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)
	err := GetMedicineUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetMedicineControllerNotFound(t *testing.T) {
	e := InitTestDB()
	offset := 99
	limit := 10
	url := fmt.Sprintf("/users/medicines?offset=%d&limit=%d", offset, limit)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)
	err := GetMedicineUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetMedicineControllerInvalidOffset(t *testing.T) {
	e := InitTestDB()
	limit := 10
	url := fmt.Sprintf("/users/medicines?limit=%d", limit)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)
	err := GetMedicineUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetMedicineControllerInvalidLimit(t *testing.T) {
	e := InitTestDB()
	offset := 0
	url := fmt.Sprintf("/users/medicines?offset=%d", offset)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)
	err := GetMedicineUserController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetMedicineByIDControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/users/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("3")
	err := GetMedicineUserByIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetMedicineByIDControllerInvalidID(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/users/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("invalid_id")
	err := GetMedicineUserByIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetMedicineByIDControllerNotFound(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/users/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("999")
	err := GetMedicineUserByIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetMedicineAdminControllerValid(t *testing.T) {
	e := InitTestDB()
	offset := 0
	limit := 10
	url := fmt.Sprintf("/admins/medicines?offset=%d&limit=%d", offset, limit)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)
	err := GetMedicineAdminController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetMedicineAdminControllerNotFound(t *testing.T) {
	e := InitTestDB()
	offset := 99
	limit := 10
	url := fmt.Sprintf("/admins/medicines?offset=%d&limit=%d", offset, limit)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)
	err := GetMedicineAdminController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetMedicineAdminControllerInvalidOffset(t *testing.T) {
	e := InitTestDB()
	limit := 10
	url := fmt.Sprintf("/admins/medicines?limit=%d", limit)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)
	err := GetMedicineAdminController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetMedicineAdminControllerInvalidLimit(t *testing.T) {
	e := InitTestDB()
	offset := 0
	url := fmt.Sprintf("/admins/medicines?offset=%d", offset)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	fmt.Println(rec.Code)
	err := GetMedicineAdminController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetMedicineAdminByIDControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/admins/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("1")
	err := GetMedicineAdminByIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetMedicineAdminByIDControllerInvalidID(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/admins/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("invalid_id")
	err := GetMedicineAdminByIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetMedicineAdminByIDControllerNotFound(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/admins/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("999")
	err := GetMedicineAdminByIDController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestCreateMedicineControllerBadRequest(t *testing.T) {
	e := InitTestDB()

	// Create a sample MedicineRequest
	medicineRequest := web.MedicineRequest{
		Code: "ABC123",
	}

	_, err := json.Marshal(medicineRequest)
	assert.NoError(t, err)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	imageFile, err := os.Open("../image/paracetamol.jpg")
	assert.NoError(t, err)
	defer imageFile.Close()

	part, err := writer.CreateFormFile("image", "paracetamol.jpg")
	assert.NoError(t, err)
	_, err = io.Copy(part, imageFile)
	assert.NoError(t, err)

	err = writer.Close()
	assert.NoError(t, err)

	url := "/admins/medicines"

	req := httptest.NewRequest(http.MethodPost, url, body)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetRequest(req)

	err = CreateMedicineController(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateMedicineAdminControllerValid(t *testing.T) {
	e := InitTestDB()
	updateData := web.MedicineUpdateRequest{
		Name: "Updated Name",
	}
	req := httptest.NewRequest(http.MethodPut, "/admins/medicines", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	body, err := json.Marshal(updateData)
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("1")
	err = UpdateMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateMedicineAdminControllerInvalidID(t *testing.T) {
	e := InitTestDB()
	updateData := web.MedicineUpdateRequest{
		Name: "Updated Name",
	}
	req := httptest.NewRequest(http.MethodPut, "/admins/medicines", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	body, err := json.Marshal(updateData)
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("invalid_id")
	err = UpdateMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateMedicineAdminControllerNotFound(t *testing.T) {
	e := InitTestDB()
	updateData := web.MedicineUpdateRequest{
		Name: "Updated Name",
	}
	req := httptest.NewRequest(http.MethodPut, "/admins/medicines", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	body, err := json.Marshal(updateData)
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("999")
	err = UpdateMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestUpdateImageMedicineAdminControllerInvalidID(t *testing.T) {
	e := InitTestDB()
	updateData := web.MedicineUpdateRequest{
		Name: "Updated Name",
	}
	req := httptest.NewRequest(http.MethodPut, "/admins/:medicine_id/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	body, err := json.Marshal(updateData)
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("invalid_id")
	err = UpdateImageMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateImageMedicineAdminControllerNotFound(t *testing.T) {
	e := InitTestDB()
	updateData := web.MedicineUpdateRequest{
		Name: "Updated Name",
	}
	req := httptest.NewRequest(http.MethodPut, "/admins/:medicine_id/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	body, err := json.Marshal(updateData)
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("999")
	err = UpdateImageMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteMedicineAdminControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/admins/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("7")
	err := DeleteMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteMedicineAdminControllerInvalidID(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/admins/:medicine_id/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("invalid_id")
	err := DeleteMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestDeleteMedicineAdminControllerNotFound(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/admins/:medicine_id/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("999")
	err := DeleteMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteImageMedicineAdminControllerInvalidID(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/admins/:medicine_id/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("invalid_id")
	err := DeleteImageMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestDeleteImageMedicineAdminControllerNotFound(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/admins/:medicine_id/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("999")
	err := DeleteImageMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteImageMedicineAdminByIDControllerInternalServerError(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodDelete, "/admins/:medicine_id/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("7")
	err := DeleteImageMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestGetImageMedicineAdminByIDControllerValid(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/admins/:medicine_id/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("7")
	err := GetImageMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetImageMedicineAdminByIDControllerInvalidID(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/admins/:medicine_id/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("invalid_id")
	err := GetImageMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetImageMedicineAdminByIDControllerNotFound(t *testing.T) {
	e := InitTestDB()
	req := httptest.NewRequest(http.MethodGet, "/admins/:medicine_id/medicines/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluMTIzQGdtYWlsLmNvbSIsImV4cCI6MTcwMjgzMTAzMiwiaWQiOjEsInJvbGUiOiJhZG1pbiJ9.AbfCi12gYEE88p_bsM3vdfU_v6RRjXawVBPnTsc5z5I")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:medicine_id/")
	c.SetParamNames("medicine_id")
	c.SetParamValues("999")
	err := GetImageMedicineController(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
