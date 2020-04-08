package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateImpressionTrackingUser(t *testing.T) {
	vwoInstance := GetTempInstance()
	variationID := 1
	campaignID := 283
	goalID := 281
	revenueGoal := 5
	userID := "Chris"

	DemoImpression := CreateImpressionTrackingGoal(vwoInstance, variationID, userID, campaignID, goalID, revenueGoal)
	assert.Equal(t, userID, DemoImpression.UID, "Non Matching UIDs")
	assert.Equal(t, variationID, DemoImpression.Combination, "Non Matching VariationIDs")
	assert.Equal(t, goalID, DemoImpression.GoalID, "Non Matching GoalIDs")
	assert.Equal(t, campaignID, DemoImpression.ExperimentID, "Non Matching CampaignIDs")
	assert.Equal(t, revenueGoal, DemoImpression.R, "Non Matching Revenues")
	assert.Equal(t, vwoInstance.SettingsFile.AccountID, DemoImpression.AccountID, "Non Matching Account IDs")
	var temp1 string
	assert.IsType(t, temp1, DemoImpression.SID, "Incorrect SID type")
	var temp2 float32
	assert.IsType(t, temp2, DemoImpression.Random, "Incorrect Random type")
}
