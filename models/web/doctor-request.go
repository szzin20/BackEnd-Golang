package web

type DoctorRegisterRequest struct {
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Fullname       string `json:"fullname" form:"fullname" validate:"required"`
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"required,min=10,max=15"`
	Price          int    `json:"price" form:"price" validate:"required,min=0"`
	Gender         string `json:"gender" form:"gender" validate:"required"`
	Specialist     string `json:"specialist" form:"specialist" validate:"required"`
	NoSTR          int    `json:"no_str" form:"no_str" validate:"required"`
	Experience     string `json:"experience" form:"experience" validate:"required"`
	Alumnus        string `json:"alumnus" form:"alumnus" validate:"required"`
}

type DoctorLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=10,max=15"`
}

type DoctorUpdateRequest struct {
	Fullname       string `json:"fullname" form:"fullname" validate:"omitempty,min=3"`
	Email          string `json:"email" form:"email" validate:"omitempty,email"`
	Password       string `json:"password" form:"password" validate:"omitempty,min=10,max=15"`
	Gender         string `json:"gender" form:"gender" validate:"omitempty"`
	Specialist     string `json:"specialist" form:"specialist" validate:"omitempty"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture" validate:"omitempty"`
	Status         bool   `json:"status" form:"status" validate:"omitempty,eq=true|eq=false"`
	Experience     string `json:"experience" form:"experience" validate:"omitempty"`
	Alumnus        string `json:"alumnus" form:"alumnus" validate:"omitempty"`
}
