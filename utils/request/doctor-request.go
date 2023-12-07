package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToDoctorRegisterRequest(doctor web.DoctorRegisterRequest) *schema.Doctor {
	return &schema.Doctor{
		Fullname:       doctor.Fullname,
		Email:          doctor.Email,
		Password:       doctor.Password,
		Price:          doctor.Price,
		Specialist:     doctor.Specialist,
		ProfilePicture: doctor.ProfilePicture,
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
		ProfilePicture: doctor.ProfilePicture,
		Fullname:       doctor.Fullname,
		Email:          doctor.Email,
		Password:       doctor.Password,
		Price:          doctor.Price,
		Specialist:     doctor.Specialist,
		Gender:         doctor.Gender,
		Experience:     doctor.Experience,
		Alumnus:        doctor.Alumnus,
		NoSTR:          doctor.NoSTR,
	}
}

func ConvertToManageUserUpdateRequest(managePatient web.UpdateManageUserRequest) *schema.DoctorTransaction {
	return &schema.DoctorTransaction{
		HealthDetails: managePatient.HealthDetails,
		PatientStatus: managePatient.PatientStatus,
	}
}

func ConvertToUpdateStatusRequest(doctors web.ChangeDoctorStatusRequest) *schema.Doctor {
	return &schema.Doctor{
		Status: doctors.Status,
	}
}
