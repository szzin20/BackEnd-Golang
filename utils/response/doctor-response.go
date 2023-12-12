package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToDoctorRegisterResponse(doctor *schema.Doctor) web.DoctorRegisterResponse {
	return web.DoctorRegisterResponse{
		ID:             doctor.ID,
		Fullname:       doctor.Fullname,
		Email:          doctor.Email,
		Price:          doctor.Price,
		Gender:         doctor.Gender,
		Specialist:     doctor.Specialist,
		ProfilePicture: doctor.ProfilePicture,
		NoSTR:          doctor.NoSTR,
		Experience:     doctor.Experience,
		Alumnus:        doctor.Alumnus,
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
		ProfilePicture: doctor.ProfilePicture,
		Fullname:       doctor.Fullname,
		Gender:         doctor.Gender,
		Email:          doctor.Email,
		Price:          doctor.Price,
		Specialist:     doctor.Specialist,
		Experience:     doctor.Experience,
		Alumnus:        doctor.Alumnus,
		NoSTR:          doctor.NoSTR,
	}
}

func ConvertToGetAllDoctorResponse(doctors []schema.Doctor) []web.DoctorAllResponse {
	var results []web.DoctorAllResponse

	for _, doctor := range doctors {
		doctorResponse := web.DoctorAllResponse{
			ID:             doctor.ID,
			ProfilePicture: doctor.ProfilePicture,
			Fullname:       doctor.Fullname,
			Price:          doctor.Price,
			Specialist:     doctor.Specialist,
			Status:         doctor.Status,
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
			Price:          doctor.Price,
			Specialist:     doctor.Specialist,
			Experience:     doctor.Experience,
			Alumnus:        doctor.Alumnus,
			NoSTR:          doctor.NoSTR,
		}

		results = append(results, doctorResponse)
	}

	return results
}

func ConvertToGetIDDoctorResponse(doctor *schema.Doctor) web.DoctorIDResponse {
	return web.DoctorIDResponse{
		ID:             doctor.ID,
		ProfilePicture: doctor.ProfilePicture,
		Fullname:       doctor.Fullname,
		Status:         doctor.Status,
		Specialist:     doctor.Specialist,
		Price:          doctor.Price,
		Experience:     doctor.Experience,
		NoSTR:          doctor.NoSTR,
		Alumnus:        doctor.Alumnus,
	}
}
func ConvertToGetDoctorbyAdminResponse(doctor *schema.Doctor) web.DoctorIDResponseByAdmin {
	return web.DoctorIDResponseByAdmin{
		ID:             doctor.ID,
		ProfilePicture: doctor.ProfilePicture,
		Fullname:       doctor.Fullname,
		Gender:         doctor.Gender,
		Email:          doctor.Email,
		Price:          doctor.Price,
		Specialist:     doctor.Specialist,
		Alumnus:        doctor.Alumnus,
		Experience:     doctor.Experience,
		NoSTR:          doctor.NoSTR,
	}
}

func ConvertToManageUserResponse(managePatient schema.DoctorTransaction, user schema.User) web.ManageUserResponse {
	return web.ManageUserResponse{
		UserID:              user.ID,
		Fullname:            user.Fullname,
		ProfilePicture:      user.ProfilePicture,
		DoctorTransactionID: managePatient.ID,
		CreatedAt:           managePatient.CreatedAt,
		HealthDetails:       managePatient.HealthDetails,
		PatientStatus:       managePatient.PatientStatus,
	}
}

func ConvertToConsultationResponse(consultation schema.DoctorTransaction, user schema.User, room schema.Roomchat) web.DoctorConsultationResponse {
	return web.DoctorConsultationResponse{
		RoomchatID:          room.ID,
		UserID:              user.ID,
		Fullname:            user.Fullname,
		ProfilePicture:      user.ProfilePicture,
		DoctorTransactionID: consultation.ID,
		Price:               consultation.Price,
		CreatedAt:           consultation.CreatedAt,
	}
}

func ConvertToDoctorStatusResponse(doctor *schema.Doctor) web.DoctorStatusResponse {
	return web.DoctorStatusResponse{
		Status: doctor.Status,
	}
}
