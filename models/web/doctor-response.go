package web


type DoctorLoginResponse struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Token    string `json:"token" form:"token"`
}

type DoctorUpdateResponse struct {
	Fullname       string `json:"fullname" form:"fullname" `
	Email          string `json:"email" form:"email"`
	Gender         string `json:"gender" form:"gender" `
	Specialist     string `json:"tag" form:"tag" `
	ProfilePicture string `json:"profile_picture" form:"profile_picture" `
	NoSTR          int    `json:"registration_certificate" form:"registration_certificate" `
	Experience     string `json:"experience" form:"experience" `
	Alumnus        string `json:"alumnus" form:"alumnus" `
	Status         bool   `json:"status" form:"status"`
}

type DoctorAllResponse struct {
	Fullname   string `json:"fullname" form:"fullname" `
	Price      int    `json:"price" form:"price" `
	Gender     string `json:"gender" form:"gender"`
	Status     bool   `json:"status" form:"status"`
	Specialist string `json:"specialist" form:"specialist" `
	Experience string `json:"experience" form:"experience" `
	Alumnus    string `json:"alumnus" form:"alumnus" `
}

type DoctorResponse struct {
	Fullname       string `json:"fullname" form:"fullname" `
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password" `
	Gender         string `json:"gender" form:"gender" `
	Specialist     string `json:"specialist" form:"specialist" `
	ProfilePicture string `json:"profile_picture" form:"profile_picture" `
	NoSTR          int    `json:"no_str" form:"no_str" `
	Experience     string `json:"experience" form:"experience" `
	Alumnus        string `json:"alumnus" form:"alumnus" `
}

