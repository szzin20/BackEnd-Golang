package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToDoctorRegisterResponse(doctor *schema.Doctor) web.DoctorResgisterResponse {
	return web.DoctorResgisterResponse{
		Fullname:       doctor.Fullname,
		Email:          doctor.Email,
		Price:          doctor.Price,
		Gender:         doctor.Gender,
		Specialist:     doctor.Specialist,
		ProfilePicture: doctor.ProfilePicture,
		NoSTR:          doctor.NoSTR,
		Experience:     doctor.Experience,
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
		Fullname:       doctor.Fullname,
		Email:          doctor.Email,
		Price:          doctor.Price,
		Gender:         doctor.Gender,
		Specialist:     doctor.Specialist,
		ProfilePicture: doctor.ProfilePicture,
		NoSTR:          doctor.NoSTR,
		Experience:     doctor.Experience,
	}
}

func ConvertToGetDoctorResponse(doctor *schema.Doctor) web.DoctorResponse {
	return web.DoctorResponse{
		Fullname:       doctor.Fullname,
		Email:          doctor.Email,
		Password:       doctor.Password,
		Gender:         doctor.Gender,
		Specialist:     doctor.Specialist,
		ProfilePicture: doctor.ProfilePicture,
		NoSTR:          doctor.NoSTR,
		Experience:     doctor.Experience,
	}
}

func ConvertToGetAllDoctorResponse(doctors []schema.Doctor) []web.DoctorAllResponse {
	var results []web.DoctorAllResponse

	// Iterasi melalui setiap dokter dan konversi ke format respons
	for _, doctor := range doctors {
		doctorResponse := web.DoctorAllResponse{
			Fullname:   doctor.Fullname,
			Gender:     doctor.Gender,
			Status:     doctor.Status,
			Price:      doctor.Price,
			Specialist: doctor.Specialist,
			Experience: doctor.Experience,
		}

		results = append(results, doctorResponse)
	}

	return results
}

// patient

// func ConvertToDoctorPatientsResponse(doctor schema.Doctor) web.DoctorPatientsResponse {
// 	response := web.DoctorPatientsResponse{
// 		DoctorID:       doctor.ID,
// 		TransactionID:  0,
// 		DoctorFullname: doctor.Fullname,
// 		Patients:       []web.PatientInfoResponse{},
// 	}

// 	for _, transaction := range doctor.DoctorTransaction {
// 		var user schema.User
// 		// Handle error if user is not found
// 		if err := configs.DB.First(&user, transaction.UserID).Error; err != nil {
// 			continue
// 		}

// 		patientInfo := web.PatientInfoResponse{
// 			ID:        uint(user.ID),
// 			Name:      user.Fullname,
// 			Email:     user.Email,
// 			Height:    user.Height,
// 			Weight:    user.Height,
// 			Gender:    user.Gender,
// 			Birthdate: user.Birthdate,
// 			BloodType: user.BloodType,
// 			Status:    transaction.PaymentStatus,
// 		}

// 		response.Patients = append(response.Patients, patientInfo)
// 	}

// 	return response
// }
