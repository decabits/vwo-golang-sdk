package core

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/service"
	"github.com/google/logger"
	"github.com/stretchr/testify/assert"
)

// UserStorage interface for testing
type UserStorage schema.UserStorage

// UserStorageData struct for testing
type UserStorageData struct{}

// Get function is used to get the data from user storage
func (us *UserStorageData) Get(userID, campaignKey string) schema.UserData {
	return schema.UserData{
		UserID:        userID,
		CampaignKey:   campaignKey,
		VariationName: "Control",
	}
}

// Set function
func (us *UserStorageData) Set(userID, campaignKey, variationName string) {
}

// Exist function
func (us *UserStorageData) Exist() bool {
	return false
}

// GetInstance function creates and return a temporary VWO instance for testing
func GetInstance(path string) schema.VwoInstance {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.ProcessSettingsFile(path); err != nil {
		log.Println("Error Processing Settings File: ", err)
	}
	settingsFileManager.Process()
	settingsFile := settingsFileManager.GetSettingsFile()

	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()

	storage := &UserStorageData{}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settingsFile,
		UserStorage:       storage,
		Logger:            logs,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}
func TestBucketUserToVariation(t *testing.T) {
	vwoInstance := GetInstance("./testData/settings6.json")

	campaign := vwoInstance.SettingsFile.Campaigns[1]
	userID := "Linda"
	actual, _ := BucketUserToVariation(vwoInstance, userID, campaign)
	expected := campaign.Variations[2]
	assert.Equal(t, expected, actual, "Variations did not match")

	campaign = vwoInstance.SettingsFile.Campaigns[0]
	userID = "Linda"
	actual, _ = BucketUserToVariation(vwoInstance, userID, campaign)
	assert.Empty(t, actual, "Variation expected to be empty")
}

func TestGetBucketerVariation(t *testing.T) {
	vwoInstance := GetInstance("./testData/settings6.json")

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
	vwoinstance := GetInstance("./testData/settings6.json")

	userID := "James"
	campaign := vwoinstance.SettingsFile.Campaigns[1]
	actual := IsUserPartOfCampaign(vwoinstance, userID, campaign)
	assert.True(t, actual, "User should be part of the campaign")

	userID = "Christy"
	campaign = vwoinstance.SettingsFile.Campaigns[0]
	actual = IsUserPartOfCampaign(vwoinstance, userID, campaign)
	assert.False(t, actual, "User should not be part of the campaign")
}

func TestGetBucketValueForUser(t *testing.T) {
	vwoInstance := GetInstance("./testData/settings6.json")

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
