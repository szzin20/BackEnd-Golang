package helper

func GenderAndBloodTypeIsValid(gender, bloodType string) bool {
	if (gender == "female" || gender == "male") && (bloodType == "A" || bloodType == "B" || bloodType == "AB" || bloodType == "O") {
		return true
	}
	return false
}
