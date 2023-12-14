package test

import (
	"bytes"
	"net/http"
	"os"
	"testing"
)

func TesAdminLogin(t *testing.T) {
	AdminEmail := os.Getenv("TEST_ADMIN_EMAIL")
	AdminPassword := os.Getenv("TEST_ADMIN_PASSWORD")

	if AdminEmail == "" || AdminPassword == "" {
		t.Fatal("TEST_ADMIN_EMAIL and/or TEST_ADMIN_PASSWORD environment variables not set")
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
			url := BaseURL + "/admins/login"

			var requestBody string

			switch tt.name {
			case "ValidRequest":
				requestBody = `{"email": "` + AdminEmail + `", "password": "` + AdminPassword + `"}`
			case "MissingEmail":
				requestBody = `{"password": "` + AdminPassword + `"}`
			case "MissingPassword":
				requestBody = `{"email": "` + AdminEmail + `"}`
			case "InvalidPassword":
				requestBody = `{"email": "` + AdminEmail + `", "password": "wrong_password"}`
			case "InvalidEmail":
				requestBody = `{"email": "wrongemail@mail.com", "password": "` + AdminPassword + `"}`
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
