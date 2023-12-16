package helper

import (
	"errors"
	"fmt"
	"healthcare/configs"
	"healthcare/models/schema"

	"gorm.io/gorm"
)

// VerifyOTPByEmail 
func VerifyOTPByEmail(email, providedOTP, userType string) error {
	storedOTP, err := getStoredOTP(email, userType)
	if err != nil {
		return err
	}

	if storedOTP != providedOTP {
		return errors.New("invalid OTP: provided OTP does not match stored OTP")
	}

	return nil
}


// SaveOTP 
func SaveOTP(email, otp string, userType string) error {
	var user schema.User
	var doctor schema.Doctor
	var admin schema.Admin

	var model interface{}

	switch userType {
	case "user":
		if err := configs.DB.Where("email = ?", email).First(&user).Error; err != nil {
			return errors.New("user not found for the given email")
		}
		model = &user
	case "doctor":
		if err := configs.DB.Where("email = ?", email).First(&doctor).Error; err != nil {
			return errors.New("doctor not found for the given email")
		}
		model = &doctor
	case "admin":
		if err := configs.DB.Where("email = ?", email).First(&admin).Error; err != nil {
			return errors.New("admin not found for the given email")
		}
		model = &admin
	default:
		return errors.New("invalid user type")
	}

	// Update the OTP
	if err := configs.DB.Model(model).Update("OTP", otp).Error; err != nil {
		return err
	}

	return nil
}


func getStoredOTP(email string, userType string) (string, error) {
	var user schema.User
	var doctor schema.Doctor
	var admin schema.Admin

	switch userType {
	case "user":
		if err := configs.DB.Where("email = ?", email).First(&user).Error; err == nil {
			return user.OTP, nil
		}
	case "doctor":
		if err := configs.DB.Where("email = ?", email).First(&doctor).Error; err == nil {
			return doctor.OTP, nil
		}
	case "admin":
		if err := configs.DB.Where("email = ?", email).First(&admin).Error; err == nil {
			return admin.OTP, nil
		}
	default:
		return "", errors.New("invalid user type")
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
