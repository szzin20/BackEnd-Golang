package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToComplaintRequest(complaint web.CreateComplaintRequest, TransactionID uint) *schema.Complaint {
	return &schema.Complaint{
		TransactionID: TransactionID,
		Message:       complaint.Message,
		Image:         complaint.Image,
		Audio:         complaint.Audio,
	}
}

func ConvertToComplaintsRequest(Complaints web.UpdateComplaintRequest) *schema.DoctorTransaction {
	return &schema.DoctorTransaction{
		HealthDetails: Complaints.HealthDetails,
		PatientStatus: Complaints.PatientStatus,
	}
}

