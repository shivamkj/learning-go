package validator

import (
	"fmt"
	"reflect"
)

func (e emailValidator) Validate(tag string, field reflect.StructField, value any) string {
	if e.tagRegex.MatchString(tag) {
		email, ok := value.(string)
		if !ok {
			return fmt.Sprintf("%s must be a valid email", field.Name)
		}

		// Simplified email validation for demonstration purposes
		// You might want to use a more sophisticated approach in a real-world scenario
		isValid := e.emailRegex.MatchString(email)
		if !isValid {
			return fmt.Sprintf("%s must be a valid email", field.Name)
		}
	}

	return ""
}

func (m minValidator) Validate(tag string, field reflect.StructField, value any) string {
	if match := m.tagRegex.FindStringSubmatch(tag); len(match) == 2 {
		min := match[1]

		if field.Type.Kind() == reflect.Int {
			if value.(int) < toInt(min) {
				return fmt.Sprintf("%s must be at least %s", field.Name, min)
			}
		} else {
			if len(fmt.Sprintf("%v", value)) < toInt(min) {
				return fmt.Sprintf("%s must be at least %s characters long", field.Name, min)
			}
		}

	}

	return ""
}

func (m maxValidator) Validate(tag string, field reflect.StructField, value any) string {
	if match := m.tagRegex.FindStringSubmatch(tag); len(match) == 2 {
		min := match[1]

		if field.Type.Kind() == reflect.Int {
			if value.(int) < toInt(min) {
				return fmt.Sprintf("%s must be at most %s", field.Name, min)
			}
		} else {
			if len(fmt.Sprintf("%v", value)) < toInt(min) {
				return fmt.Sprintf("%s must be at most %s characters long", field.Name, min)
			}
		}
	}

	return ""
}

func (m numericValidator) Validate(tag string, field reflect.StructField, value any) string {
	if m.tagRegex.MatchString(tag) {
		var isNumeric bool = false
		switch value.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			isNumeric = true
		}
		if !isNumeric {
			return fmt.Sprintf("%s must be numeric", field.Name)
		}
	}

	return ""
}

func (r requiredValidator) Validate(tag string, field reflect.StructField, value any) string {
	if r.tagRegex.MatchString(tag) {
		v := reflect.ValueOf(value)
		isEmpty := v.Interface() == reflect.Zero(v.Type()).Interface()
		if isEmpty {
			return fmt.Sprintf("%s is required", field.Name)
		}
	}

	return ""
}
