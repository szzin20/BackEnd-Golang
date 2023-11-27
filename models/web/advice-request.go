package web

type AdviceRequest struct {
	Message string `json:"message" form:"message" validate:"omitempty"`
	Image   string `json:"image" form:"image" validate:"omitempty"`
	Audio   string `json:"audio" form:"audio" validate:"omitempty"`
}
