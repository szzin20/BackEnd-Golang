package response

// import (
// 	"healthcare/models/schema"
// 	"healthcare/models/web"
// )

// func ConvertToComplaintResponse(UserComplaint schema.User, transaction schema.DoctorTransaction, doctor schema.Doctor) web.ComplaintResponse {
// 	return web.ComplaintResponse{
// 		UserID:              UserComplaint.ID,
// 		UserFullname:        UserComplaint.Fullname,
// 		DoctorID:            doctor.ID,
// 		DoctorFullname:      doctor.Fullname,
// 		Specialist:          doctor.Specialist,
// 		DoctorStatus:        doctor.Status,
// 		DoctorTransactionID: transaction.ID,
// 		CreatedAt:           transaction.CreatedAt,
// 		HealthDetails:       transaction.HealthDetails,
// 		PatientStatus:       transaction.PatientStatus,
// 	}
// }

// func ConvertToComplaintsResponse(UserComplaint schema.User, transaction schema.DoctorTransaction) web.UserComplaintResponse {
// 	return web.UserComplaintResponse{
// 		UserFullname:        UserComplaint.Fullname,
// 		Gender:              UserComplaint.Gender,
// 		Weight:              UserComplaint.Weight,
// 		DoctorTransactionID: transaction.ID,
// 		HealthDetails:       transaction.HealthDetails,
// 		PatientStatus:       transaction.PatientStatus,
// 		CreatedAt:           transaction.CreatedAt,
// 	}
// }


