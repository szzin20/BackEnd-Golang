package test

import (
	"net/http"
	"strings"
	"testing"
)

var AdminToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluc3VwZXJyQG1haWwuY29tIiwiZXhwIjoxNzAyNzk1MjU1LCJpZCI6MSwicm9sZSI6ImFkbWluIn0.VFt7z567Mtkl9ID_Z_B6GTIWRR1BQa_aLHFSmaDEYjY"
var UserToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InBhdGllbnQydGVzdEBnbWFpbC5jb20iLCJleHAiOjE3MDI3OTUxNjksImlkIjoxMzgsInJvbGUiOiJ1c2VyIn0.e5sUfwf2OxlgV-5ClgCum-CIpYw8wouGI-c3tSHeDuM"

func addAdminToken(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+AdminToken)
}

func addUserToken(req *http.Request) {
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

func TestGetAllMedicineByAdminToken(t *testing.T) {
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
