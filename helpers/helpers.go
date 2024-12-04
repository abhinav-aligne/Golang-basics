package helpers

import "strings"

func ValidateUserInput(firstName string, lastName string, email string) (bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	return isValidName, isValidEmail
}
