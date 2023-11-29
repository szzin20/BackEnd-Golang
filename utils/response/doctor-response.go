package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToDoctorRegisterResponse(doctor *schema.Doctor) web.DoctorRegisterResponse {
	return web.DoctorRegisterResponse{
		Fullname:       doctor.Fullname,
		Email:          doctor.Email,
		Price:          doctor.Price,
		Gender:         doctor.Gender,
		Specialist:     doctor.Specialist,
		ProfilePicture: doctor.ProfilePicture,
		NoSTR:          doctor.NoSTR,
		Experience:     doctor.Experience,
		Alumnus:        doctor.Alumnus,
		Status:         doctor.Status,
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
		Fullname:         doctor.Fullname,
		Email:            doctor.Email,
		Gender:           doctor.Gender,
		Specialist:       doctor.Specialist,
		ProfilePicture:   doctor.ProfilePicture,
		NoSTR:            doctor.NoSTR,
		Status:           doctor.Status,
		Experience:       doctor.Experience,
		Alumnus:          doctor.Alumnus,
	}
}

func ConvertToGetDoctorResponse(doctor *schema.Doctor) web.DoctorProfile {
	return web.DoctorProfile{
		Fullname:       doctor.Fullname,
		Email:          doctor.Email,
		Gender:         doctor.Gender,
		Specialist:     doctor.Specialist,
		ProfilePicture: doctor.ProfilePicture,
		NoSTR:          doctor.NoSTR,
		Experience:     doctor.Experience,
		Alumnus:        doctor.Alumnus,
	}
}

func ConvertToGetAllDoctorResponse(doctors []schema.Doctor) []web.DoctorAllResponse {
	var results []web.DoctorAllResponse

	for _, doctor := range doctors {
		doctorResponse := web.DoctorAllResponse{
			ID:             doctor.ID,
			ProfilePicture: doctor.ProfilePicture,
			Fullname:       doctor.Fullname,
			Status:         doctor.Status,
			Price:          doctor.Price,
			Specialist:     doctor.Specialist,
		}

		results = append(results, doctorResponse)
	}

	return results
}

func ConvertToGetAllDoctorByAdminResponse(doctors []schema.Doctor) []web.DoctorAllResponseByAdmin {
	var results []web.DoctorAllResponseByAdmin

	// Iterasi melalui setiap dokter dan konversi ke format respons
	for _, doctor := range doctors {
		doctorResponse := web.DoctorAllResponseByAdmin{

			ID:             doctor.ID,
			ProfilePicture: doctor.ProfilePicture,
			Fullname:       doctor.Fullname,
			Gender:         doctor.Gender,
			Email:          doctor.Email,
			Status:         doctor.Status,
			Price:          doctor.Price,
			Specialist:     doctor.Specialist,
			NoSTR:          doctor.NoSTR,
			Role:           doctor.Role,
			Experience:     doctor.Experience,
			Alumnus:        doctor.Alumnus,
		}

		results = append(results, doctorResponse)
	}

	return results
}

func ConvertToGetIDDoctorResponse(doctor *schema.Doctor) web.DoctorIDResponse {
	return web.DoctorIDResponse{
		ID:               doctor.ID,
		ProfilePicture:   doctor.ProfilePicture,
		Fullname:         doctor.Fullname,
		Status:           doctor.Status,
		Specialist:       doctor.Specialist,
		Price:            doctor.Price,
		Experience:       doctor.Experience,
		NoSTR:            doctor.NoSTR,		Alumnus:          doctor.Alumnus,
	}
}

func ConvertToManageUserResponse(managePatient schema.DoctorTransaction, user schema.User) web.ManageUserResponse {
	return web.ManageUserResponse{
		UserID:              user.ID,
		Fullname:            user.Fullname,
		DoctorTransactionID: managePatient.ID,
		CreatedAt:           managePatient.CreatedAt,
		HealthDetails:       managePatient.HealthDetails,
		PatientStatus:       managePatient.PatientStatus,
	}
}

func ConvertToConsultationResponse(Consultation schema.DoctorTransaction, user schema.User, doctor schema.Doctor) web.DoctorConsultationResponse {
	return web.DoctorConsultationResponse{
		UserID:              Consultation.ID,
		Fullname:            user.Fullname,
		DoctorTransactionID: Consultation.ID,
		Price:               doctor.Price,
	}
}
