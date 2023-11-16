package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateDTRequest(DoctorTransaction web.CreateDoctorTransactionRequest) *schema.DoctorTransaction {
	return &schema.DoctorTransaction{
		DoctorID:      DoctorTransaction.DoctorID,
		UserID:        DoctorTransaction.UserID,
		HealthDetails: DoctorTransaction.HealthDetails,
		PaymentMethod: DoctorTransaction.PaymentMethod,
		Price:         DoctorTransaction.Price,
		ImageURL:      DoctorTransaction.ImageURL,
	}
}
