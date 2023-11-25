package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToComplaintResponse(complaint *schema.Complaint) web.CreateComplaintResponse{
	return web.CreateComplaintResponse{
		TransactionID: complaint.TransactionID,
		ID:            complaint.ID,
		Message:       complaint.Message,
		Image:         complaint.Image,
		Audio:         complaint.Audio,
		CreatedAt:     complaint.CreatedAt,
	}
}

func ConvertToDoctorTransactionResponse(doctorTransaction *schema.DoctorTransaction) web.DoctorTransactionResponse {
	doctorTransactions := web.DoctorTransactionResponse{
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

	return doctorTransactions
}
