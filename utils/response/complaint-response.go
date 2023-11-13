package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToComplaintResponse(complaint *schema.Complaint) web.ComplaintResponse {
	return web.ComplaintResponse{
		ID:            complaint.ID,
		TransactionID: complaint.TransactionID,
		Title:         complaint.Title,
		Content:       complaint.Content,
		Status:        complaint.Status,
		CreatedAt:     complaint.CreatedAt,
	}
}

func ConvertToGetAllComplaintsResponse(complaints []schema.Complaint) []web.ComplaintResponse {
	var results []web.ComplaintResponse
	for _, complaint := range complaints {
		complaintResponse := web.ComplaintResponse{
			ID:            complaint.ID,
			TransactionID: complaint.TransactionID,
			Title:         complaint.Title,
			Content:       complaint.Content,
			Status:        complaint.Status,
			CreatedAt:     complaint.CreatedAt,
		}
		results = append(results, complaintResponse)
	}
	return results
}
