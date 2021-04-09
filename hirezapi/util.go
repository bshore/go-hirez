package hirezapi

import (
	"strings"
)

const (
	notSmiteErrMsg = "is only available for Smite requests"
	notPaladinsErrMsg = "is only available for Paladins requests"
	notRankedErrMsg = "only applies to queues: Joust1v1Ranked, Joust3v3Ranked, and ConquestRanked (440, 450, and 451)"
)

// IsNotSmitePath checks that the client BasePath is for Smite
func IsSmitePath(basePath string) bool {
	return strings.Contains(basePath, "smite")
}

// IsPaladinsPath checks that the client BasePath is for Paladins
func IsPaladinsPath(basePath string) bool {
	return strings.Contains(basePath, "paladins")
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