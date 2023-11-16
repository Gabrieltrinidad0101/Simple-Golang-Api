package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ExtractErrorMessages(verr validator.ValidationErrors) string {
	var messages []string
	for _, fieldErr := range verr {
		fieldName := fieldErr.Field()
		fieldType := fieldErr.Type()
		tagName := fieldErr.Tag()
		message := fmt.Sprintf("Field %s Type %s: is %s", fieldName, fieldType, tagName)
		messages = append(messages, message)
	}
	return fmt.Sprintf("[%s]", join(messages, ", "))
}

// join concatenates strings with the given separator
func join(slice []string, sep string) string {
	result := ""
	for i, s := range slice {
		result += s
		if i < len(slice)-1 {
			result += sep
		}
	}
	return result
}
