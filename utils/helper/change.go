package helper

import (
	"fmt"

	"gorm.io/gorm"
)

// UpdatePasswordInDatabase
func UpdatePasswordInDatabase(db *gorm.DB, tableName, email, hashedPassword, otp string) error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	// Update the password based on email and OTP
	err := db.Table(tableName).Where("email = ? AND otp = ?", email, otp).Update("password", hashedPassword).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteOTPFromDatabase
func DeleteOTPFromDatabase(db *gorm.DB, tableName, email string) error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}

	// Delete the OTP entry based on email
	err := db.Table(tableName).Where("email = ?", email).Update("OTP", "").Error
	if err != nil {
		return err
	}

	return nil
}
