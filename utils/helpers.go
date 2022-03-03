package utils

import (
	"strings"
	"time"

	"github.com/bshore/go-hirez/models"
)

const (
	// NotSmiteErrMsg is an error message from making non Smite requests to the Smite URL
	NotSmiteErrMsg = "is only available for Smite requests"
	// NotPaladinsErrMsg is an error message from making non Paladins requests to the Paladins URL
	NotPaladinsErrMsg = "is only available for Paladins requests"
	// NotRankedErrMsg is an error message from making a request to an endpoint marked as "Ranked Only"
	NotRankedErrMsg = "only applies to queues: Joust1v1Ranked, Joust3v3Ranked, and ConquestRanked (440, 450, and 451)"
)

// IsNotSmitePath checks that the client BasePath is for Smite
func IsSmitePath(basePath string) bool {
	return strings.Contains(basePath, "smite")
}

// IsPaladinsPath checks that the client BasePath is for Paladins
func IsPaladinsPath(basePath string) bool {
	return strings.Contains(basePath, "paladins")
}

// IsNotRanked returns true if the passed in queueID is not a ranked queueID
func IsNotRanked(queueID string) bool {
	return !Contains(
		[]string{
			models.Joust1v1Ranked, // 440
			models.Joust3v3Ranked, // 450
			models.ConquestRanked, // 451
		}, queueID)
}

// Contains checks if string slice 'a' contains string 'b'
func Contains(a []string, b string) bool {
	for _, v := range a {
		if b == v {
			return true
		}
	}
	return false
}

// FormatTime formats a go timestamp to the format expected/returned by the HiRez API
func FormatTime(t time.Time) string {
	return t.Format(models.TimeFormat)
}
