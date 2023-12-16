package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateArticleRequest(article web.CreateArticle) *schema.Article {
	return &schema.Article{
		Title:   article.Title,
		Content: article.Content,
	}
}
