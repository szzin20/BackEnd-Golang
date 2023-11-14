package web

type MedicineRequest struct {
	Code     string `json:"code" form:"code"`
	Name     string `json:"name" form:"name"`
	Merk     string `json:"merk" form:"merk"`
	Category string `json:"category" form:"category"`
	Type     string `json:"type" form:"type"`
	Stock    int    `json:"stock" form:"stock"`
	Price    int    `json:"price" form:"price"`
	Details  string `json:"details" form:"details"`
	Image    string `json:"image" form:"image"`
}
