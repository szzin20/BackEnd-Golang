package web

type DoctorResgisterResponse struct {
	Fullname           string `json:"fullname" form:"fullname"`
	Email              string `json:"email" form:"email"`
	Price              int    `json:"price" form:"price"`
	Tag                string `json:"tag" form:"tag"`
	ProfilePicture     string `json:"profile_picture" form:"profile_picture"`
	RegistrationLetter string `json:"registration_letter" form:"registration_letter"`
}

type DoctorLoginResponse struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Token    string `json:"token" form:"token"`
}

type DoctorUpdateResponse struct {
	Fullname           string `json:"fullname" form:"fullname"`
	Email              string `json:"email" form:"email"`
	Price              int    `json:"price" form:"price"`
	Tag                string `json:"tag" form:"tag"`
	ProfilePicture     string `json:"profile_picture" form:"profile_picture"`
	RegistrationLetter string `json:"registration_letter" form:"registration_letter"`
}

type DoctorAllResponse struct {
	Fullname           string `json:"fullname" form:"fullname"`
	Email              string `json:"email" form:"email"`
	Price              int    `json:"price" form:"price"`
	Tag                string `json:"tag" form:"tag"`
	Status             string `json:"status" form:"status"`
	ProfilePicture     string `json:"profile_picture" form:"profile_picture"`
	RegistrationLetter string `json:"registration_letter" form:"registration_letter"`
}
