package api

import (
	"github.com/decabits/vwo-golang-sdk/schema"
	"testing"

	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetFeatureVariableValue(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")
	options := schema.Options{}

	userID := "Varun"
	campaignKey := "notPresent"
	variableKey := ""
	value := GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	assert.Nil(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab5"
	variableKey = ""
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	assert.Nil(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "phpab5"
	variableKey = ""
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	assert.Nil(t, value, "Campaign Not Valid")

	/*
	Comment out line 27 for testing
	*/

	userID = "Liza"
	campaignKey = "phpab3"
	variableKey = ""
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	assert.Nil(t, value, "Variation Not alloted")

	userID = "Gimmy"
	campaignKey = "php3"
	variableKey = "string1"
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	assert.Nil(t, value, "No variable with name found")

	userID = "Gimmy"
	campaignKey = "phpab4"
	variableKey = "int1"
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	expected1 := 301
	assert.Equal(t,expected1, value, "Variable Not found")

	userID = "Gimmy"
	campaignKey = "php3"
	variableKey = "string2"
	value = GetFeatureVariableValue(vwoInstance, campaignKey, variableKey, userID, options)
	expected2 := "abcd"
	assert.Equal(t,expected2, value, "Variable Not found")

	
}
