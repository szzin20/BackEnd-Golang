package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToDoctorRegisterRequest(doctor web.DoctorRegisterRequest) *schema.Doctor {
	return &schema.Doctor{
		Fullname:       doctor.Fullname,
		Email:          doctor.Email,
		Price:          doctor.Price,
		Specialist:     doctor.Specialist,
		ProfilePicture: doctor.ProfilePicture,
		Gender:         doctor.Gender,
		NoSTR:          doctor.NoSTR,
		Experience: doctor.Experience,
	}
}

func ConvertToDoctorLoginRequest(doctor web.DoctorLoginRequest) *schema.Doctor {
	return &schema.Doctor{
		Email:    doctor.Email,
		Password: doctor.Password,
	}
}

