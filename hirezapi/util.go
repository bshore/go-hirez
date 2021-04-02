package hirezapi

import (
	"regexp"
)

const (
	consolePattern string = `^.*(xbox|ps4).*$`
	isConsoleErrMsg string = "is only supported for PC requests"

  smitePattern string = `^.*(smite).*$`
	notSmiteErrMsg string = "is only supported for Smite requests"
)

var (
	regexConsole = regexp.MustCompile(consolePattern)
	regexSmite = regexp.MustCompile(smitePattern)
)

// IsConsolePath compares the basePath against a console-specific regex
func IsConsolePath(basePath string) bool {
	return regexConsole.MatchString(basePath)
}

// IsNotSmitePath compares the basePath against a smite-specific regex
func IsNotSmitePath(basePath string) bool {
	return !regexSmite.MatchString(basePath)
}