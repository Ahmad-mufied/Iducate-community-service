package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase" // Use this package to convert to snake case
	"strings"
)

// FormatValidationErrors formats validation errors into a map with snake_case field names and nested error messages.
func FormatValidationErrors(err error) map[string]interface{} {
	errors := make(map[string]interface{})

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			// Split the namespace to get nested field names
			fieldParts := strings.Split(fieldError.Namespace(), ".")

			// Skip the first part (add_book_request)
			fieldParts = fieldParts[1:]

			// Convert each part to snake_case
			for i := range fieldParts {
				fieldParts[i] = strcase.ToSnake(fieldParts[i])
			}

			// The last part is the actual field name
			fieldName := fieldParts[len(fieldParts)-1]
			nestedFields := fieldParts[:len(fieldParts)-1]

			// Construct the nested map structure
			currentMap := errors
			for _, part := range nestedFields {
				if _, exists := currentMap[part]; !exists {
					currentMap[part] = make(map[string]interface{})
				}
				currentMap = currentMap[part].(map[string]interface{})
			}

			// Add the error message to the appropriate field using snake_case fieldName
			currentMap[fieldName] = fmt.Sprintf("The %s field is %s", fieldName, fieldError.Tag())
		}
	} else {
		errors["error"] = "An unexpected error occurred during validation"
	}

	return errors
}
