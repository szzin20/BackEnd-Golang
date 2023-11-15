package web

type ComplaintsUpdateRequest struct {
	DoctorTransactionID uint    `json:"doctorTransactionID" form:"doctorTransactionID"`
	Title               string `json:"title" form:"title"`
	Content             string `json:"content" form:"content"`
	Status              bool   `json:"status" form:"status"`
}
