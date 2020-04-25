package vwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFeatureVariableValue(t *testing.T) {
	vwoInstance := VWOInstance{}
	err := vwoInstance.getInstance("./testdata/testGetFeatureVariableValue.json")
	assert.Nil(t, err, "error fetching instance")

	userID := ""
	campaignKey := ""
	variableKey := ""
	value := vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID)
	assert.Nil(t, value, "Invalid params")

	userID = "Varun"
	campaignKey = "notPresent"
	variableKey = "float2"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID)
	assert.Nil(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "php1"
	variableKey = "float2"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID)
	assert.Nil(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "phpab1"
	variableKey = "float2"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID)
	assert.Nil(t, value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "php2"
	variableKey = "float2"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID)
	assert.Nil(t, value, "Variation Not alloted as none exist")

	userID = "Gimmy"
	campaignKey = "php3"
	variableKey = "string1"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID)
	assert.Nil(t, value, "No variable with name found")

	userID = "Kate"
	campaignKey = "php3"
	variableKey = "float2"
	actual := vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID)
	expected := 10.67
	assert.Equal(t, expected, actual, "Value mismatch for variable")

	userID = "Kate"
	campaignKey = "php4"
	variableKey = "bool1"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID)
	expected1 := true
	assert.Equal(t, expected1, value, "Value mismatch for variable")

	userID = "Kate"
	campaignKey = "php4"
	variableKey = "int1"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID)
	expected2 := 301
	assert.Equal(t, float64(expected2), value, "Value mismatch for variable")

	userID = "Kate"
	campaignKey = "php4"
	variableKey = "string2"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID)
	expected4 := "abcd"
	assert.Equal(t, expected4, value, "Value mismatch for variable")
}
