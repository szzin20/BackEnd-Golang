package web

type DoctorLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=10,max=15"`
}

type DoctorUpdateRequest struct {
	Fullname       string `json:"fullname" form:"fullname" validate:"max=30"`
	Email          string `json:"email" form:"email" validate:"required,email"`
	Password       string `json:"password" form:"password" validate:"min=10,max=15"`
	Gender         string `json:"gender" form:"gender" validate:"required"`
	Specialist     string `json:"specialist" form:"specialist" validate:"required"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Status         bool   `json:"status" form:"status" validate:"eq=true|eq=false"`
	Experience     string `json:"experience" form:"experience" validate:"required"`
	Alumnus        string `json:"alumnus" form:"alumnus" validate:"required"`
}
