package web

type ChatbotRequest struct {
	Request string `json:"request" validate:"required"`
}
