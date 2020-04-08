package utils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/stretchr/testify/assert"
)

func TestGetVariableValueForVariation(t *testing.T) {
	var campaign1 schema.Campaign
	if err := json.Unmarshal([]byte(Campaign2), &campaign1); err != nil {
		fmt.Println(err)
	}
	var campaign2 schema.Campaign
	if err := json.Unmarshal([]byte(Campaign1), &campaign2); err != nil {
		fmt.Println(err)
	}

	variation := campaign2.Variations[0]
	variableKey := "flaot1"
	variable := GetVariableValueForVariation(campaign2, variation, variableKey)
	assert.Empty(t, variable, "Expected object should be empty")

	variation = campaign1.Variations[1]
	variableKey = "bool2"
	variable = GetVariableValueForVariation(campaign1, variation, variableKey)
	actual := variable.ID
	expected := 4
	assert.Equal(t, expected, actual, "Expected and Actual IDs should be same")

	variation = campaign1.Variations[2]
	variableKey = "string2"
	variable = GetVariableValueForVariation(campaign1, variation, variableKey)
	actual = variable.ID
	expected = 3
	assert.Equal(t, expected, actual, "Expected and Actual IDs should be same")

}

func TestGetVariableForFeature(t *testing.T) {
	var campaign schema.Campaign
	if err := json.Unmarshal([]byte(Campaign1), &campaign); err != nil {
		fmt.Println(err)
	}
	variableKey := "int2"
	variable := GetVariableForFeature(campaign, variableKey)
	actual := variable.ID
	expected := 1
	assert.Equal(t, expected, actual, "Expected and Actual IDs should be same")
	variableKey = "float2"
	variable = GetVariableForFeature(campaign, variableKey)
	actual = variable.ID
	expected = 1
	assert.NotEqual(t, expected, actual, "Expected and Actual IDs should not be same")
}
