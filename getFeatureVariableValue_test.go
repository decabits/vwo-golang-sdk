package vwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFeatureVariableValue(t *testing.T) {
	assertOutput := assert.New(t)

	vwoInstance := VWOInstance{}
	err := vwoInstance.getInstance("./testdata/testGetFeatureVariableValue.json")
	assertOutput.Nil(err, "error fetching instance")

	userID := ""
	campaignKey := ""
	variableKey := ""
	value := vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	assertOutput.Nil(value, "Invalid params")

	userID = "Varun"
	campaignKey = "notPresent"
	variableKey = "float2"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	assertOutput.Nil(value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "php1"
	variableKey = "float2"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	assertOutput.Nil(value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "phpab1"
	variableKey = "float2"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	assertOutput.Nil(value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "php2"
	variableKey = "float2"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	assertOutput.Nil(value, "Variation Not alloted as none exist")

	userID = "Gimmy"
	campaignKey = "php3"
	variableKey = "string1"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	assertOutput.Nil(value, "No variable with name found")

	userID = "Kate"
	campaignKey = "php3"
	variableKey = "float2"
	actual := vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	expected := 10.67
	assertOutput.Equal(expected, actual, "Value mismatch for variable")

	userID = "Kate"
	campaignKey = "php4"
	variableKey = "bool1"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	expected1 := true
	assertOutput.Equal(expected1, value, "Value mismatch for variable")

	userID = "Kate"
	campaignKey = "php4"
	variableKey = "int1"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	expected2 := 301
	assertOutput.Equal(float64(expected2), value, "Value mismatch for variable")

	userID = "Kate"
	campaignKey = "php4"
	variableKey = "string2"
	value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	expected4 := "abcd"
	assertOutput.Equal(expected4, value, "Value mismatch for variable")
}
