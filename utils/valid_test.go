package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateActivate(t *testing.T) {
	actual := ValidateActivate("", "")
	assert.False(t, actual)

	actual = ValidateActivate("campaignKey", "userID")
	assert.True(t, actual)
}

func TestValidateGetFeatureVariableValue(t *testing.T) {
	actual := ValidateGetFeatureVariableValue("", "", "")
	assert.False(t, actual)

	actual = ValidateGetFeatureVariableValue("campaignKey", "variableKey", "userID")
	assert.True(t, actual)
}

func TestValidateGetVariationName(t *testing.T) {
	actual := ValidateGetVariationName("", "")
	assert.False(t, actual)

	actual = ValidateGetVariationName("campaignKey", "userID")
	assert.True(t, actual)
}

func TestValidateIsFeatureEnabled(t *testing.T) {
	actual := ValidateIsFeatureEnabled("", "")
	assert.False(t, actual)

	actual = ValidateIsFeatureEnabled("campaignKey", "userID")
	assert.True(t, actual)
}

func TestValidatePush(t *testing.T) {
	actual := ValidatePush("", "", "")
	assert.False(t, actual)

	actual = ValidatePush("tagKey", "tagValue", "userID")
	assert.True(t, actual)
}

func TestValidateTrack(t *testing.T) {
	actual := ValidateTrack("", "", "")
	assert.False(t, actual)

	actual = ValidateTrack("campaignKey", "userID", "goalIdentifier")
	assert.True(t, actual)
}
