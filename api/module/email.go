package module

import (
	"regexp"
)

// ValidateFormat ...
func ValidateFormat(email string) (matched bool, err error) {
	matched, err = regexp.MatchString("", email)
	if err != nil {
		return false, err
	}

	return matched, nil
}

// ValidateDomain ...
func ValidateDomain(domain string) (bool, error) {
	return false, nil
}

// ValidateSMTP ...
func ValidateSMTP(domain string) (bool, error) {
	return false, nil
}
