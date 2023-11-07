package helper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateBalance(request interface{}) error {
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			message := fmt.Sprintf("field '%s' failed '%s' validation", field, tag)
			validationErrors = append(validationErrors, message)
		}
		return errors.New(strings.Join(validationErrors, "; "))
	}
	return nil
}
