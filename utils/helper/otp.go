package helper

import (
	"errors"
	"fmt"
	"healthcare/configs"
	"healthcare/models/schema"

	"gorm.io/gorm"
)

// VerifyOTPByEmail 
func VerifyOTPByEmail(email, providedOTP string) error {
	storedOTP, err := getStoredOTP(email)
	if err != nil {
		return err
	}

	if storedOTP != providedOTP {
		return errors.New("invalid OTP: provided OTP does not match stored OTP")
	}

	return nil
}

// SaveOTP 
func SaveOTP(email, otp string) error {
	var user schema.User
	var doctor schema.Doctor

	userError := configs.DB.Where("email = ?", email).First(&user).Error
	doctorError := configs.DB.Where("email = ?", email).First(&doctor).Error

	if userError == nil || doctorError == nil {
		var model interface{}
		if userError == nil {
			model = &user
		} else {
			model = &doctor
		}

		// Update the OTP
		if err := configs.DB.Model(model).Update("OTP", otp).Error; err != nil {
			return err
		}

		return nil
	}

	return errors.New("model not found for the given email")
}

func getStoredOTP(email string) (string, error) {
	var user schema.User
	var doctor schema.Doctor

	if configs.DB.Where("email = ?", email).First(&user).Error == nil {
		return user.OTP, nil
	} else if configs.DB.Where("email = ?", email).First(&doctor).Error == nil {
		return doctor.OTP, nil
	}

	return "", errors.New("model not found for the given email")
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
