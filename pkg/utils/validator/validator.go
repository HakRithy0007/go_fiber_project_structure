package custom_validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type ValidateResponse struct {
	FailedFiels string
	Tag         string
	Value       string
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func init() {
	validate = validator.New()
}

func ValidatePaging(s interface{}) error {
	return validate.Struct(s)
}

func ValidateStruct(value interface{}) []*ValidateResponse {
	var errors []*ValidateResponse
	validate := validator.New()
	err := validate.Struct(value)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidateResponse
			element.FailedFiels = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateStructPaging(value interface{}) ([]string, error) {
	err := validate.Struct(value)

	if err != nil {
		if validationError, ok := err.(validator.ValidationErrors); ok {
			var errs []string
			for _, validationError := range validationError {
				errs = append(errs, fmt.Sprintf(
					"Fieled '%s' failed on the '%s' tag",
					validationError.Field(),
					validationError.Tag(),
				),
				)
			}
			return errs, err
		}
	}
	return nil, nil
}
