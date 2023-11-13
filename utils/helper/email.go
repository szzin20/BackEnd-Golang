package helper

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// sendEmail sends an email
func sendEmail(to, subject, body string) error {
	// Set up email server configuration
	smtpServer := os.Getenv("SMTPSERVER")
	smtpPortStr := os.Getenv("SMTPPORT")
	smtpUsername := os.Getenv("SMTPUSERNAME")
	smtpPassword := os.Getenv("SMTPPASSWORD")

	// Convert smtpPortStr to int
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return err
	}

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", smtpUsername)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	// Set up the email server dialer
	d := gomail.NewDialer(smtpServer, smtpPort, smtpUsername, smtpPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// sendVerificationEmail sends a verification email
func SendVerificationEmail(to, verificationCode string) error {
	subject := "Email Verification"
	body := "Your verification code is: " + verificationCode

	return sendEmail(to, subject, body)
}

// sendLoginNotificationEmail sends a login notification email
func SendLoginNotificationEmail(to, username, body string ) error {
	subject := "Login Notification"
	body = "Hello, " + username + "! You have successfully logged in."

	return sendEmail(to, subject, body)
}