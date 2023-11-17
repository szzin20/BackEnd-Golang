package web

// type DoctorRegisterRequest struct {
// 	Fullname                string `json:"fullname" form:"fullname" validate:"required"`
// 	Email                   string `json:"email" form:"email" validate:"required,email"`
// 	Password                string `json:"password" form:"password" validate:"required,min=8"`
// 	Price                   int    `json:"price" form:"price" validate:"required,min=0"`
// 	Tag                     string `json:"tag" form:"tag" validate:"required"`
// 	ProfilePicture          string `json:"profile_picture" form:"profile_picture" validate:"url"`
// 	RegistrationCertificate int    `json:"registration_certificate" form:"registration_certificate" Validate:"required`
// }

type DoctorLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type DoctorUpdateRequest struct {
	Fullname                string `json:"fullname" form:"fullname" validate:"required"`
	Email                   string `json:"email" form:"email" validate:"required,email"`
	Password                string `json:"password" form:"password" validate:"min=8"`
	Price                   int    `json:"price" form:"price" validate:"min=0"`
	Tag                     string `json:"tag" form:"tag" validate:"required"`
	ProfilePicture          string `json:"profile_picture" form:"profile_picture" validate:"url"`
	RegistrationCertificate int    `json:"registration_certificate" form:"registration_certificate" Validate:"required"`
}
