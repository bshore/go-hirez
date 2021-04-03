package hirezapi

import (
	"regexp"
)

const (
	consolePattern string = `^.*(xbox|ps4).*$`
	isConsoleErrMsg string = "is only supported for PC requests"

  smitePattern string = `^.*(smite).*$`
	notSmiteErrMsg string = "is only supported for Smite requests"

	notRankedErrMsg = "only applies to queues: Joust1v1Ranked, Joust3v3Ranked, and ConquestRanked (440, 450, and 451)"
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

func IsNotRanked(queueID string) bool {
	return !Contains(
		[]string{
			Joust1v1Ranked, // 440
			Joust3v3Ranked, // 450
			ConquestRanked, // 451
		}, queueID)
}

func Contains(a []string, b string) bool {
	for _, v := range a {
		if b == v {
			return true
		}
	}
	return false
}