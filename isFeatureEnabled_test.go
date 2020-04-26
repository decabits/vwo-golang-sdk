package vwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFeatureEnabled(t *testing.T) {
	assertOutput := assert.New(t)

	vwoInstance := VWOInstance{}
	err := vwoInstance.getInstance("./testdata/testIsFeatureEnabled.json")
	assertOutput.Nil(err, "error fetching instance")

	userID := ""
	campaignKey := ""
	value := vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	assertOutput.False(value, "Invalid params")

	userID = "Varun"
	campaignKey = "notPresent"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	assertOutput.False(value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "php1"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	assertOutput.False(value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "phpab1"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	assertOutput.False(value, "Campaign Not Valid")

	userID = "Robbie"
	campaignKey = "php2"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	assertOutput.False(value, "No Variation from campaign Not alloted")

	userID = "Gimmy"
	campaignKey = "php3"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	assertOutput.True(value, "Campaign Feature Rollout")

	userID = "Kate"
	campaignKey = "php4"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	assertOutput.True(value, "Campaign Feature Test")

	userID = "Gimmy"
	campaignKey = "php4"
	value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	assertOutput.False(value, "Campaign Feature Rollout")
}
