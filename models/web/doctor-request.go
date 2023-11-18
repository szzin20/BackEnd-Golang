package web

type DoctorRegisterRequest struct {
	ProfilePicture string `json:"profile_picture" form:"profile_picture" `
	Fullname       string `json:"fullname" form:"fullname" validate:"required"`
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"min=8",max15"`
	Price          int    `json:"price" form:"price" validate:"min=0"`
	Gender         string `json:"gender" form:"gender" validate:"required"`
	Specialist     string `json:"specialist" form:"specialist" validate:"required"`
	NoSTR          int    `json:"no_str" form:"no_str" validate:"required"`
	Experience     string `json:"experience" form:"experience" validate:"required"`
}

type DoctorLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=15"`
}

type DoctorUpdateRequest struct {
	Fullname       string `json:"fullname" form:"fullname" validate:"required"`
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"min=8"`
	Price          int    `json:"price" form:"price" validate:"min=0"`
	Gender         string `json:"gender" form:"gender" validate:"required"`
	Specialist     string `json:"specialist" form:"specialist" validate:"required"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Experience     string `json:"experience" form:"experience" validate:"required"`
}

