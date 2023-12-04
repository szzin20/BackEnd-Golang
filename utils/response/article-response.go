package response

import (
	"healthcare/configs"
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToGetAllArticles(articles []schema.Article) []web.ArticleOnlyResponses {
	var results []web.ArticleOnlyResponses

	for _, article := range articles {
		articleResponse := web.ArticleOnlyResponses{
			ID:        article.ID,
			Title:     article.Title,
			Content:   article.Content,
			Image:     article.Image,
			CreatedAt: article.CreatedAt,
		}

		results = append(results, articleResponse)
	}

	return results
}

func ConvertToDoctorArticles(doctors schema.Doctor) web.DoctorArticle {
	return web.DoctorArticle{
		DoctorID:       doctors.ID,
		Fullname:       doctors.Fullname,
		ProfilePicture: doctors.ProfilePicture,
	}
}

func ConvertToSliceDoctorsArticles(doctors []schema.Doctor) []web.DoctorArticle {
	var result []web.DoctorArticle
	for _, v := range doctors {
		result = append(result, ConvertToDoctorArticles(v))
	}
	return result
}

func ConvertToGetArticleResponse(article *schema.Article) web.ArticleOnlyResponses {
	return web.ArticleOnlyResponses{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Image:     article.Image,
		CreatedAt: article.CreatedAt,
	}
}

// Fungsi untuk mendapatkan data dokter berdasarkan ID
func getDoctorByID(doctorID uint) (schema.Doctor, error) {
	var doctor schema.Doctor
	if err := configs.DB.First(&doctor, doctorID).Error; err != nil {
		return doctor, err
	}

	return doctor, nil
}

func ConvertToArticleDoctor(article schema.Article) web.DoctorArticleResponse {
	doctor, err := getDoctorByID(article.DoctorID)
	if err != nil {
		return web.DoctorArticleResponse{}
	}

	return web.DoctorArticleResponse{
		ID:             article.ID,
		Title:          article.Title,
		Content:        article.Content,
		Image:          article.Image,
		CreatedAt:      article.CreatedAt,
		Fullname:       doctor.Fullname,
		ProfilePicture: doctor.ProfilePicture,
	}
}

func ListConvertToArticleDoctors(articles []schema.Article) []web.DoctorArticleResponse {
	articleDoctors := make([]web.DoctorArticleResponse, len(articles))

	for i, article := range articles {
		articleDoctors[i] = ConvertToArticleDoctor(article)
	}

	return articleDoctors
}
