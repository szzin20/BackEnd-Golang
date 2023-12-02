package helper

import (
	"errors"
	"fmt"
)

type OTPInfo struct {
	Email string
	OTP   interface{} 
}

var storedOTPs []OTPInfo

// VerifyOTPByEmail .
func VerifyOTPByEmail(email string, providedOTP interface{}) error {
	storedOTP, err := getStoredOTP(email)
	if err != nil {
		return err
	}

	storedOTPStr := fmt.Sprint(storedOTP.OTP)
	providedOTPStr := fmt.Sprint(providedOTP)

	if storedOTPStr != providedOTPStr {
		return errors.New("invalid OTP: provided OTP does not match stored OTP")
	}

	return nil
}

func getStoredOTP(email string) (OTPInfo, error) {
	for _, otpInfo := range storedOTPs {
		if otpInfo.Email == email {
			return otpInfo, nil
		}
	}

	return OTPInfo{}, errors.New("OTP not found for the given email")
}
