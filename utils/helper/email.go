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
		return errors.New("incomplete smtp configuration. please set all required environment variables.")
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


func SendNotificationEmail(to, fullname, notificationType, userType, userEmail, userPassword string, includeCredentials bool, roomNumber int) error {
	go func() {
		var subject, body string

		switch notificationType {
		case "login":
			subject = "Healthify Notification"
			body = "Hello, " + fullname + "! You have successfully logged in."
		case "register":
			subject = "Healthify Notification"
			if userType == "doctor" {
				body = fmt.Sprintf("Hallo %s,\n\n<br>Selamat! Akun Anda telah berhasil dibuat di platform kami. Sekarang Anda memiliki akses penuh untuk menjelajahi layanan kami yang memudahkan manajemen pasien dan informasi medis.\n<br><br>Dengan akun ini, Anda dapat dengan mudah mengelola konsultasi pasien, melacak riwayat pasien, mengelola artikel kesehatan, dan mengakses obat-obatan yang tersedia untuk dijadikan rekomendasi obat pada pasien.\n<br><br>Langkah berikutnya, silakan masuk dengan email dan password yang terdaftar dibawah ini :\n<br><br>Email : %s\n<br>Password : %s\n<br><br>Email dan password ini bersifat rahasia, jangan berikan kepada siapapun, agar tidak ada penyalah gunaan akun.\n\n<br><br>Terima kasih atas kepercayaan Anda pada layanan kami. Semoga akun baru ini membantu meningkatkan efisiensi dan kualitas layanan medis Anda.", fullname, userEmail, userPassword)
			} else {
				body = fmt.Sprintf("Hallo %s,\n<br>Kamu berhasil daftar di aplikasi Healthify Care System!\n<br><br>Kami mengarahkan kamu untuk langsung mulai pada halaman beranda, agar kamu dapat memulai perjalanan menuju hidup sehat bersama Healthify.\n<br><br>Dengan mendaftar, Kamu menyetujui Kebijakan Privasi Kesehatan Healthify.", fullname)
			}
		case "complaints":
			subject = "Healthify Notification"
			body = "Hello, " + fullname + "! You have a new consultation request that requires immediate attention. Please review and attend to it promptly."
		default:
			err := errors.New("invalid notification type")
			log.Println(err)
			return
		}

		imageURL := "https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjAfO1adC7X4vJbrrL32Y-50nSyTIRi0X9GZg38xX8Pp7wLQaGhUAActrcIXOflN7mc8Q6vlodQl21TieiybFKuDY1XOrcznX_tDyvwr7vimXxHv80ijlFyTHeiyXmYuYUB77UlBU3PbuvKNsC2FHsdtXH6_W4I-XmtWHThHf4TwMUFjQY2CMbMwxcMK-Fr/s328/Frame%202.png"
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
					.bold {
						font-weight: bold;
					}
					.blue-text {
						color: #007bff;
					}
					.button {
						display: inline-block;
						padding: 10px 20px;
						font-size: 16px;
						text-align: center;
						text-decoration: none;
						background-color: %s; 
						color: #ffffff;
						border-radius: 5px;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1><img src="%s" alt="Healthify Notification"></h1>
					<p>%s</p>
					%s <!-- Button -->
				</div>
			</body>
			</html>
		`, getButtonColor(notificationType), imageURL, body, getButtonHTML(notificationType, roomNumber))

		err := SendEmail(to, subject, body, htmlBody)
		if err != nil {
			log.Printf("Failed to send email: %v\n", err)
		}
	}()

	return nil
}

func getButtonColor(notificationType string) string {
	switch notificationType {
	case "complaints":
		return "#20B2AA"
	default:
		return "#007bff"
	}
}

// Update the getButtonHTML function to include the room number
func getButtonHTML(notificationType string, roomNumber int) string {
	switch notificationType {
	case "complaints":
		link := fmt.Sprintf("https://healthify-doctor.vercel.app/chat/user?status=all&room=%d", roomNumber)
		return fmt.Sprintf(`<a class="button" href="%s" style="background-color: #20B2AA; text-decoration: none; color: #ffffff; padding: 10px 20px; font-size: 16px; border-radius: 5px; transition: background-color 0.3s;">Attend to Complaints</a>`, link)
	default:
		return ""
	}
}


// SendOTPViaEmail sends a one-time password (OTP) via email.
func SendOTPViaEmail(email string) error {
	// Generate OTP
	otp, err := GenerateRandomCode()
	if err != nil {
		log.Printf("Failed to generate OTP: %v\n", err)
		return err
	}

	// Save OTP to the database
	if err := SaveOTP(email, otp); err != nil {
		log.Printf("Failed to save OTP to the database: %v\n", err)
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
