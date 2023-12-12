package test

import (
	"net/http"
	"testing"
)

var BaseURL = "https://www.healthify.my.id"

func TestGetAvailableDoctor(t *testing.T) {
	tests := []struct {
		name       string
		queryParam string
		expected   int
	}{
		{"ValidParams", "?limit=5&offset=0", http.StatusOK},
		{"MissingLimit", "?offset=0", http.StatusBadRequest},
		{"MissingOffset", "?limit=5", http.StatusBadRequest},
		{"MissingBoth", "", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := BaseURL + "/users/doctors/available" + tt.queryParam

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

func TestGetSpecializeDoctor(t *testing.T) {
	tests := []struct {
		name        string
		queryParams string
		expected    int
	}{
		{"MissingSpecialization", "", http.StatusBadRequest},
		{"ValidParams", "?specialist=gigi&limit=5&offset=0", http.StatusOK},
		{"MissingLimit", "?specialist=pediatrics&offset=0", http.StatusBadRequest},
		{"MissingOffset", "?specialist=dermatology&limit=5", http.StatusBadRequest},
		{"MissingBoth", "?specialist=gigi", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := BaseURL + "/users/doctors" + tt.queryParams

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
