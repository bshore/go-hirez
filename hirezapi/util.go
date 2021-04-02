package hirezapi

import (
	"regexp"
)

const consolePattern string = `^.*(xbox|ps4).*$`

var (
	regexConsole = regexp.MustCompile(consolePattern)
)

// IsPCOnly compares the basePath against a console-specific regex
// If the basePath contains 'xbox' or 'ps4' we return false
func IsPCOnly(basePath string) bool {
	return !regexConsole.MatchString(basePath)
}