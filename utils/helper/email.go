package helper

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// SendEmail sends an email using SMTP server configuration from environment variables.
func SendEmail(to, subject, body, htmlBody string) error {
	// SMTP configuration
	smtpServer := os.Getenv("SMTPSERVER")
	smtpPortStr := os.Getenv("SMTPPORT")
	smtpUsername := os.Getenv("SMTPUSERNAME")
	smtpPassword := os.Getenv("SMTPPASSWORD")

	// Check if all environment variables are set
	if smtpServer == "" || smtpPortStr == "" || smtpUsername == "" || smtpPassword == "" {
		return errors.New("Incomplete SMTP configuration. Please set all required environment variables.")
	}

	// Convert smtpPortStr to int
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return err
	}

	// Configure email dialer
	d := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUsername)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Add HTML body if available
	if htmlBody != "" {
		m.AddAlternative("text/html", htmlBody)
	}

	// Send email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email successfully sent to %s\n", to)

	return nil
}

// SendNotificationEmail sends a notification email based on the notification type.
func SendNotificationEmail(to, fullname, notificationType, userType string) error {
	go func() {
		var subject, body string

		switch notificationType {
		case "login":
			subject = "Login Notification"
			body = "Hello, " + fullname + "! You have successfully logged in."
		case "register":
			subject = "Registration Notification"
			body = "Hello, " + fullname + "! You have successfully registered."
		case "complaint":
			subject = "Consultation Notification"
			body = "Hello, " + fullname + "! You have a new consultation request that requires immediate attention. Please review and attend to it promptly."
		default:
			err := errors.New("Invalid notification type")
			log.Println(err)
			return
		}

		htmlBody := fmt.Sprintf(`
			<!DOCTYPE html>
			<html>
			<head>
				<style>
					/* CSS */
					body {
						font-family: Arial, sans-serif;
						background-color: #f4f4f4;
						padding: 20px;
					}
					.container {
						background-color: #ffffff;
						padding: 20px;
						border-radius: 10px;
					}
					h1 {
						color: #007bff;
					}
					p {
						color: #333;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>%s</h1>
					<p>%s</p>
				</div>
			</body>
			</html>
		`, subject, body)

		err := SendEmail(to, subject, body, htmlBody)
		if err != nil {
			log.Printf("Failed to send email: %v\n", err)
		}
	}()

	return nil
}

// SendOTPViaEmail sends a one-time password (OTP) via email.
func SendOTPViaEmail(email string) error {
	// Generate OTP
	otp, err := GenerateRandomCode()
	if err != nil {
		log.Printf("Failed to generate OTP: %v\n", err)
		return err
	}

	go func() {
		// Email body
		subject := "Your One-Time Password"
		body := fmt.Sprintf("Your OTP is: %s", otp)

		htmlBody := fmt.Sprintf(`
			<!DOCTYPE html>
			<html>
			<head>
				<style>
					/* CSS styling for email body */
					body {
						font-family: Arial, sans-serif;
						background-color: #f4f4f4;
						padding: 20px;
					}
					.container {
						background-color: #ffffff;
						padding: 20px;
						border-radius: 10px;
					}
					h1 {
						color: #007bff;
					}
					p {
						color: #333;
						font-size: 16px;
					}
					/* Styling for OTP code */
					.otp-code {
						color: #ff0000; /* Red color */
						font-size: 20px; /* Larger font size */
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>%s</h1>
					<p>Your OTP is: <strong class="otp-code">%s</strong></p>
				</div>
			</body>
			</html>
		`, subject, otp)

		// Send email
		err := SendEmail(email, subject, body, htmlBody)
		if err != nil {
			log.Printf("Failed to send OTP email: %v\n", err)
		}
	}()

	return nil
}
