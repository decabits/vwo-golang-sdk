package api

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsFeatureEnabled(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFiles/settings4.json")
	options := schema.Options{}

	userID := "Varun"
	campaignKey := "notPresent"
	value := IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "php1"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "phpab1"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "Campaign Not Valid")

	userID = "Robbie"
	campaignKey = "php2"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "No Variation from campaign Not alloted")

	userID = "Gimmy"
	campaignKey = "php3"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.True(t, value, "Campaign Feature Rollout")

	userID = "Kate"
	campaignKey = "php4"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.True(t, value, "Campaign Feature Test")

	userID = "Gimmy"
	campaignKey = "php4"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "Campaign Feature Rollout")
}
