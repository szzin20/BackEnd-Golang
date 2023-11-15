package web

type UserRegisterRequest struct {
	Fullname  string `json:"fullname" form:"fullname"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserUpdateRequest struct {
	Fullname  string `json:"fullname" form:"fullname"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	ImageURL  string `json:"image_url" form:"image_url"`
	Gender    string `json:"gender" form:"gender"`
	Birthdate string `json:"birthdate" form:"birthdate"`
	BloodType string `json:"blood_type" form:"blood_type"`
	Height    int    `json:"height" form:"height"`
	Weight    int    `json:"weight" form:"weight"`
}
