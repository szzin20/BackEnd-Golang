package web

type DoctorResgisterResponse struct {
	Name               string `json:"name" form:"name"`
	Email              string `json:"email" form:"email"`
	Password           string `json:"password" form:"password"`
	Status             bool   `json:"status" form:"status"`
	Price              int    `json:"price" form:"price"`
	Tag                string `json:"tag" form:"tag"`
	ImageURL           string `json:"image_url" form:"image_url"`
	RegistrationLetter string `json:"registration_letter" form:"registration_letter"`
}

type DoctorLoginResponse struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}

type DoctorUpdateResponse struct {
	Name               string `json:"name" form:"name"`
	Email              string `json:"email" form:"email"`
	Password           string `json:"password" form:"password"`
	Status             bool   `json:"status" form:"status"`
	Price              int    `json:"price" form:"price"`
	Tag                string `json:"tag" form:"tag"`
	ImageURL           string `json:"image_url" form:"image_url"`
	RegistrationLetter string `json:"registration_letter" form:"registration_letter"`
}
