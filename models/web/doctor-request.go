package web

type DoctorRegisterRequest struct {
	Fullname           string `json:"fullname" form:"fullname"`
	Email              string `json:"email" form:"email"`
	Password           string `json:"password" form:"password"`
	Price              int    `json:"price" form:"price"`
	Tag                string `json:"tag" form:"tag"`
	ProfilePicture     string `json:"profile_picture" form:"profile_picture"`
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
	ProfilePicture     string `json:"profile_picture" form:"profile_picture"`
	RegistrationLetter string `json:"registration_letter" form:"registration_letter"`
}
