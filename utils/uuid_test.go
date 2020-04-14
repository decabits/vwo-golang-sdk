package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFor(t *testing.T) {
	vwoInstance := GetInstance("../settingsFile.json")
	
	accountID := 89499
	userID := "Chris"
	actual := generateFor(vwoInstance, userID, accountID)
	expected := "C50CFF01A27E51F080BAA50397B18BCF"
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

func TestGenerate(t *testing.T) {
	name := "Robert"
	actual := generate(parse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"),name)
	expected := parse("b0337440-9cc4-59da-9afb-70c0f468a5d1")
	assert.Equal(t, expected, actual, "UUIDs did not match")
}