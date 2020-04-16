package api

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsFeatureEnabled(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")
	options := schema.Options{}

	userID := "Varun"
	campaignKey := "notPresent"
	value := IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab5"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "phpab2"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "Campaign Not Valid")

	userID = "Robbie"
	campaignKey = "php4"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "No Variation from campaign Not alloted")

	userID = "Gimmy"
	campaignKey = "php2"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "Campaign Feature Test")

	userID = "Kate"
	campaignKey = "phpab6"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.True(t, value, "Campaign Feature Test")

	userID = "Kate"
	campaignKey = "phpab4"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.True(t, value, "Campaign Feature Rollout")
}
