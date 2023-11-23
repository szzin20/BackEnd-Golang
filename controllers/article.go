package controllers

import (
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

func CreateArticle(c echo.Context) error {
	userID := c.Get("userID").(int)

	var article web.CreateArticle

	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Request Body Tidak Valid!"))
	}
	if err := helper.ValidateStruct(article); err != nil {
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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid image file format. Supported formats: jpg, jpeg, png"))
	}

	image, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error upload image to Cloud Storage"))
	}

	articleRequest := request.ConvertToCreateArticleRequest(article)

	articleRequest.Image = image
	articleRequest.DoctorID = uint(userID)

	if err := configs.DB.Create(&articleRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal Membuat Article!"))
	}

	response := response.ConvertToCreateArticleResponse(articleRequest)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Article berhasil dibuat!", response))
}

func UpdateArticleById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Medicine ID"))
	}

	userID := c.Get("userID").(int)

	var existingArticle schema.Article
	result := configs.DB.First(&existingArticle, id, userID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data artikel tidak ditemukan"))
	}

	var updateArticle web.CreateArticle

	if err := c.Bind(&updateArticle); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Request Body Tidak Valid!"))
	}
	if err := helper.ValidateStruct(updateArticle); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
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

	image, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error upload image to Cloud Storage"))
	}

	updateArticle.Image = image

	result = configs.DB.Model(&existingArticle).Updates(updateArticle)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal update article"))
	}

	response := response.ConvertToCreateArticleResponse(&existingArticle)
	return c.JSON(http.StatusOK, helper.SuccessResponse("Artikel berhasil diperbarui", response))

}

func GetAllArticles(c echo.Context) error {
	var articles []schema.Article

	err := configs.DB.Find(&articles).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Data artikel kosong"))
	}

	if len(articles) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data artikel kosong"))
	}

	response := response.ConvertToGetAllArticles(articles)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data Artikel Berhasil Diambil!", response))
}

func GetArticleByID(c echo.Context) error {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Article ID Tidak Valid"))
	}

	var article schema.Article
	if err = configs.DB.First(&article, articleID).Error; err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Artikel tidak ditemukan"))
	}

	response := response.ConvertToGetArticleResponse(&article)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Artikel berhasil diambil", response))
}

func GetAllArticlesByTitle(c echo.Context) error {

	param := c.QueryParam("title")
	if param == "" {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Parameter title required!"))
	}

	var articles []schema.Article
	err := configs.DB.Where("title LIKE ?", "%"+param+"%").Find(&articles).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data artikel tidak ditemukan"))
	}

	if len(articles) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data artikel tidak ditemukan"))
	}

	response := response.ConvertToGetAllArticles(articles)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data Artikel Berhasil Diambil!", response))
}

func DeleteArticleById(c echo.Context) error {
	userID := c.Get("userID")

	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Article ID Tidak Valid"))
	}

	var existingArticle schema.Article
	result := configs.DB.First(&existingArticle, articleID, userID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Data artikel tidak ditemukan"))
	}

	if err := configs.DB.Delete(&existingArticle).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal menghapus artikel"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Artikel berhasil dihapus", nil))
}

// get all for doctor
// get by id for doctor

func DoctorGetAllArticles(c echo.Context) error {
	userID := c.Get("userID")

	var articles []schema.Article

	err := configs.DB.Where("doctor_id = ?", userID).Find(&articles).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil data artikel"))
	}

	if len(articles) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Artikel Tidak Ditemukan"))
	}

	response := response.ConvertToGetAllArticles(articles)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Data Artikel Berhasil Diambil!", response))
}

func DoctorGetArticleByID(c echo.Context) error {
	userID := c.Get("userID")

	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Article ID Tidak Valid"))
	}

	var article schema.Article
	if err = configs.DB.First(&article, articleID, userID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Gagal mengambil Article by ID"))
	}

	response := response.ConvertToGetArticleResponse(&article)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Artikel berhasil diambil", response))
}
