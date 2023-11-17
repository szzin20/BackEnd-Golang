package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

// func ConvertToDoctorRegisterRequest(doctor web.DoctorRegisterRequest) *schema.Doctor {
// 	return &schema.Doctor{
// 		Fullname:           doctor.Fullname,
// 		Email:              doctor.Email,
// 		Password:           doctor.Password,
// 		Price:              doctor.Price,
// 		Tag:                doctor.Tag,
// 		ProfilePicture:     doctor.ProfilePicture,
// 		RegistrationLetter: doctor.RegistrationLetter,
// 	}
// }

func ConvertToDoctorLoginRequest(doctor web.DoctorLoginRequest) *schema.Doctor {
	return &schema.Doctor{
		Email:    doctor.Email,
		Password: doctor.Password,
	}
}

