package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateDoctorTransactionResponse(doctorTransaction *schema.DoctorTransaction, doctor schema.Doctor) web.CreateDoctorTransactionResponse {
	return web.CreateDoctorTransactionResponse{
		ID:                  doctorTransaction.ID,
		Fullname:            doctor.Fullname,
		Specialist:          doctor.Specialist,
		Price:               doctor.Price,
		PaymentMethod:       doctorTransaction.PaymentMethod,
		PaymentStatus:       doctorTransaction.PaymentStatus,
		PaymentConfirmation: doctorTransaction.PaymentConfirmation,
		CreatedAt:           doctorTransaction.CreatedAt,
	}
}

func ConvertToGetDoctorTransactionResponse(doctorTransaction schema.DoctorTransaction, doctor schema.Doctor) web.CreateDoctorTransactionResponse {
	return web.CreateDoctorTransactionResponse{
		ID:                  doctorTransaction.ID,
		Fullname:            doctor.Fullname,
		Specialist:          doctor.Specialist,
		Price:               doctor.Price,
		PaymentMethod:       doctorTransaction.PaymentMethod,
		PaymentStatus:       doctorTransaction.PaymentStatus,
		PaymentConfirmation: doctorTransaction.PaymentConfirmation,
		CreatedAt:           doctorTransaction.CreatedAt,
	}
}

func ConvertToGetUserTransactionbyAdminResponse(doctorTransaction schema.DoctorTransaction) web.GetUserTransactionbyAdminResponse {
	return web.GetUserTransactionbyAdminResponse{
		ID:                  doctorTransaction.ID,
		DoctorID:            doctorTransaction.DoctorID,
		PaymentMethod:       doctorTransaction.PaymentMethod,
		Price:               doctorTransaction.Price,
		CreatedAt:           doctorTransaction.CreatedAt,
		PaymentConfirmation: doctorTransaction.PaymentConfirmation,
		PaymentStatus:       doctorTransaction.PaymentStatus,
	}
}

func ListConvertToGetUserTransactionbyAdminResponse(data []schema.DoctorTransaction) []web.GetUserTransactionbyAdminResponse {
    response := []web.GetUserTransactionbyAdminResponse{}
    for _, v := range data {
        responseDetail := ConvertToGetUserTransactionbyAdminResponse(v)
        response = append(response, responseDetail)
    }
    return response
}

func ConvertToGetAllDoctorTransactionsResponse(doctorTransaction schema.DoctorTransaction, doctor schema.Doctor) web.DoctorTransactionsResponse {
	return web.DoctorTransactionsResponse{
		ID:       doctorTransaction.ID,
		Fullname: doctor.Fullname,
	}
}

func ConvertToGetAllDoctorTransactionDetailsResponse(doctorTransaction schema.DoctorTransaction, doctor schema.Doctor) web.DoctorTransactionDetailsResponse {
	return web.DoctorTransactionDetailsResponse{
		ID:            doctorTransaction.ID,
		Fullname:      doctor.Fullname,
		Specialist:    doctor.Specialist,
		PatientStatus: doctorTransaction.PatientStatus,
	}
}

// func ConvertToDoctorTransactionDetailsResponse(doctorTransaction *schema.DoctorTransaction) web.DoctorTransactionDetailsResponse {
// 	doctorTransactions := web.DoctorTransactionDetailsResponse{
// 		ID:       doctorTransaction.ID,
// 		DoctorID: doctorTransaction.DoctorID,
// 		UserID:   doctorTransaction.UserID,
// 	}

// 	var complaintResults []web.ComplaintsResponse
// 	for _, complaint := range doctorTransaction.Complaint {
// 		complaintResponses := web.ComplaintsResponse{
// 			ID:            complaint.ID,
// 			Message:       complaint.Message,
// 			Image:         complaint.Image,
// 			Audio:         complaint.Audio,
// 			CreatedAt:     complaint.CreatedAt,
// 		}
// 		complaintResults = append(complaintResults, complaintResponses)
// 	}
// 	doctorTransactions.Complaint = complaintResults

// 	var adviceResults []web.AdvicesResponse
// 	for _, advice := range doctorTransaction.Advice {
// 		adviceResponses := web.AdvicesResponse{
// 			ID:            advice.ID,
// 			Message:       advice.Message,
// 			Image:         advice.Image,
// 			Audio:         advice.Audio,
// 			CreatedAt:     advice.CreatedAt,
// 		}
// 		adviceResults = append(adviceResults, adviceResponses)
// 	}
// 	doctorTransactions.Advice = adviceResults

// 	return doctorTransactions
// }
