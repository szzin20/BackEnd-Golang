package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToComplaintRequest(complaint web.ComplaintRequest, TransactionID int) *schema.Complaint {
	return &schema.Complaint{
		Title:   complaint.Title,
		Content: complaint.Content,
	}
}
