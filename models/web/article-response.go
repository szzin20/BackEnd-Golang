package web

import "time"

type ArticleResponse struct {
	ID        uint            `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Image     string          `json:"image"`
	CreatedAt time.Time       `json:"created_at"`
	Doctor    []DoctorArticle `json:"doctor"`
}

type DoctorArticle struct {
	DoctorID       uint   `json:"doctor_id"`
	Fullname       string `json:"fullname"`
	ProfilePicture string `json:"profile_picture"`
}

type DoctorArticleResponse struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Image          string    `json:"image"`
	CreatedAt      time.Time `json:"created_at"`
	Fullname       string    `json:"fullname"`
	ProfilePicture string    `json:"profile_picture"`
}

type ArticleOnlyResponses struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}
