package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFor(t *testing.T) {
	vwoInstance := GetTempInstance()
	accountID := 89499
	userID := "Chris"

	actual := generateFor(vwoInstance, userID, accountID)
	expected := "C50CFF01A27E51F080BAA50397B18BCF"
	assert.Equal(t, expected, actual, "Expected and Actual UUIDs should be same")
}
