package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToAdminMedicineResponse(medicine *schema.Medicine) web.MedicineResponse {
	return web.MedicineResponse{
		Code:     medicine.Code,
		Name:     medicine.Name,
		Merk:     medicine.Merk,
		Category: medicine.Category,
		Type:     medicine.Type,
		Stock:    medicine.Stock,
		Price:    medicine.Price,
		Details:  medicine.Details,
		Image:    medicine.Image,
	}
}

func ConvertToPatientMedicineResponse(medicine *schema.Medicine) web.MedicineResponse {
	return web.MedicineResponse{
		Name:     medicine.Name,
		Merk:     medicine.Merk,
		Category: medicine.Category,
		Type:     medicine.Type,
		Stock:    medicine.Stock,
		Price:    medicine.Price,
		Details:  medicine.Details,
		Image:    medicine.Image,
	}
}
