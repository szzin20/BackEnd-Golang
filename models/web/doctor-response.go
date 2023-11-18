package web

type DoctorLoginResponse struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Token    string `json:"token" form:"token"`
}

type DoctorUpdateResponse struct {
	Fullname       string `json:"fullname" form:"fullname"`
	Email          string `json:"email" form:"email"`
	Gender         string `json:"gender" form:"gender"`
	Specialist     string `json:"specialist" form:"specialist"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	NoSTR          int    `json:"no_str" form:"no_str"`
	Experience     string `json:"experience" form:"experience"`
	Alumnus        string `json:"alumnus" form:"alumnus"`
	Status         bool   `json:"status" form:"status"`
}

type DoctorAllResponse struct {
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Fullname       string `json:"fullname" form:"fullname"`
	NoSTR          int    `json:"no_str" form:"no_str"`
	Price          int    `json:"price" form:"price"`
	Gender         string `json:"gender" form:"gender"`
	Status         bool   `json:"status" form:"status"`
	Specialist     string `json:"specialist" form:"specialist"`
	Experience     string `json:"experience" form:"experience"`
	Alumnus        string `json:"alumnus" form:"alumnus"`
}