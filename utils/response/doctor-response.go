package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToDoctorRegisterResponse(doctor *schema.Doctor) web.DoctorResgisterResponse {
	return web.DoctorResgisterResponse{
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

func ConvertToDoctorLoginResponse(doctor *schema.Doctor) web.DoctorLoginResponse {
	return web.DoctorLoginResponse{
		Email:    doctor.Email,
		Password: doctor.Password,
	}
}

func ConvertToDoctorUpdateResponse(doctor *schema.Doctor) web.DoctorUpdateResponse {
	return web.DoctorUpdateResponse{
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

func ConvertToGetDoctorResponse(doctor *schema.Doctor) web.DoctorResgisterResponse {
	return web.DoctorResgisterResponse{
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