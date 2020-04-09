package utils

import (
	"math"
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
	vwoInstance := GetTempInstance()

	campaignKey := "php2"
	campaign, _ := GetCampaign(vwoInstance.SettingsFile, campaignKey)
	actual := campaign.ID
	expected := 283
	assert.Equal(t, expected, actual, "Expected and Actual Campaign IDs should be same")

	campaignKey = "php1"
	campaign, _ = GetCampaign(vwoInstance.SettingsFile, campaignKey)
	assert.Empty(t, campaign, "Expected campaign should be empty")
}

func TestGetCampaignVariation(t *testing.T) {
	vwoInstance := GetTempInstance()

	campaign := vwoInstance.SettingsFile.Campaigns[0]
	variationName := "Control"
	variation, _ := GetCampaignVariation(campaign, variationName)
	expected := variation.ID
	actual := 1
	assert.Equal(t, expected, actual, "Expected and Actual Variation IDs should be same")

	variationName = "Variation-3"
	variation, _ = GetCampaignVariation(campaign, variationName)
	assert.Empty(t, variation, "Expected Variation should be empty")
}

func TestGetCampaignGoal(t *testing.T) {
	vwoInstance := GetTempInstance()

	campaign := vwoInstance.SettingsFile.Campaigns[0]
	goalName := "rev"
	goal, _ := GetCampaignGoal(campaign, goalName)
	expected := goal.ID
	actual := 2
	assert.Equal(t, expected, actual, "Expected and Actual Goal IDs should be same")

	goalName = "demo"
	goal, _ = GetCampaignGoal(campaign, goalName)
	assert.Empty(t, goal, "Expected Goal should be empty")
}

func TestGetControlVariation(t *testing.T) {
	vwoInstance := GetTempInstance()

	campaign := vwoInstance.SettingsFile.Campaigns[0]
	variation := GetControlVariation(campaign)
	assert.NotEmpty(t, variation, "Expected variation should be present in the campaign")
}

func TestScaleVariations(t *testing.T) {
	vwoInstance := GetTempInstance()

	variations := vwoInstance.SettingsFile.Campaigns[0].Variations
	variations = ScaleVariations(variations)
	actualWeightSum := 0.0
	for _, variation := range variations {
		actualWeightSum += variation.Weight
	}
	expectedWeightSum := 100.0
	assert.Equal(t, expectedWeightSum, math.Ceil(actualWeightSum), "Sum of weights should be hundred")

	for i := range variations {
		variations[i].Weight = 0.0
	}
	variations = ScaleVariations(variations)
	actualWeightSum = 0.0
	for _, variation := range variations {
		actualWeightSum += variation.Weight
	}
	expectedWeightSum = 100.0
	assert.Equal(t, expectedWeightSum, math.Ceil(actualWeightSum), "Sum of weights should be hundred")
}

func TestGetVariationAllocationRanges(t *testing.T) {
	vwoInstance := GetTempInstance()

	variations := vwoInstance.SettingsFile.Campaigns[0].Variations
	variations = GetVariationAllocationRanges(vwoInstance, variations)

	assert.NotEmpty(t, variations, "No Variations recieved")
  
	startVal := 1
	endVal := 1
	for _, variation := range variations {
		Range := GetVariationBucketingRange(variation.Weight)
		assert.Equal(t, startVal, variation.StartVariationAllocation, "Start Allocation range failed to match")
		endVal = min(startVal+Range-1, constants.MaxTrafficValue)
		assert.Equal(t, endVal, variation.EndVariationAllocation, "End Allocation range failed to match")
		startVal += Range
	}
}
