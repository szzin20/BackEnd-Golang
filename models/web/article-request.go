package web

type CreateArticle struct {
	Title   string `json:"title" form:"title" validate:"required,min=1,max=50"`
	Content string `json:"content" form:"content" validate:"required,min=1,max=3000"`
	Image   string `json:"image"`
}
