package web

type ComplaintRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}
