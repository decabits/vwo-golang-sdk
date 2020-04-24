package vwo

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/stretchr/testify/assert"
)

func TestTrack(t *testing.T) {
	vwoInstance := VWOInstance{}
	vwoInstance.getInstance("./testData/testTrack.json")

	userID := "Varun"
	campaignKey := "notPresent"
	goalIdentifier := ""
	value := vwoInstance.Track(campaignKey, userID, goalIdentifier)
	assert.False(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab1"
	goalIdentifier = ""
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier)
	assert.False(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "php1"
	goalIdentifier = ""
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier)
	assert.False(t, value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "phpab2"
	goalIdentifier = "test"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier)
	assert.False(t, value, "Goal Not Found")

	userID = "Liza"
	campaignKey = "phpab3"
	goalIdentifier = "rev"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier)
	assert.False(t, value, "Invalid Goal type")

	userID = "Liza"
	campaignKey = "phpab3"
	options := schema.Options{
		RevenueGoal: 0,
	}
	goalIdentifier = "rev"
	value = vwoInstance.TrackWithOptions(campaignKey, userID, goalIdentifier, options)
	assert.False(t, value, "Revenue Not defined")

	userID = "Liza"
	campaignKey = "php2"
	options = schema.Options{
		RevenueGoal: 10,
	}
	goalIdentifier = "custom"
	value = vwoInstance.TrackWithOptions(campaignKey, userID, goalIdentifier, options)
	assert.False(t, value, "No Variation in Campaign")

	options = schema.Options{
		RevenueGoal: 12,
	}
	userID = "Misty"
	campaignKey = "phpab3"
	goalIdentifier = "custom"
	value = vwoInstance.TrackWithOptions(campaignKey, userID, goalIdentifier, options)
	assert.True(t, value, "Variation should be defined")
}
