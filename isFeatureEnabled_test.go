package vwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFeatureEnabled(t *testing.T) {
	vwoInstance := VWOInstance{}
	vwoInstance.getInstance("./testData/testIsFeatureEnabled.json")

	userID := "Varun"
	campaignKey := "notPresent"
	value := vwoInstance.IsFeatureEnabled(campaignKey, userID)
	assert.False(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "php1"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID)
	assert.False(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "phpab1"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID)
	assert.False(t, value, "Campaign Not Valid")

	userID = "Robbie"
	campaignKey = "php2"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID)
	assert.False(t, value, "No Variation from campaign Not alloted")

	userID = "Gimmy"
	campaignKey = "php3"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID)
	assert.True(t, value, "Campaign Feature Rollout")

	userID = "Kate"
	campaignKey = "php4"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID)
	assert.True(t, value, "Campaign Feature Test")

	userID = "Gimmy"
	campaignKey = "php4"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID)
	assert.False(t, value, "Campaign Feature Rollout")
}
