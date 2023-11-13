package web

type DoctorRegisterRequest struct {
	Fullname           string `json:"fullname" form:"fullname"`
	Email              string `json:"email" form:"email"`
	Password           string `json:"password" form:"password"`
	Price              int    `json:"price" form:"price"`
	Tag                string `json:"tag" form:"tag"`
	ImageURL           string `json:"image_url" form:"image_url"`
	RegistrationLetter string `json:"registration_letter" form:"registration_letter"`
}

type DoctorLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type DoctorUpdateRequest struct {
	Fullname           string `json:"fullname" form:"fullname"`
	Email              string `json:"email" form:"email"`
	Password           string `json:"password" form:"password"`
	Price              int    `json:"price" form:"price"`
	Tag                string `json:"tag" form:"tag"`
	ImageURL           string `json:"image_url" form:"image_url"`
	RegistrationLetter string `json:"registration_letter" form:"registration_letter"`
}
