package web

import "time"

type MedicineResponse struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Merk      string    `json:"merk"`
	Category  string    `json:"category"`
	Type      string    `json:"type"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	Details   string    `json:"details"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

type MedicineUserResponse struct {
	ID       uint   `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Merk     string `json:"merk"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	Details  string `json:"details"`
	Image    string `json:"image"`
}

type MedicineUpdateResponse struct {
	ID        uint      `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Merk      string    `json:"merk"`
	Category  string    `json:"category"`
	Type      string    `json:"type"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
}

type MedicineImageResponse struct {
	Image string `json:"image"`
}
