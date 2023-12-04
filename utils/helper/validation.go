package helper

import (
	"healthcare/models/schema"
	"time"
)

func GenderIsValid(gender string) bool {
	return gender == "female" || gender == "male" || gender == ""
}

func BloodTypeIsValid(bloodType string) bool {
	return bloodType == "A" || bloodType == "B" || bloodType == "AB" || bloodType == "O" || bloodType == ""
}

func BirthdateIsValid(birthdate string) bool {
	if birthdate == "" {
		return true
	}
	_, err := time.Parse("2006-01-02", birthdate)
	return err == nil
}

func PaymentMethodIsValid(paymentMethod string) bool {
	return paymentMethod == "manual transfer bni" || paymentMethod == "manual transfer bca" || paymentMethod == "manual transfer bri" || paymentMethod == ""
}

func PaymentStatusIsValid(paymentStatus string) bool {
	return paymentStatus == "pending" || paymentStatus == "success" || paymentStatus == "cancelled" || paymentStatus == ""
}

func GetMessageContent(lastMessage schema.Message) string {
	if lastMessage.Audio != "" {
		return "Audio Message"
	} else if lastMessage.Image != "" {
		return "Image Message"
	} else {
		return lastMessage.Message
	}
}
