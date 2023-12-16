package web

type MedicineRequest struct {
	Code     string `json:"code" form:"code" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Merk     string `json:"merk" form:"merk" validate:"required"`
	Category string `json:"category" form:"category" validate:"required"`
	Type     string `json:"type" form:"type" validate:"required"`
	Stock    int    `json:"stock" form:"stock" validate:"required,min=0"`
	Price    int    `json:"price" form:"price" validate:"required,min=0"`
	Details  string `json:"details" form:"details" validate:"required"`
	Image    string `json:"image" form:"image"`
}

type MedicineUpdateRequest struct {
	Code     string `json:"code" form:"code" validate:"omitempty"`
	Name     string `json:"name" form:"name" validate:"omitempty"`
	Merk     string `json:"merk" form:"merk" validate:"omitempty"`
	Category string `json:"category" form:"category" validate:"omitempty"`
	Type     string `json:"type" form:"type" validate:"omitempty"`
	Stock    int    `json:"stock" form:"stock" validate:"omitempty,min=0"`
	Price    int    `json:"price" form:"price" validate:"omitempty,min=0"`
	Details  string `json:"details" form:"details" validate:"omitempty"`
}

type MedicineImageRequest struct {
	Image string `json:"image" form:"image" validate:"omitempty"`
}
