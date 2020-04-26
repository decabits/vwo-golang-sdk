package vwo

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/stretchr/testify/assert"
)

func TestTrack(t *testing.T) {
	assertOutput := assert.New(t)

	vwoInstance := VWOInstance{}
	err := vwoInstance.getInstance("./testdata/testTrack.json")
	assertOutput.Nil(err, "error fetching instance")

	userID := ""
	campaignKey := ""
	goalIdentifier := ""
	value := vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Invalid params")

	userID = "Varun"
	campaignKey = "notPresent"
	goalIdentifier = "custom"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab1"
	goalIdentifier = "custom"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "php1"
	goalIdentifier = "custom"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "phpab2"
	goalIdentifier = "test"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Goal Not Found")

	userID = "Liza"
	campaignKey = "phpab3"
	goalIdentifier = "rev"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Invalid Goal type")

	userID = "Liza"
	campaignKey = "phpab3"
	options := schema.Options{
		RevenueGoal: 0,
	}
	goalIdentifier = "rev"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, options)
	assertOutput.False(value, "Revenue Not defined")

	userID = "Liza"
	campaignKey = "php2"
	options = schema.Options{
		RevenueGoal: 10,
	}
	goalIdentifier = "custom"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, options)
	assertOutput.False(value, "No Variation in Campaign")

	options = schema.Options{
		RevenueGoal: 12,
	}
	userID = "Misty"
	campaignKey = "phpab3"
	goalIdentifier = "custom"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, options)
	assertOutput.True(value, "Variation should be defined")
}
