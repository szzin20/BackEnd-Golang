package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToMedicineTransactionRequest(mt web.MedicineTransactionRequest, userID uint) *schema.MedicineTransaction {

	medicineDetails := make([]schema.MedicineDetails, len(mt.MedicineDetails))

	for i, mdReq := range mt.MedicineDetails {
		medicineDetails[i] = schema.MedicineDetails{
			MedicineID: mdReq.MedicineID,
			Quantity:   mdReq.Quantity,
		}
	}

	return &schema.MedicineTransaction{
		UserID:          userID,
		Name:            mt.Name,
		Address:         mt.Address,
		HP:              mt.HP,
		PaymentMethod:   mt.PaymentMethod,
		MedicineDetails: medicineDetails,
	}
}
