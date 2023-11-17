package web

import "healthcare/models/schema"

type ComplaintsUpdateResponse struct {
	DoctorTransactionID uint                   `json:"doctor_transactionID" form:"doctor_transactionID"`
	Title               string                 `json:"title" form:"title"`
	Content             string                 `json:"content" form:"content"`
	Status              schema.ComplaintStatus `json:"status" form:"status"`
}
