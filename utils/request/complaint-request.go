package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToComplaintRequest(complaint web.ComplaintRequest, TransactionID uint) *schema.Complaint {
	return &schema.Complaint{
		Message: complaint.Message,
		Image:   complaint.Image,
		Audio:   complaint.Audio,
	}
}
