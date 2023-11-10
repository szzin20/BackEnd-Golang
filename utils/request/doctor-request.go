package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToDoctorRegisterRequest(doctor web.DoctorRegisterRequest) *schema.Doctor {
	return &schema.Doctor{
		Name:               doctor.Name,
		Email:              doctor.Email,
		Password:           doctor.Password,
		Status:             doctor.Status,
		Price:              doctor.Price,
		Tag:                doctor.Tag,
		ImageURL:           doctor.ImageURL,
		RegistrationLetter: doctor.RegistrationLetter,
	}
}

func ConvertToDoctorLoginRequest(doctor web.DoctorLoginRequest) *schema.Doctor {
	return &schema.Doctor{
		Email:    doctor.Email,
		Password: doctor.Password,
	}
}