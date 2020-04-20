package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFor(t *testing.T) {
	vwoInstance := GetInstance()

	accountID := 12345
	userID := "Chris"
	actual := generateFor(vwoInstance, userID, accountID)
	expected := "810E5C0420C2541C8E9BDACD245E6476"
	assert.Equal(t, expected, actual, "Expected and Actual UUIDs should be same")

	accountID = 12345
	userID = "__123__"
	actual = generateFor(vwoInstance, userID, accountID)
	expected = "50A5B167FB6356A796F91D8951E480EE"
	assert.Equal(t, expected, actual, "Expected and Actual UUIDs should be same")

	accountID = 12345
	userID = "We@#dcs3232.f3"
	actual = generateFor(vwoInstance, userID, accountID)
	expected = "AAB4580A6BB3525FAA31DC341752D501"
	assert.Equal(t, expected, actual, "Expected and Actual UUIDs should be same")
}
