package utils

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/stretchr/testify/assert"
)

func TestCreateImpressionForPush(t *testing.T) {
	vwoInstance := GetInstance()
	URL := "https://dev.visualwebsiteoptimizer.com/server-side/push"

	userID := "Chris"
	tagValue := ""
	tagKey := ""
	DemoImpression := CreateImpressionForPush(vwoInstance, tagKey, tagValue, userID)
	assert.Equal(t, userID, DemoImpression.UID, "Non Matching UIDs")
	assert.Equal(t, vwoInstance.SettingsFile.AccountID, DemoImpression.AccountID, "Non Matching Account IDs")
	assert.Equal(t, URL, DemoImpression.URL, "Non Matching URLs")
	assert.Equal(t, "=", DemoImpression.U, "Non Matching Parameters")

	userID = "Lizzie"
	tagValue = "testVal"
	tagKey = "testKey"
	DemoImpression = CreateImpressionForPush(vwoInstance, tagKey, tagValue, userID)
	assert.Equal(t, userID, DemoImpression.UID, "Non Matching UIDs")
	assert.Equal(t, vwoInstance.SettingsFile.AccountID, DemoImpression.AccountID, "Non Matching Account IDs")
	assert.Equal(t, URL, DemoImpression.URL, "Non Matching URLs")
	assert.Equal(t, "testKey=testVal", DemoImpression.U, "Non Matching Parameters")
}

func TestCreateImpressionTrackingUser(t *testing.T) {
	vwoInstance := GetInstance()
	variationID := 1
	campaignID := 283
	userID := "Chris"
	URL := "https://dev.visualwebsiteoptimizer.com/server-side/track-user"

	DemoImpression := CreateImpressionTrackingUser(vwoInstance, campaignID, variationID, userID)
	assert.Equal(t, userID, DemoImpression.UID, "Non Matching UIDs")
	assert.Equal(t, variationID, DemoImpression.Combination, "Non Matching VariationIDs")
	assert.Equal(t, campaignID, DemoImpression.ExperimentID, "Non Matching CampaignIDs")
	assert.Equal(t, vwoInstance.SettingsFile.AccountID, DemoImpression.AccountID, "Non Matching Account IDs")
	assert.Equal(t, URL, DemoImpression.URL, "Non Matching URLs")

}
func TestCreateImpressionTrackingGoal(t *testing.T) {
	vwoInstance := GetInstance()
	variationID := 1
	campaignID := 283
	goalID := 281
	revenueGoal := 5
	userID := "Chris"
	URL := "https://dev.visualwebsiteoptimizer.com/server-side/track-goal"

	DemoImpression := CreateImpressionTrackingGoal(vwoInstance, variationID, userID, campaignID, goalID, revenueGoal)
	assert.Equal(t, userID, DemoImpression.UID, "Non Matching UIDs")
	assert.Equal(t, variationID, DemoImpression.Combination, "Non Matching VariationIDs")
	assert.Equal(t, goalID, DemoImpression.GoalID, "Non Matching GoalIDs")
	assert.Equal(t, campaignID, DemoImpression.ExperimentID, "Non Matching CampaignIDs")
	assert.Equal(t, revenueGoal, DemoImpression.R, "Non Matching Revenues")
	assert.Equal(t, vwoInstance.SettingsFile.AccountID, DemoImpression.AccountID, "Non Matching Account IDs")
	assert.Equal(t, URL, DemoImpression.URL, "Non Matching URLs")
	var temp1 string
	assert.IsType(t, temp1, DemoImpression.SID, "Incorrect SID type")
	var temp2 float32
	assert.IsType(t, temp2, DemoImpression.Random, "Incorrect Random type")
}

func TestGetCommonProperties(t *testing.T) {
	userID := "Gimmy"
	vwoInstance := GetInstance()
	impression := getCommonProperties(vwoInstance, userID)

	assert.Equal(t, userID, impression.UID, "Non Matching UIDs")
	assert.Equal(t, vwoInstance.SettingsFile.AccountID, impression.AccountID, "Non Matching AccountIDs")
	assert.Equal(t, constants.Platform, impression.Ap, "Non Matching Platforms")
	assert.Equal(t, constants.SDKName, impression.Sdk, "Non Matching SDK")
	assert.Equal(t, constants.SDKVersion, impression.SdkV, "Non Matching SDK Version")
	var temp1 float32
	assert.IsType(t, temp1, impression.Random, "Random number should be of type float32")
	var temp2 string
	assert.IsType(t, temp2, impression.SID, "Random number should be of type float32")

}
