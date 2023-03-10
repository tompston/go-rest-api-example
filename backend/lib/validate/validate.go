package validate

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Message     string `json:"message"`
}

// Return an understandabe / UI friendly message that potentially
// can be displayed on the client side, when validation for
// a field fails. Be free to add new tags!
func messageForTag(fe validator.FieldError) string {

	field := fe.Field()
	param := fe.Param()
	tag := fe.Tag()

	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email address"
	case "min":
		return field + " should be at least " + param + " characters long"
	case "max":
		return field + " should be less than " + param + " characters long"
	}
	return field + " failed validation for " + tag + ", " + param
}

// Pass in the sent struct and check if the
// fields pass validation
func ValidateStruct(i interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validator.New().Struct(i)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorResponse{
				FailedField: err.Field(),
				Message:     messageForTag(err),
			})
		}
	}
	return errors
}
