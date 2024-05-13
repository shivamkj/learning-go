package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type validator interface {
	Validate(string, reflect.StructField, any) string
}

// Validate performs struct validation and returns error messages.
func Validate(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return errors.New("validation: input is not a struct")
	}

	var validationErrors []string

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i).Interface()

		// Validate required fields
		if tag, ok := field.Tag.Lookup("validate"); ok {

			for _, validate := range validators {
				if err := validate.Validate(tag, field, value); err != "" {
					validationErrors = append(validationErrors, err)
					continue
				}
			}

		}
	}

	if len(validationErrors) > 0 {
		return ValidationErrors{Errors: validationErrors}
	}

	return nil
}

// ValidationErrors represents the validation error messages.
type ValidationErrors struct {
	Errors []string
}

// Error returns the formatted validation error messages.
func (ve ValidationErrors) Error() string {
	return fmt.Sprintf("Validation errors:\n%s", strings.Join(ve.Errors, "\n"))
}
