package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVariableValueForVariation(t *testing.T) {
	vwoInstance := GetInstance("../settingsFile.json")
  campaign := vwoInstance.SettingsFile.Campaigns[1]
  
	variation := campaign.Variations[0]
	variableKey := "string1"
	variable := GetVariableValueForVariation(vwoInstance, campaign, variation, variableKey)
	assert.Empty(t, variable, "Expected object should be empty")

	variation = campaign.Variations[1]
	variableKey = "string2"
	variable = GetVariableValueForVariation(vwoInstance, campaign, variation, variableKey)
	assert.Equal(t, campaign.Variations[0].Variables[0], variable, "Expected and Actual IDs should be same")

	variation = campaign.Variations[0]
	variableKey = "bool2"
	variable = GetVariableValueForVariation(vwoInstance, campaign, variation, variableKey)
	assert.Equal(t, variation.Variables[1], variable, "Expected and Actual IDs should be same")
}

func TestGetVariableForFeature(t *testing.T) {
	vwoInstance := GetInstance("../settingsFile.json")

	variables := vwoInstance.SettingsFile.Campaigns[0].Variables
	variableKey := "int1"
	variable := GetVariableForFeature(variables, variableKey)
	assert.Equal(t, variables[0], variable, "Expected and Actual IDs should be same")

	variableKey = "float2"
	variable = GetVariableForFeature(variables, variableKey)
	assert.Empty(t, variable, "Expected variable should be empty")
}
