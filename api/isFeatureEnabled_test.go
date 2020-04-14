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
	campaignKey = "phpab5"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "Campaign Not Valid")

	/*
	Comment out line 28 for testing
	*/

	userID = "Robbie"
	campaignKey = "php3"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.False(t, value, "Campaign Not Valid")

	userID = "Kate"
	campaignKey = "php3"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.True(t, value, "Campaign Not Valid")

	userID = "Kate"
	campaignKey = "phpab4"
	value = IsFeatureEnabled(vwoInstance, campaignKey, userID, options)
	assert.True(t, value, "Campaign Not Valid")

}
