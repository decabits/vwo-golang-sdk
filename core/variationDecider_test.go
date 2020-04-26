package core

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/stretchr/testify/assert"
)

func TestPreEvaluateSegment(t *testing.T) {
	vwoInstance := getInstanceWithoutStorage("./testData/testVariation.json")

	segments := vwoInstance.SettingsFile.Campaigns[0].Segments
	options := schema.Options{}
	value := PreEvaluateSegment(vwoInstance, segments, options)
	assert.False(t, value, "Expected False as no segments")

	segments = vwoInstance.SettingsFile.Campaigns[0].Variations[0].Segments
	options = schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "Kate"},
		RevenueGoal:     12,
	}
	value = PreEvaluateSegment(vwoInstance, segments, options)
	assert.True(t, value, "Expected True")

	options = schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "Misty"},
		RevenueGoal:     12,
	}
	value = PreEvaluateSegment(vwoInstance, segments, options)
	assert.False(t, value, "Expected True")
}

func TestEvaluateSegment(t *testing.T) {
	vwoInstance := getInstanceWithoutStorage("./testData/testVariation.json")

	segments := vwoInstance.SettingsFile.Campaigns[0].Segments
	options := schema.Options{}
	value := EvaluateSegment(vwoInstance, segments, options)
	assert.True(t, value, "Expected False as mismatch")

	segments = vwoInstance.SettingsFile.Campaigns[0].Variations[0].Segments
	options = schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "Varun"},
		RevenueGoal:     12,
	}
	value = EvaluateSegment(vwoInstance, segments, options)
	assert.True(t, value, "Expected True")

	options = schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "Liza"},
		RevenueGoal:     12,
	}
	value = EvaluateSegment(vwoInstance, segments, options)
	assert.False(t, value, "Expected True")
}

func TestGetWhiteListedVariationsList(t *testing.T) {
	vwoInstance := getInstanceWithoutStorage("./testData/testVariation.json")

	options := schema.Options{}
	userID := "test"
	campaign := vwoInstance.SettingsFile.Campaigns[2]
	actual := GetWhiteListedVariationsList(vwoInstance, userID, campaign, options)
	assert.Empty(t, actual, "No WhiteListed Variations Found")

	options = schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "Varun"},
		RevenueGoal:     12,
	}
	userID = "test"
	campaign = vwoInstance.SettingsFile.Campaigns[1]
	actual = GetWhiteListedVariationsList(vwoInstance, userID, campaign, options)
	expected := campaign.Variations
	assert.Equal(t, expected, actual, "No WhiteListed Variations Found")
}

func TestFindTargetedVariation(t *testing.T) {
	assertOutput := assert.New(t)

	vwoInstance := getInstanceWithoutStorage("./testData/testVariation.json")
	options := schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "Varun"},
		RevenueGoal:     12,
	}

	userID := "Varun"
	campaign := vwoInstance.SettingsFile.Campaigns[2]
	actual, _ := FindTargetedVariation(vwoInstance, userID, campaign, options)
	assertOutput.Empty(actual, "Forced variation Disabled")

	userID = "Varun"
	campaign = vwoInstance.SettingsFile.Campaigns[3]
	actual, _ = FindTargetedVariation(vwoInstance, userID, campaign, options)
	assertOutput.Empty(actual, "Expected no variation")

	userID = "Varun"
	campaign = vwoInstance.SettingsFile.Campaigns[4]
	actual, _ = FindTargetedVariation(vwoInstance, userID, campaign, options)
	expected := campaign.Variations[0]
	assertOutput.Equal(expected, actual, "Expected single variation")

	userID = "Varun"
	campaign = vwoInstance.SettingsFile.Campaigns[5]
	actual, _ = FindTargetedVariation(vwoInstance, userID, campaign, options)
	expected = campaign.Variations[0]
	assertOutput.Equal(expected, actual, "Expected single variation")

	options = schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "Gimmy"},
		RevenueGoal:     12,
	}
	userID = "Gimmy"
	campaign = vwoInstance.SettingsFile.Campaigns[6]
	actual, _ = FindTargetedVariation(vwoInstance, userID, campaign, options)
	assertOutput.Empty(actual, "No Varaition Expected ")

}

func TestGetVariation(t *testing.T) {
	assertOutput := assert.New(t)

	vwoInstance := getInstanceWithoutStorage("./testData/testVariation.json")
	options := schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "Varun"},
		RevenueGoal:     12,
	}

	userID := "Varun"
	campaign := vwoInstance.SettingsFile.Campaigns[4]
	actual, err := GetVariation(vwoInstance, userID, campaign, options)
	expected := campaign.Variations[0]
	assertOutput.Nil(err, "Variation mis match")
	assertOutput.Equal(expected, actual, "Variation mis match")

	userID = "Liza"
	campaign = vwoInstance.SettingsFile.Campaigns[0]
	actual, err = GetVariation(vwoInstance, userID, campaign, options)
	expected = campaign.Variations[0]
	assertOutput.Nil(err, "Variation not found in userStorage")
	assertOutput.Equal(expected, actual, "Variation not found in userStorage")

	userID = "Gimmy"
	campaign = vwoInstance.SettingsFile.Campaigns[7]
	actual, err = GetVariation(vwoInstance, userID, campaign, options)
	assertOutput.NotNil(err, "Variation not in campaign")
	assertOutput.Empty(actual, "Variation not in campaign")

	userID = "Gimmy"
	campaign = vwoInstance.SettingsFile.Campaigns[3]
	actual, err = GetVariation(vwoInstance, userID, campaign, options)
	assertOutput.NotNil(err, "User not eligible for campaign")
	assertOutput.Empty(actual, "User not eligible for campaign")

	vwoInstance = getInstanceWithStorage("./testData/testVariation.json")
	options = schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "Varun"},
		RevenueGoal:     12,
	}

	userID = "Varun"
	campaign = vwoInstance.SettingsFile.Campaigns[3]
	expected = campaign.Variations[0]
	actual, err = GetVariation(vwoInstance, userID, campaign, options)
	assertOutput.Nil(err, "User not eligible for campaign")
	assertOutput.Equal(expected, actual, "User not eligible for campaign")

	userID = "user1"
	campaign = vwoInstance.SettingsFile.Campaigns[0]
	actual, err = GetVariation(vwoInstance, userID, campaign, options)
	assertOutput.Nil(err, "Actual and Expected Variation Name mismatch")
	assertOutput.NotEmpty(actual, "Actual and Expected Variation Name mismatch")
}

func TestGetVariationFromUserStorage(t *testing.T) {
	assertOutput := assert.New(t)
	vwoInstance := getInstanceWithoutStorage("./testData/testVariation.json")

	campaign := vwoInstance.SettingsFile.Campaigns[0]
	userID := "Liza"
	actual, err := GetVariationFromUserStorage(vwoInstance, userID, campaign)
	assertOutput.NotNil(err, "Actual and Expected Variation Name mismatch")
	assertOutput.Empty(actual, "Actual and Expected Variation Name mismatch")

	vwoInstance = getInstanceWithStorage("./testData/testVariation.json")

	campaign = vwoInstance.SettingsFile.Campaigns[0]
	userID = "user1"
	expected := "Control"
	actual, err = GetVariationFromUserStorage(vwoInstance, userID, campaign)
	assertOutput.Nil(err, "Actual and Expected Variation Name mismatch")
	assert.Equal(t, expected, actual, "Actual and Expected Variation Name mismatch")
}
