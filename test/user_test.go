package test

import (
	"bytes"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"testing"
)

func TestUserLogin(t *testing.T) {
	godotenv.Load()
	UserEmail := os.Getenv("TEST_USER_EMAIL")
	UserPassword := os.Getenv("TEST_USER_PASSWORD")

	if UserEmail == "" || UserPassword == "" {
		t.Fatal("TEST_USER_EMAIL and/or TEST_USER_PASSWORD environment variables not set")
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
			url := BaseURL + "/users/login"

			var requestBody string

			switch tt.name {
			case "ValidRequest":
				requestBody = `{"email": "` + UserEmail + `", "password": "` + UserPassword + `"}`
			case "MissingEmail":
				requestBody = `{"password": "` + UserPassword + `"}`
			case "MissingPassword":
				requestBody = `{"email": "` + UserEmail + `"}`
			case "InvalidPassword":
				requestBody = `{"email": "` + UserEmail + `", "password": "wrong_password"}`
			case "InvalidEmail":
				requestBody = `{"email": "wrongemail@mail.com", "password": "` + UserPassword + `"}`
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
