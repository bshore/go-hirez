package utils_test

import (
	"testing"

	"github.com/bshore/go-hirez/models"
	"github.com/bshore/go-hirez/utils"
)

func TestGenerateDesiredOutput(t *testing.T) {
	tests := []struct {
		name    string
		desired interface{}
	}{
		{
			name:    "slice of players",
			desired: []models.Player{},
		},
	}
	for _, tt := range tests {
		out, err := utils.GenerateDesiredOutput(tt.desired)
		if err != nil {
			t.Logf("test %s failed - error: %v", tt.name, err)
			t.Fail()
		}
		_ = out
	}
}
