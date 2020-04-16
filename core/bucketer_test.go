package core

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestBucketUserToVariation(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")

	campaign := vwoInstance.SettingsFile.Campaigns[2]
	userID := "Linda"
	actual, _ := BucketUserToVariation(vwoInstance, userID, campaign)
	expected := campaign.Variations[2]
	assert.Equal(t, expected, actual, "Variations did not match")

	campaign = vwoInstance.SettingsFile.Campaigns[3]
	userID = "Linda"
	actual, _ = BucketUserToVariation(vwoInstance, userID, campaign)
	assert.Empty(t, actual, "Variation expected to be empty")	
}

func TestGetBucketerVariation(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")

	variations := vwoInstance.SettingsFile.Campaigns[1].Variations
	bucketValue := 2345
	actual, _ := GetBucketerVariation(variations, bucketValue)
	expected := variations[0]
	assert.Equal(t, expected, actual, "Expected Variation do not match with Actual")

	bucketValue = 0
	actual, _ = GetBucketerVariation(variations, bucketValue)
	assert.Empty(t, actual, "Variation should be empty")

	bucketValue = 12345
	actual, _ = GetBucketerVariation(variations, bucketValue)
	assert.Empty(t, actual, "Variation should be empty")
}

func TestIsUserPartOfCampaign(t *testing.T) {
	vwoinstance := utils.GetInstance("../settingsFile.json")

	userID := "James"
	campaign := vwoinstance.SettingsFile.Campaigns[0]
	actual := IsUserPartOfCampaign(vwoinstance, userID, campaign)
	assert.True(t, actual, "User should be part of the campaign")

	userID = "Christy"
	campaign = vwoinstance.SettingsFile.Campaigns[3]
	actual = IsUserPartOfCampaign(vwoinstance, userID, campaign)
	assert.False(t, actual, "User should not be part of the campaign")
}

func TestGetBucketValueForUser(t *testing.T) {
	vwoInstance := utils.GetInstance("../settingsFile.json")

	userID := "Chris"
	actual := GetBucketValueForUser(vwoInstance, userID, constants.MaxTrafficPercent, 1)
	expected := 93
	assert.Equal(t, expected, actual, "Bucket Values do not match")

	userID = "Chris"
	actual = GetBucketValueForUser(vwoInstance, userID, constants.MaxTrafficPercent, 0.5)
	expected = 46
	assert.Equal(t, expected, actual, "Bucket Values do not match")

	userID = "Liza"
	actual = GetBucketValueForUser(vwoInstance, userID, constants.MaxTrafficValue, 1)
	expected = 3379
	assert.Equal(t, expected, actual, "Bucket Values do not match")

	userID = "Gimmy"
	actual = GetBucketValueForUser(vwoInstance, userID, constants.MaxTrafficValue, 1)
	expected = 9572
	assert.Equal(t, expected, actual, "Bucket Values do not match")
}

func TestHash(t *testing.T) {
	actual := hash("Robert")
	expected := uint32(1150261924)
	assert.Equal(t, expected, actual, "Hash values do not match")

	actual = hash("12345")
	expected = uint32(1377935000)
	assert.Equal(t, expected, actual, "Hash values do not match")
}
