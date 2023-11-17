package web

import "healthcare/models/schema"

type ComplaintsUpdateRequest struct {
	DoctorTransactionID uint                   `json:"doctorTransactionID" form:"doctorTransactionID"`
	Title               string                 `json:"title" form:"title"`
	Content             string                 `json:"content" form:"content"`
	Status              schema.ComplaintStatus `json:"status" form:"status"`
}
