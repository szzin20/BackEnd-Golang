package test

import (
	"bytes"
	"net/http"
	"os"
	"testing"
)

var BaseURL = "https://dev.healthify.my.id"

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
		{"InvalidValueSpecializationParams", "?specialist=dermathology&limit=5&offset=0", http.StatusNotFound},
		{"MissingLimit", "?specialist=gigi&offset=0", http.StatusBadRequest},
		{"MissingOffset", "?specialist=gigi&limit=5", http.StatusBadRequest},
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

func TesDoctorLogin(t *testing.T) {
	DoctorEmail := os.Getenv("TEST_DOCTOR_EMAIL")
	DoctorPassword := os.Getenv("TEST_DOCTOR_PASSWORD")

	if DoctorEmail == "" || DoctorPassword == "" {
		t.Fatal("TEST_DOCTOR_EMAIL and/or TEST_DOCTOR_PASSWORD environment variables not set")
	}

	tests := []struct {
		name         string
		expectedCode int
	}{
		{"ValidRequest", http.StatusOK},
		{"MissingEmail", http.StatusBadRequest},
		{"MissingPassword", http.StatusBadRequest},
		{"InvalidPassword", http.StatusUnauthorized},
		{"InvalidEmail", http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := BaseURL + "/doctors/login"

			var requestBody string

			switch tt.name {
			case "ValidRequest":
				requestBody = `{"email": "` + DoctorEmail + `", "password": "` + DoctorPassword + `"}`
			case "MissingEmail":
				requestBody = `{"password": "` + DoctorPassword + `"}`
			case "MissingPassword":
				requestBody = `{"email": "` + DoctorEmail + `"}`
			case "InvalidPassword":
				requestBody = `{"email": "` + DoctorEmail + `", "password": "wrong_password"}`
			case "InvalidEmail":
				requestBody = `{"email": "wrongemail@mail.com", "password": "` + DoctorPassword + `"}`
			}

			req, err := http.NewRequest("POST", url, bytes.NewBufferString(requestBody))
			if err != nil {
				t.Fatalf("Error creating POST request: %s", err)
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error making POST request: %s", err)
			}
			defer response.Body.Close()

			if response.StatusCode != tt.expectedCode {
				t.Errorf("Expected status code %d, got %d", tt.expectedCode, response.StatusCode)
			}
		})
	}
}
