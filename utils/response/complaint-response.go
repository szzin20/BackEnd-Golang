package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToGetAllComplaintsResponse(complaints []schema.Complaint) []web.ComplaintsUpdateResponse {
    var results []web.ComplaintsUpdateResponse
    // Iterasi melalui setiap keluhan dan konversi ke format respons
    for _, complaint := range complaints {
        complaintsResponse := web.ComplaintsUpdateResponse{
            DoctorTransactionID: uint(complaint.TransactionID),
            Title:               complaint.Title,
            Content:             complaint.Content,
            Status:              complaint.Status,
        }

        results = append(results, complaintsResponse)
    }

    return results
}


func ConvertToGetComplaintsResponse(complaints *schema.Complaint) web.ComplaintsUpdateResponse {
	return web.ComplaintsUpdateResponse{
		DoctorTransactionID: uint(complaints.TransactionID),
		Title:               complaints.Title,
		Content:             complaints.Content,
		Status:              complaints.Status,
	}
}

func ConvertToUpdateComplaintsResponse(complaints *schema.Complaint) web.ComplaintsUpdateResponse {
	return web.ComplaintsUpdateResponse{
		DoctorTransactionID: uint(complaints.TransactionID),
		Title:               complaints.Title,
		Content:             complaints.Content,
		Status:              complaints.Status,
	}
}