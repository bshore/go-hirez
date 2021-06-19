package utils

import (
	"strings"

	"github.com/bshore/go-hirez/models"
)


const (
	NotSmiteErrMsg    = "is only available for Smite requests"
	NotPaladinsErrMsg = "is only available for Paladins requests"
	NotRankedErrMsg   = "only applies to queues: Joust1v1Ranked, Joust3v3Ranked, and ConquestRanked (440, 450, and 451)"
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
			models.Joust1v1Ranked, // 440
			models.Joust3v3Ranked, // 450
			models.ConquestRanked, // 451
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