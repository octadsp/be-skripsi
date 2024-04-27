package validation_error

import "github.com/go-playground/validator/v10"

func GetValidationErrors(err error) []struct {
	Key string
	Tag string
} {
	validationErrors := []struct {
		Key string
		Tag string
	}{}

	for _, err := range err.(validator.ValidationErrors) {
		validationErrors = append(validationErrors, struct {
			Key string
			Tag string
		}{
			Key: err.Field(),
			Tag: err.Tag(),
		})
	}

	return validationErrors
}
