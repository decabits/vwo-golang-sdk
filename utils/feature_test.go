package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVariableValueForVariation(t *testing.T) {
	vwoInstance := GetInstance()
	campaign := vwoInstance.SettingsFile.Campaigns[2]
	userID := ""

	variation := campaign.Variations[0]
	variableKey := "string1"
	variable := GetVariableValueForVariation(vwoInstance, campaign, variation, variableKey, userID)
	assert.Empty(t, variable, "Expected object should be empty")

	variation = campaign.Variations[0]
	variableKey = "int2"
	variable = GetVariableValueForVariation(vwoInstance, campaign, variation, variableKey, userID)
	assert.Equal(t, campaign.Variations[0].Variables[0], variable, "Expected and Actual IDs should be same")
}

func TestGetVariableForFeature(t *testing.T) {
	vwoInstance := GetInstance()

	variables := vwoInstance.SettingsFile.Campaigns[2].Variables
	variableKey := "int1"
	variable := GetVariableForFeature(variables, variableKey)
	assert.Equal(t, variables[0], variable, "Expected and Actual IDs should be same")

	variableKey = "float2"
	variable = GetVariableForFeature(variables, variableKey)
	assert.Empty(t, variable, "Expected variable should be empty")
}
