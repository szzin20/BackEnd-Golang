package web

type ComplaintsUpdateResponse struct {
	DoctorTransactionID uint   `json:"doctor_transactionID" form:"doctor_transactionID"`
	Title               string `json:"title" form:"title"`
	Content             string `json:"content" form:"content"`
	Status              bool   `json:"status" form:"status"`
}
