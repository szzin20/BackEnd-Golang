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
	}
}

func ConvertToGetAllDoctorTransactionsResponse(doctorTransaction schema.DoctorTransaction, doctor schema.Doctor) web.DoctorTransactionsResponse {
	return web.DoctorTransactionsResponse{
		ID:       doctorTransaction.ID,
		Fullname: doctor.Fullname,
	}
}

func ConvertToDoctorTransactionDetailsResponse(doctorTransaction *schema.DoctorTransaction) web.DoctorTransactionDetailsResponse {
	doctorTransactions := web.DoctorTransactionDetailsResponse{
		ID:       doctorTransaction.ID,
		DoctorID: doctorTransaction.DoctorID,
		UserID:   doctorTransaction.UserID,
	}

	var complaintResponses []web.ComplaintsResponse
	for _, complaint := range doctorTransaction.Complaint {
		complaints := web.ComplaintsResponse{
			ID:            complaint.ID,
			Message:       complaint.Message,
			Image:         complaint.Image,
			Audio:         complaint.Audio,
			CreatedAt:     complaint.CreatedAt,
		}
		complaintResponses = append(complaintResponses, complaints)
	} 
	doctorTransactions.Complaint = complaintResponses

	var adviceResponses []web.AdvicesResponse
	for _, advice := range doctorTransaction.Advice {
		advices := web.AdvicesResponse{
			ID:            advice.ID,
			Message:       advice.Message,
			Image:         advice.Image,
			Audio:         advice.Audio,
			CreatedAt:     advice.CreatedAt,
		}
		adviceResponses = append(adviceResponses, advices)
	}
	doctorTransactions.Advice = adviceResponses

	return doctorTransactions
}
