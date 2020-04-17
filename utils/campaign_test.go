package utils

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/stretchr/testify/assert"
)

func TestGetVariationBucketingRange(t *testing.T) {
	var weight float64

	weight = 0
	actual := GetVariationBucketingRange(weight)
	expected := 0
	assert.Equal(t, expected, actual, "Expected and Actual Ranges should be same")

	weight = 33.333
	actual = GetVariationBucketingRange(weight)
	expected = 3334
	assert.Equal(t, expected, actual, "Expected and Actual Ranges should be same")

	weight = 102
	actual = GetVariationBucketingRange(weight)
	expected = constants.MaxTrafficValue
	assert.Equal(t, expected, actual, "Expected and Actual Ranges should be same")
}

func TestGetCampaign(t *testing.T) {
	vwoInstance := GetInstance("../settingsFiles/settings5.json")

	campaignKey := "phpab1"
	campaign, _ := GetCampaign(vwoInstance.SettingsFile, campaignKey)
	assert.Equal(t, vwoInstance.SettingsFile.Campaigns[0], campaign, "Expected and Actual Campaign IDs should be same")

	campaignKey = "notAvailable"
	campaign, _ = GetCampaign(vwoInstance.SettingsFile, campaignKey)
	assert.Empty(t, campaign, "Expected campaign should be empty")
}

func TestGetCampaignVariation(t *testing.T) {
	vwoInstance := GetInstance("../settingsFiles/settings5.json")
	campaign := vwoInstance.SettingsFile.Campaigns[1]

	variationName := "Control"
	variation, _ := GetCampaignVariation(campaign, variationName)
	assert.Equal(t, campaign.Variations[0], variation, "Expected and Actual Variation IDs should be same")

	variationName = "Variation-3"
	variation, _ = GetCampaignVariation(campaign, variationName)
	assert.Empty(t, variation, "Expected Variation should be empty")

	campaign = vwoInstance.SettingsFile.Campaigns[0]
	variationName = "Control"
	variation, _ = GetCampaignVariation(campaign, variationName)
	assert.Empty(t, variation, "Expected and Actual Variation IDs should be same")
}

func TestGetCampaignGoal(t *testing.T) {
	vwoInstance := GetInstance("../settingsFiles/settings5.json")
	campaign := vwoInstance.SettingsFile.Campaigns[1]

	goalName := "rev"
	goal, _ := GetCampaignGoal(campaign, goalName)
	assert.Equal(t, campaign.Goals[0], goal, "Expected and Actual Goal IDs should be same")

	goalName = "demo"
	goal, _ = GetCampaignGoal(campaign, goalName)
	assert.Empty(t, goal, "Expected Goal should be empty")
}

func TestGetControlVariation(t *testing.T) {
	vwoInstance := GetInstance("../settingsFiles/settings5.json")

	campaign := vwoInstance.SettingsFile.Campaigns[1]
	variation := GetControlVariation(campaign)
	assert.Equal(t, campaign.Variations[0], variation, "Expected variation should be present in the campaign")

	campaign = vwoInstance.SettingsFile.Campaigns[2]
	variation = GetControlVariation(campaign)
	assert.Empty(t, variation, "Expected variation should be empty")
}

func TestScaleVariations(t *testing.T) {
	vwoInstance := GetInstance("../settingsFiles/settings5.json")

	variations := vwoInstance.SettingsFile.Campaigns[3].Variations
	variations = ScaleVariations(variations)
	assert.Equal(t, vwoInstance.SettingsFile.Campaigns[3].Variations, variations, "List of variations did not match")
	assert.Equal(t, 100.0, variations[0].Weight, "Variation weight did not match")

	variations = vwoInstance.SettingsFile.Campaigns[4].Variations
	variations = ScaleVariations(variations)
	assert.Equal(t, vwoInstance.SettingsFile.Campaigns[4].Variations, variations, "List of variations did not match")
}

func TestGetVariationAllocationRanges(t *testing.T) {
	vwoInstance := GetInstance("../settingsFiles/settings5.json")

	variations := vwoInstance.SettingsFile.Campaigns[3].Variations
	assert.NotEmpty(t, variations, "No Variations recieved")
	variations = GetVariationAllocationRanges(vwoInstance, variations)
	assert.Equal(t, 1, variations[0].StartVariationAllocation, "Value Mismatch")
	assert.Equal(t, 10000, variations[0].EndVariationAllocation, "Value Mismatch")

	variations = vwoInstance.SettingsFile.Campaigns[4].Variations
	assert.NotEmpty(t, variations, "No Variations recieved")
	variations = GetVariationAllocationRanges(vwoInstance, variations)
	assert.Equal(t, -1, variations[0].StartVariationAllocation, "Start Allocation range failed to match")
	assert.Equal(t, -1, variations[0].EndVariationAllocation, "End Allocation range failed to match")

}

func TestMin(t *testing.T) {
	assert.Equal(t, 10, min(10,20), "Incorrect")
	assert.Equal(t, 10, min(20,10), "Incorrect")
	assert.NotEqual(t, 12, min(10,20), "Incorrect")
}

func TestMax(t *testing.T) {
	assert.Equal(t, 20, max(10,20), "Incorrect")
	assert.Equal(t, 20, max(20,10), "Incorrect")
	assert.NotEqual(t, 12, max(10,20), "Incorrect")	
}