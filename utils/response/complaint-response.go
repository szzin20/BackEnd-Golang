package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToComplaintResponse(complaint *schema.Complaint) web.ComplaintResponse {
	return web.ComplaintResponse{
		Message:   complaint.Message,
		Image:     complaint.Image,
		Audio:     complaint.Audio,
		CreatedAt: complaint.CreatedAt,
	}
}


