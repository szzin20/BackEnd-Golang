package web

type AdviceRequest struct {
	Message string `json:"message" form:"message"`
	Image   string `json:"image" form:"image"`
	Audio   string `json:"audio" form:"audio"`
}
