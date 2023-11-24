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
		Password:       doctor.Password,
		Gender:         doctor.Gender,
		NoSTR:          doctor.NoSTR,
		Experience:     doctor.Experience,
		Alumnus:        doctor.Alumnus,
	}
}

func ConvertToDoctorLoginRequest(doctor web.DoctorLoginRequest) *schema.Doctor {
	return &schema.Doctor{
		Email:    doctor.Email,
		Password: doctor.Password,
	}
}

func ConvertToDoctorUpdateRequest(doctor web.DoctorUpdateRequest) *schema.Doctor {
	return &schema.Doctor{
		Fullname:         doctor.Fullname,
		Email:            doctor.Email,
		Specialist:       doctor.Specialist,
		ProfilePicture:   doctor.ProfilePicture,
		AboutDoctor:      doctor.AboutDoctor,
		LocationPractice: doctor.LocationPractice,
		Gender:           doctor.Gender,
		Password:         doctor.Password,
		Status:           doctor.Status,
		Experience:       doctor.Experience,
		Alumnus:          doctor.Alumnus,
	}
}

// func ConvertToDoctorTransactionRequest(updateRequest web.DoctorPatientRequest) *schema.DoctorTransaction {
// 	return &schema.DoctorTransaction{
// 		HealthDetails: updateRequest.HealthDetails,
// 		PatientStatus: updateRequest.PatientStatus,
// 	}
// }
