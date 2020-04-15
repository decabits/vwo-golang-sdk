package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

func TestPreEvaluateSegment(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")

	segments := vwoInstance.SettingsFile.Campaigns[0].Segments
	options := schema.Options{}
	value := PreEvaluateSegment(vwoInstance, segments, options)
	assert.False(t, value, "Expected False")

	segments = vwoInstance.SettingsFile.Campaigns[1].Variations[1].Segments
	options = schema.Options{
		VWOUserID:   "Varun",
		RevenueGoal: 12,
	}
	value = PreEvaluateSegment(vwoInstance, segments, options)
	assert.True(t, value, "Expected True")
}

func TestEvaluateSegment(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")


	segments := vwoInstance.SettingsFile.Campaigns[0].Segments
	options := schema.Options{}
	value := EvaluateSegment(vwoInstance, segments, options)
	assert.False(t, value, "Expected False")

	segments = vwoInstance.SettingsFile.Campaigns[0].Variations[0].Segments
	options = schema.Options{}
	value = EvaluateSegment(vwoInstance, segments, options)
	assert.True(t, value, "Expected True")

	segments = vwoInstance.SettingsFile.Campaigns[0].Variations[1].Segments
	options = schema.Options{
		VWOUserID:   "Varun",
		RevenueGoal: 12,
	}
	value = EvaluateSegment(vwoInstance, segments, options)
	assert.True(t, value, "Expected True")
}

func TestGetWhiteListedVariationsList(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")

	options := schema.Options{}
	userID := "test"
	campaign := vwoInstance.SettingsFile.Campaigns[0]
	actual := GetWhiteListedVariationsList(vwoInstance, userID, campaign, options)
	assert.Empty(t, actual, "No WhiteListed Variations Found")

	options = schema.Options{
		VWOUserID:   "Varun",
		RevenueGoal: 12,
	}
	userID = "test"
	campaign = vwoInstance.SettingsFile.Campaigns[1]
	actual = GetWhiteListedVariationsList(vwoInstance, userID, campaign, options)
	expected := campaign.Variations
	assert.Equal(t, expected, actual, "No WhiteListed Variations Found")
}

func TestFindTargetedVariation(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")
	options := schema.Options{
		VWOUserID:   "Varun",
		RevenueGoal: 12,
	}

	userID := "Varun"
	campaign := vwoInstance.SettingsFile.Campaigns[2]
	actual, _ := FindTargetedVariation(vwoInstance, userID, campaign, options)
	assert.Empty(t, actual, "Expected no variation")

	userID = "Varun"
	campaign = vwoInstance.SettingsFile.Campaigns[0]
	actual, _ = FindTargetedVariation(vwoInstance, userID, campaign, options)
	assert.Empty(t, actual, "Expected no variation")

	userID = "Varun"
	campaign = vwoInstance.SettingsFile.Campaigns[4]
	actual, _ = FindTargetedVariation(vwoInstance, userID, campaign, options)
	expected := campaign.Variations[0]
	assert.Equal(t, expected, actual, "Expected no variation")

	userID = "Liza"
	campaign = vwoInstance.SettingsFile.Campaigns[1]
	actual, _ = FindTargetedVariation(vwoInstance, userID, campaign, options)
	expected = campaign.Variations[1]
	assert.Equal(t, expected, actual, "Expected no variation")	
}

func TestGetVariation(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")
	options := schema.Options{
		VWOUserID:   "Liza",
		RevenueGoal: 12,
	}

	/*
	Comment line 55-63 and 71-73for testing, as UserStorage 
	testing is not included in sdk testing
	*/

	userID := "Liza"
	campaign := vwoInstance.SettingsFile.Campaigns[0]
	actual, _ := GetVariation(vwoInstance, userID, campaign, options)
	expected := campaign.Variations[0]
	assert.Equal(t, expected, actual, "Variation mis match")
	
	userID = "Varun"
	campaign = vwoInstance.SettingsFile.Campaigns[4]
	actual, _ = GetVariation(vwoInstance, userID, campaign, options)
	expected = campaign.Variations[0]
	assert.Equal(t, expected, actual, "Variation mis match")

	userID = "Gimmy"
	campaign = vwoInstance.SettingsFile.Campaigns[4]
	actual, _ = GetVariation(vwoInstance, userID, campaign, options)
	assert.Empty(t, actual, "No variation will be alloted")

	userID = "Misty"
	campaign = vwoInstance.SettingsFile.Campaigns[3]
	actual, _ = GetVariation(vwoInstance, userID, campaign, options)
	assert.Empty(t, actual, "No variation will be alloted")

	userID = "Robbie"
	campaign = vwoInstance.SettingsFile.Campaigns[7]
	actual, _ = GetVariation(vwoInstance, userID, campaign, options)
	assert.Empty(t, actual, "No variation will be alloted")

}