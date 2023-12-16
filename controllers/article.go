package controllers

import (
	"fmt"
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
	"healthcare/utils/helper"
	"healthcare/utils/helper/constanta"
	"healthcare/utils/request"
	"healthcare/utils/response"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateArticle(c echo.Context) error {
	userID := c.Get("userID").(int)

	var article web.CreateArticle

	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidBody))
	}
	if err := helper.ValidateStruct(article); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	file, fileHeader, err := c.Request().FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrImageFileRequired))
	}
	defer file.Close()

	if fileHeader.Size > 10*1024*1024 { // 10 MB limit
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("image file size exceeds the limit (10 MB)"))
	}

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
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidImageFormat))
	}

	image, err := helper.UploadFilesToGCS(c, fileHeader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error upload image to Cloud Storage"))
	}

	articleRequest := request.ConvertToCreateArticleRequest(article)

	articleRequest.Image = image
	articleRequest.DoctorID = uint(userID)

	if err := configs.DB.Create(&articleRequest).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("failed to create article"))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse(constanta.SuccessActionCreated+"article", nil))
}

func UpdateArticleById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidIDParam))
	}

	userID := c.Get("userID").(int)

	var existingArticle schema.Article
	result := configs.DB.First(&existingArticle, id, userID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound))
	}

	var updateArticle web.CreateArticle

	if err := c.Bind(&updateArticle); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidBody))
	}
	if err := helper.ValidateStruct(updateArticle); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	if file, fileHeader, err := c.Request().FormFile("image"); err == nil {
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
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidImageFormat))
		}

		if fileHeader.Size > 10*1024*1024 { // 10 MB limit
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("image file size exceeds the limit (10 MB)"))
		}

		image, err := helper.UploadFilesToGCS(c, fileHeader)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("error upload image to Cloud Storage"))
		}

		updateArticle.Image = image

	} else if err != http.ErrMissingFile {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrImageFileRequired))

	} else {
		updateArticle.Image = existingArticle.Image
	}

	result = configs.DB.Model(&existingArticle).Updates(updateArticle)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionUpdated+"article"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionUpdated+"article", nil))

}

func GetAllArticlesPagination(offset int, limit int, queryInput []schema.Article) ([]schema.Article, int64, error) {

	if offset < 0 || limit < 0 {
		return nil, 0, nil
	}

	queryAll := queryInput
	var total int64

	query := configs.DB.Model(&queryAll)

	query.Find(&queryAll).Count(&total)

	query = query.Order("created_at desc").Limit(limit).Offset(offset)

	result := query.Find(&queryAll)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	if offset >= int(total) {
		return nil, 0, fmt.Errorf("not found")
	}

	return queryAll, total, nil
}

func GetAllArticles(c echo.Context) error {
	params := c.QueryParams()
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("limit"+constanta.ErrQueryParamRequired))
	}

	offset, err := strconv.Atoi(params.Get("offset"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("offset"+constanta.ErrQueryParamRequired))
	}

	var articles []schema.Article

	article, total, err := GetAllArticlesPagination(offset, limit, articles)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound))
		}
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	pagination := helper.Pagination(offset, limit, total)

	response := response.ListConvertToArticleDoctors(article)

	return c.JSON(http.StatusOK, helper.PaginationResponse(constanta.SuccessActionGet+"articles", response, pagination))
}

func GetArticleByID(c echo.Context) error {
	articleID, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidIDParam))
	}

	var article schema.Article

	if err = configs.DB.First(&article, articleID).Error; err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound))
	}

	response := response.ConvertToArticleDoctor(article)

	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionGet, response))
}

func GetAllArticlesByTitle(c echo.Context) error {
	param := c.QueryParam("title")
	if param == "" {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("title"+constanta.ErrQueryParamRequired))
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("limit"+constanta.ErrQueryParamRequired))
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("offset"+constanta.ErrQueryParamRequired))
	}

	var articles []schema.Article
	var total int64

	query := configs.DB.Where("title LIKE ?", "%"+param+"%")

	query.Model(&articles).Count(&total)

	query = query.Limit(limit).Offset(offset)

	err = query.Find(&articles).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound))
	}

	if len(articles) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound))
	}

	pagination := helper.Pagination(offset, limit, total)
	response := response.ListConvertToArticleDoctors(articles)

	return c.JSON(http.StatusOK, helper.PaginationResponse(constanta.SuccessActionGet+"article", response, pagination))
}

func DeleteArticleById(c echo.Context) error {
	userID := c.Get("userID").(int)

	articleID, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ErrInvalidIDParam))
	}

	var existingArticle schema.Article
	result := configs.DB.First(&existingArticle, articleID, userID)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound))
	}

	if err := configs.DB.Delete(&existingArticle).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionDeleted+"article"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionDeleted+"article", nil))
}

// get all for doctor
// get by id for doctor

func DoctorGetAllArticles(c echo.Context) error {
	userID := c.Get("userID").(int)

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("limit"+constanta.ErrQueryParamRequired))
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("offset"+constanta.ErrQueryParamRequired))
	}

	var articles []schema.Article
	var total int64

	query := configs.DB.Where("doctor_id = ?", userID)

	query.Model(&articles).Count(&total)

	query = query.Order("created_at desc").Limit(limit).Offset(offset)

	err = query.Find(&articles).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse(constanta.ErrActionGet+"articles"))
	}

	if len(articles) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound))
	}

	response := response.ConvertToGetAllArticles(articles)
	pagination := helper.Pagination(offset, limit, total)

	return c.JSON(http.StatusOK, helper.PaginationResponse(constanta.SuccessActionGet+"articles", response, pagination))
}

func DoctorGetArticleByID(c echo.Context) error {
	userID := c.Get("userID").(int)

	articleID := c.Param("article_id")

	var article schema.Article
	if err := configs.DB.Where("id = ? AND doctor_id = ?", articleID, userID).First(&article).Error; err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ErrNotFound))
	}

	response := response.ConvertToGetArticleResponse(&article)

	return c.JSON(http.StatusOK, helper.SuccessResponse(constanta.SuccessActionGet+"article", response))
}
