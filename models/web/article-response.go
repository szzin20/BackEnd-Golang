package web

type ArticleResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Image    string `json:"image"`
	DoctorID uint   `json:"doctor_id"`
}
