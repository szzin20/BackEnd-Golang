package test

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strings"
	"testing"
)

func addAdminToken(req *http.Request) {
	godotenv.Load()
	var AdminToken = os.Getenv("ADMIN_TOKEN")
	req.Header.Add("Authorization", "Bearer "+AdminToken)
}

func addUserToken(req *http.Request) {
	godotenv.Load()
	var UserToken = os.Getenv("USER_TOKEN")
	req.Header.Add("Authorization", "Bearer "+UserToken)
}

func TestGetMedicineByUser(t *testing.T) {
	tests := []struct {
		name       string
		queryParam string
		expected   int
	}{
		{"ValidParams", "?limit=5&offset=0", http.StatusOK},
		{"ValidKeyword", "?keyword=bodrex&limit=5&offset=0", http.StatusOK},
		{"ValidID", "/2", http.StatusOK},
		{"MissingLimit", "?offset=0", http.StatusBadRequest},
		{"MissingOffset", "?limit=5", http.StatusBadRequest},
		{"MissingBoth", "", http.StatusBadRequest},
		{"InvalidID", "/invalid", http.StatusBadRequest},
		{"NotFoundID", "/0", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := BaseURL + "/users/medicines" + tt.queryParam

			if strings.Contains(tt.name, "ValidID") || strings.Contains(tt.name, "InvalidID") {
				url = BaseURL + "/users/medicines" + tt.queryParam
			}

			response, err := http.Get(url)
			if err != nil {
				t.Fatalf("Error making GET request: %s", err)
			}
			defer response.Body.Close()

			if response.StatusCode != tt.expected {
				t.Errorf("Expected status code %d, got %d", tt.expected, response.StatusCode)
			}
		})
	}
}

func TestGetMedicineByAdmin(t *testing.T) {
	tests := []struct {
		name       string
		queryParam string
		tokenFunc  func(req *http.Request)
		expected   int
	}{
		{"ValidAdminToken", "?limit=5&offset=0", addAdminToken, http.StatusOK},
		{"ValidKeyword", "?keyword=bodrex&limit=5&offset=0", addAdminToken, http.StatusOK},
		{"ValidID", "/2", addAdminToken, http.StatusOK},
		{"MissingLimit", "?offset=0", addAdminToken, http.StatusBadRequest},
		{"MissingOffset", "?limit=5", addAdminToken, http.StatusBadRequest},
		{"MissingBoth", "", addAdminToken, http.StatusBadRequest},
		{"InvalidID", "/invalid", addAdminToken, http.StatusBadRequest},
		{"NotFoundID", "/0", addAdminToken, http.StatusNotFound},
		{"MissingToken", "?limit=5&offset=0", nil, http.StatusUnauthorized},
		{"InvalidToken", "?limit=5&offset=0", addUserToken, http.StatusForbidden},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := BaseURL + "/admins/medicines" + tt.queryParam

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Fatalf("Error creating GET request: %s", err)
			}

			if tt.tokenFunc != nil {
				tt.tokenFunc(req)
			}

			response, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Error making GET request: %s", err)
			}
			defer response.Body.Close()

			if response.StatusCode != tt.expected {
				t.Errorf("Expected status code %d, got %d", tt.expected, response.StatusCode)
			}
		})
	}
}

func TestCreateMedicineByAdmin(t *testing.T) {
	tests := []struct {
		name      string
		payload   string
		tokenFunc func(req *http.Request)
		expected  int
	}{
		//{"ValidMedicineByAdmin", `{
		//	"code": "PX1509",
		//	"name": "Ibuprofen",
		//	"merk": "Arbupon",
		//	"category": "Obat antiinflamasi nonsteroid",
		//	"type": "12345",
		//	"price": 12000,
		//	"stock": 100,
		//	"details": "12345",
		//}`, addAdminToken, http.StatusCreated},
		{"FieldsRequired", `{"invalid_field": "value"}`, addAdminToken, http.StatusBadRequest},
		{"MissingToken", `{"invalid_field": "value"}`, nil, http.StatusUnauthorized},
		{"InvalidToken", `{"invalid_field": "value"}`, addUserToken, http.StatusForbidden},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := BaseURL + "/admins/medicines"

			req, err := http.NewRequest("POST", url, strings.NewReader(tt.payload))
			if err != nil {
				t.Fatalf("Error creating POST request: %s", err)
			}

			req.Header.Set("Content-Type", "application/json")

			if tt.tokenFunc != nil {
				tt.tokenFunc(req)
			}

			response, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Error making POST request: %s", err)
			}
			defer response.Body.Close()

			if response.StatusCode != tt.expected {
				t.Errorf("Expected status code %d, got %d", tt.expected, response.StatusCode)
			}
		})
	}
}

func TestPutMedicineByIDByAdmin(t *testing.T) {
	tests := []struct {
		name       string
		medicineID int
		payload    string
		tokenFunc  func(req *http.Request)
		expected   int
	}{
		{"ValidUpdateByAdmin", 54, `{
			"price": 15000,
			"stock": 120
		}`, addAdminToken, http.StatusOK},
		{"InvalidUpdateByAdmin", 999, `{"invalid_field": "value"}`, addAdminToken, http.StatusNotFound},
		{"MissingToken", 1, `{"invalid_field": "value"}`, nil, http.StatusUnauthorized},
		{"InvalidToken", 1, `{"invalid_field": "value"}`, addUserToken, http.StatusForbidden},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := BaseURL + fmt.Sprintf("/admins/medicines/%d", tt.medicineID)

			req, err := http.NewRequest("PUT", url, strings.NewReader(tt.payload))
			if err != nil {
				t.Fatalf("Error creating PUT request: %s", err)
			}

			req.Header.Set("Content-Type", "application/json")

			if tt.tokenFunc != nil {
				tt.tokenFunc(req)
			}

			response, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Error making PUT request: %s", err)
			}
			defer response.Body.Close()

			if response.StatusCode != tt.expected {
				t.Errorf("Expected status code %d, got %d", tt.expected, response.StatusCode)
			}
		})
	}
}

func TestDeleteMedicineByIDByAdmin(t *testing.T) {
	tests := []struct {
		name       string
		medicineID int
		tokenFunc  func(req *http.Request)
		expected   int
	}{
		{"ValidDeleteByAdmin", 61, addAdminToken, http.StatusOK},
		{"InvalidDeleteByAdmin", 999, addAdminToken, http.StatusNotFound},
		{"MissingToken", 1, nil, http.StatusUnauthorized},
		{"InvalidToken", 1, addUserToken, http.StatusForbidden},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := BaseURL + fmt.Sprintf("/admins/medicines/%d", tt.medicineID)

			req, err := http.NewRequest("DELETE", url, nil)
			if err != nil {
				t.Fatalf("Error creating DELETE request: %s", err)
			}

			if tt.tokenFunc != nil {
				tt.tokenFunc(req)
			}

			response, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("Error making DELETE request: %s", err)
			}
			defer response.Body.Close()

			if response.StatusCode != tt.expected {
				t.Errorf("Expected status code %d, got %d", tt.expected, response.StatusCode)
			}
		})
	}
}
