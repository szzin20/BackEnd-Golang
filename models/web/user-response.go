package web

type UserRegisterResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserLoginResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserUpdateResponse struct {
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
	Gender         string `json:"gender"`
	Birthdate      string `json:"birthdate"`
	BloodType      string `json:"blood_type"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
}

type UserAllResponseByAdmin struct {
	ID             uint   `json:"id" form:"id"`
	Fullname       string `json:"fullname" form:"fullname"`
	Email          string `json:"email" form:"email"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Gender         string `json:"gender" form:"gender"`
	Birthdate      string `json:"birthdate" form:"birthdate"`
	BloodType      string `json:"blood_type" form:"blood_type"`
	Height         int    `json:"height" form:"height"`
	Weight         int    `json:"weight" form:"weight"`
	// DoctorTransaction []DoctorTransaction `gorm:"ForeignKey:DoctorID;references:ID"`
}
