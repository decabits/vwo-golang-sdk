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
	vwoInstance := GetInstance("../settingsFile.json")

	campaignKey := "phpab3"
	campaign, _ := GetCampaign(vwoInstance.SettingsFile, campaignKey)
	assert.Equal(t, vwoInstance.SettingsFile.Campaigns[3], campaign, "Expected and Actual Campaign IDs should be same")

	campaignKey = "p007"
	campaign, _ = GetCampaign(vwoInstance.SettingsFile, campaignKey)
	assert.Empty(t, campaign, "Expected campaign should be empty")
}

func TestGetCampaignVariation(t *testing.T) {
	vwoInstance := GetInstance("../settingsFile.json")
	campaign := vwoInstance.SettingsFile.Campaigns[1]

	variationName := "Control"
	variation, _ := GetCampaignVariation(campaign, variationName)
	assert.Equal(t, campaign.Variations[0], variation, "Expected and Actual Variation IDs should be same")

	variationName = "Variation-3"
	variation, _ = GetCampaignVariation(campaign, variationName)
	assert.Empty(t, variation, "Expected Variation should be empty")
}

func TestGetCampaignGoal(t *testing.T) {
	vwoInstance := GetInstance("../settingsFile.json")
	campaign := vwoInstance.SettingsFile.Campaigns[1]

	goalName := "rev"
	goal, _ := GetCampaignGoal(campaign, goalName)
	assert.Equal(t, campaign.Goals[0], goal, "Expected and Actual Goal IDs should be same")

	goalName = "demo"
	goal, _ = GetCampaignGoal(campaign, goalName)
	assert.Empty(t, goal, "Expected Goal should be empty")
}

func TestGetControlVariation(t *testing.T) {
	vwoInstance := GetInstance("../settingsFile.json")

	campaign := vwoInstance.SettingsFile.Campaigns[2]
	variation := GetControlVariation(campaign)
	assert.Equal(t, campaign.Variations[0], variation, "Expected variation should be present in the campaign")

	campaign = vwoInstance.SettingsFile.Campaigns[3]
	variation = GetControlVariation(campaign)
	assert.Empty(t, variation, "Expected variation should be empty")
}

func TestScaleVariations(t *testing.T) {
	vwoInstance := GetInstance("../settingsFile.json")

	variations := vwoInstance.SettingsFile.Campaigns[0].Variations
	variations = ScaleVariations(variations)
	assert.Equal(t, vwoInstance.SettingsFile.Campaigns[0].Variations, variations, "List of variations did not match")
	assert.Equal(t, 50.0, variations[0].Weight, "Variation weight did not match")
	assert.Equal(t, 50.0, variations[1].Weight, "Variation weight did not match")

	variations = GetInstance("../settingsFile.json").SettingsFile.Campaigns[2].Variations
	variations = ScaleVariations(variations)
	vwoInstance = GetInstance("../settingsFile.json")
	assert.Equal(t, vwoInstance.SettingsFile.Campaigns[2].Variations, variations, "List of variations did not match")
}

func TestGetVariationAllocationRanges(t *testing.T) {
	vwoInstance := GetInstance("../settingsFile.json")

	variations := vwoInstance.SettingsFile.Campaigns[0].Variations

	assert.NotEmpty(t, variations, "No Variations recieved")

	startVal := 1
	endVal := 1
	for _, variation := range variations {
		Range := GetVariationBucketingRange(variation.Weight)
		assert.Equal(t, startVal, variation.StartVariationAllocation, "Start Allocation range failed to match")
		endVal = startVal + Range - 1
		assert.Equal(t, endVal, variation.EndVariationAllocation, "End Allocation range failed to match")
		startVal += Range
	}
}
