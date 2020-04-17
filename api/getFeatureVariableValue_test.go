package api

import (
	"github.com/decabits/vwo-golang-sdk/schema"
	"testing"

	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetFeatureVariableValue(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFiles/settings3.json")
	options := schema.Options{}

	userID := "Varun"
	campaignKey := "notPresent"
	variableKey := ""
	value := GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	assert.Nil(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "php1"
	variableKey = ""
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	assert.Nil(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "phpab1"
	variableKey = ""
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	assert.Nil(t, value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "php2"
	variableKey = ""
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	assert.Nil(t, value, "Variation Not alloted as none exist")

	userID = "Gimmy"
	campaignKey = "php3"
	variableKey = "string1"
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	assert.Nil(t, value, "No variable with name found")

	userID = "Kate"
	campaignKey = "php3"
	variableKey = "float2"
	actual := GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	expected := 10.67
	assert.Equal(t, expected, actual, "Value mismatch for variable")

	userID = "Kate"
	campaignKey = "php4"
	variableKey = "bool1"
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	expected1 := true
	assert.Equal(t,expected1, value, "Value mismatch for variable")

	userID = "Kate"
	campaignKey = "php4"
	variableKey = "int1"
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	expected2 := 301
	assert.Equal(t,float64(expected2), value, "Value mismatch for variable")

	userID = "Kate"
	campaignKey = "php4"
	variableKey = "string2"
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	expected4 := "abcd"
	assert.Equal(t,expected4, value, "Value mismatch for variable")
}