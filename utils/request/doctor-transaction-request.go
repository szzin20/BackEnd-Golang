package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateDoctorTransactionRequest(doctorTransaction web.CreateDoctorTransactionRequest, userID, doctorID uint, fullname string, specialist string, price int) *schema.DoctorTransaction {
	return &schema.DoctorTransaction{
		UserID:              userID,
		DoctorID:            doctorID,
		Price:               price,
		PaymentMethod:       doctorTransaction.PaymentMethod,
		PaymentConfirmation: doctorTransaction.PaymentConfirmation,
	}
}
