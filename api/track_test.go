package api

import (

	// "github.com/decabits/vwo-golang-sdk/schema"
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/stretchr/testify/assert"
)

func TestTrack(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")

	userID := "Varun"
	campaignKey := "notPresent"
	goalIdentifier := ""
	value := Track(vwoInstance, campaignKey, userID, goalIdentifier)
	assert.False(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab5"
	goalIdentifier = ""
	value = Track(vwoInstance, campaignKey, userID, goalIdentifier)
	assert.False(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "phpab4"
	goalIdentifier = ""
	value = Track(vwoInstance, campaignKey, userID, goalIdentifier)
	assert.False(t, value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "php2"
	goalIdentifier = "test"
	value = Track(vwoInstance, campaignKey, userID, goalIdentifier)
	assert.False(t, value, "Goal Not Found")

	userID = "Liza"
	campaignKey = "phpab4"
	goalIdentifier = "custom"
	value = Track(vwoInstance, campaignKey, userID, goalIdentifier)
	assert.False(t, value, "Invalid Goal type")

	userID = "Liza"
	campaignKey = "phpab4"
	goalIdentifier = "rev"
	value = Track(vwoInstance, campaignKey, userID, goalIdentifier)
	assert.False(t, value, "Revenue Not defined")

	/*
		Comment out line 45 for testing, It allocates ranges to variation
	*/

	options := schema.Options{
		RevenueGoal: 12,
	}
	userID = "Misty"
	campaignKey = "phpab6"
	// fmt.Println(utils.CheckCampaignType(utils.GetCampaign(vwoInstance.SettingsFile, campaignKey), constants.CampaignTypeFeatureRollout))
	goalIdentifier = "rev"
	value = TrackWithOptions(vwoInstance, campaignKey, userID, goalIdentifier, options)
	assert.True(t, value, "Variation should be defined")

}
