package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToDoctorRegisterResponse(doctor *schema.Doctor) web.DoctorResgisterResponse {
	return web.DoctorResgisterResponse{
		Fullname:               doctor.Fullname,
		Email:              doctor.Email,
		Password:           doctor.Password,
		Price:              doctor.Price,
		Tag:                doctor.Tag,
		ImageURL:           doctor.ImageURL,
		RegistrationLetter: doctor.RegistrationLetter,
	}
}

func ConvertToDoctorLoginResponse(doctor *schema.Doctor) web.DoctorLoginResponse {
	return web.DoctorLoginResponse{
		Fullname: doctor.Fullname,
		Email:    doctor.Email,
	}
}

func ConvertToDoctorUpdateResponse(doctor *schema.Doctor) web.DoctorUpdateResponse {
	return web.DoctorUpdateResponse{
		Fullname:               doctor.Fullname,
		Email:              doctor.Email,
		Password:           doctor.Password,
		Price:              doctor.Price,
		Tag:                doctor.Tag,
		ImageURL:           doctor.ImageURL,
		RegistrationLetter: doctor.RegistrationLetter,
	}
}

func ConvertToGetDoctorResponse(doctor *schema.Doctor) web.DoctorResgisterResponse {
	return web.DoctorResgisterResponse{
		Fullname:               doctor.Fullname,
		Email:              doctor.Email,
		Password:           doctor.Password,
		Price:              doctor.Price,
		Tag:                doctor.Tag,
		ImageURL:           doctor.ImageURL,
		RegistrationLetter: doctor.RegistrationLetter,
	}
}

