package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateArticleResponse(article *schema.Article) web.ArticleResponse {
	return web.ArticleResponse{
		ID:       article.ID,
		Title:    article.Title,
		Content:  article.Content,
		Image:    article.Image,
		DoctorID: article.DoctorID,
	}
}

func ConvertToGetAllArticles(articles []schema.Article) []web.ArticleResponse {
	var results []web.ArticleResponse

	for _, article := range articles {
		articleResponse := web.ArticleResponse{
			ID:       article.ID,
			Title:    article.Title,
			Content:  article.Content,
			Image:    article.Image,
			DoctorID: article.DoctorID,
		}

		results = append(results, articleResponse)
	}

	return results
}

func ConvertToGetArticleResponse(article *schema.Article) web.ArticleResponse {
	return web.ArticleResponse{
		ID:       article.ID,
		Title:    article.Title,
		Content:  article.Content,
		Image:    article.Image,
		DoctorID: article.DoctorID,
	}
}
