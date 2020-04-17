package utils

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/stretchr/testify/assert"
)

func TestCheckCampaignType(t *testing.T) {
	vwoInstance := GetInstance("../settingsFiles/settings5.json")

	campaign := vwoInstance.SettingsFile.Campaigns[0]
	campaignType := constants.CampaignTypeVisualAB
	value := CheckCampaignType(campaign, campaignType)
	assert.True(t, value, "Campaign should match")

	campaign = vwoInstance.SettingsFile.Campaigns[4]
	campaignType = constants.CampaignTypeFeatureTest
	value = CheckCampaignType(campaign, campaignType)
	assert.True(t, value, "Campaign should not match")

	campaign = vwoInstance.SettingsFile.Campaigns[2]
	campaignType = constants.CampaignTypeFeatureRollout
	value = CheckCampaignType(campaign, campaignType)
	assert.True(t, value, "Campaign should not match")

	campaign = vwoInstance.SettingsFile.Campaigns[2]
	campaignType = constants.CampaignTypeFeatureTest
	value = CheckCampaignType(campaign, campaignType)
	assert.False(t, value, "Campaign should not match")
}

func TestGetKeyValue(t *testing.T) {
	vwoInstance := GetInstance("../settingsFiles/settings5.json")

	segment := vwoInstance.SettingsFile.Campaigns[4].Variations[0].Segments
	actualKey, actualValue := GetKeyValue(segment)
	expectedKey := "or"
	assert.Equal(t, expectedKey, actualKey, "Expected and Actual Keys should be same")
	var Temp []interface{}
	assert.IsType(t, Temp, actualValue, "Type Mismatch")

	var tempSegment map[string]interface{}
	actualKey, actualValue = GetKeyValue(tempSegment)
	assert.Equal(t, "", actualKey, "Nil Value expected")
	assert.Nil(t, actualValue, "Nil Value expected")
}
