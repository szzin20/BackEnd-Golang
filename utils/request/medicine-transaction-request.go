package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToMedicineTransactionRequest(mt web.MedicineTransactionRequest) *schema.MedicineTransaction {

	medicineDetails := make([]schema.MedicineDetails, len(mt.MedicineDetailsRequest))

	for i, mdReq := range mt.MedicineDetailsRequest {
		medicineDetails[i] = schema.MedicineDetails{
			MedicineID: mdReq.MedicineID,
			Quantity:   mdReq.Quantity,
		}
	}

	return &schema.MedicineTransaction{
		Address:         mt.Address,
		HP:              mt.HP,
		PaymentMethod:   mt.PaymentMethod,
		MedicineDetails: medicineDetails,
	}
}
