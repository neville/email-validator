package module

import (
	"net"
	"regexp"
	"strings"
)

// ValidateFormat ...
func ValidateFormat(email string) (matched bool, err error) {
	matched, err = regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	if err != nil {
		return false, err
	}

	return matched, nil
}

// ValidateDomain ...
func ValidateDomain(email string) (bool, error) {
	domain := extractDomain(email)

	ipAddresses, err := net.LookupIP(domain)
	if err != nil {
		return false, err
	}
	if len(ipAddresses) > 0 {
		return true, nil
	}

	return false, nil
}

// ValidateSMTP ...
func ValidateSMTP(email string) (bool, error) {
	//domain := extractDomain(email)

	return false, nil
}

func extractDomain(email string) (domain string) {
	addressSymbolIndex := strings.Index(email, "@")
	if addressSymbolIndex == -1 {
		return
	}

	domain = email[addressSymbolIndex+1:]
	return domain
}
