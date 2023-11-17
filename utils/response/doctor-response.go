package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

// func ConvertToDoctorRegisterResponse(doctor *schema.Doctor) web.DoctorResgisterResponse {
// 	return web.DoctorResgisterResponse{
// 		Fullname:           doctor.Fullname,
// 		Email:              doctor.Email,
// 		Price:              doctor.Price,
// 		Tag:                doctor.Tag,
// 		ProfilePicture:     doctor.ProfilePicture,
// 		RegistrationLetter: doctor.RegistrationLetter,
// 	}
// }

func ConvertToDoctorLoginResponse(doctor *schema.Doctor) web.DoctorLoginResponse {
	return web.DoctorLoginResponse{
		Fullname: doctor.Fullname,
		Email:    doctor.Email,
	}
}

func ConvertToDoctorUpdateResponse(doctor *schema.Doctor) web.DoctorUpdateResponse {
	return web.DoctorUpdateResponse{
		Fullname:           doctor.Fullname,
		Email:              doctor.Email,
		Price:              doctor.Price,
		Tag:                doctor.Tag,
		ProfilePicture:     doctor.ProfilePicture,
		RegistrationCertificate: doctor.RegistrationCertificate,
	}
}

func ConvertToGetDoctorResponse(doctor *schema.Doctor) web.DoctorUpdateResponse {
	return web.DoctorUpdateResponse{
		Fullname:                doctor.Fullname,
		Email:                   doctor.Email,
		Price:                   doctor.Price,
		Tag:                     doctor.Tag,
		ProfilePicture:          doctor.ProfilePicture,
		RegistrationCertificate: doctor.RegistrationCertificate,
	}
}

func ConvertToGetAllDoctorResponse(doctors []schema.Doctor) []web.DoctorAllResponse {
	var results []web.DoctorAllResponse

	// Iterasi melalui setiap dokter dan konversi ke format respons
	for _, doctor := range doctors {
		doctorResponse := web.DoctorAllResponse{
			Fullname:           doctor.Fullname,
			Email:              doctor.Email,
			Price:              doctor.Price,
			Tag:                doctor.Tag,
			Status:             doctor.Status,
			ProfilePicture:     doctor.ProfilePicture,
			RegistrationCertificate: doctor.RegistrationCertificate,
		}

		results = append(results, doctorResponse)
	}

	return results
}

func ConvertToDoctorLogoutResponse(doctor *schema.Doctor) web.DoctorLogOutResponse {
	return web.DoctorLogOutResponse{
		Fullname: doctor.Fullname,
		Email:    doctor.Email,
		Tag:      doctor.Tag,
	}
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

