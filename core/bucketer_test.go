package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/service"
	"github.com/google/logger"
	"github.com/stretchr/testify/assert"
)

// UserStorage interface
type UserStorage interface {
	Get(userID, campaignKey string) schema.UserData
	Set(userID, campaignKey, variationName string)
}

// UserStorageData struct
type UserStorageData struct{}

// data is an example of how data is stored
var data = `{
    "php1": [{
            "UserID": "user1",
            "CampaignKey": "php1",
            "VariationName": "Control"
        },
        {
            "UserID": "user2",
            "CampaignKey": "php1",
            "VariationName": "Variation-1"
        }
    ]
}`

// Get function is used to get the data from user storage
func (us *UserStorageData) Get(userID, campaignKey string) schema.UserData {
	var userDatas map[string][]schema.UserData
	// Conect your database here to fetch the current data
	// Uncomment the below part to user JSON as data base
	if err := json.Unmarshal([]byte(data), &userDatas); err != nil {
		fmt.Print("Could not unmarshall")
	}
	if len(userDatas) == 0 {
		return schema.UserData{}
	}
	userData, ok := userDatas[campaignKey]
	if ok {
		for _, userdata := range userData {
			if userdata.UserID == userID {
				return userdata
			}
		}
	}
	return schema.UserData{}
}

// Set function
func (us *UserStorageData) Set(userID, campaignKey, variationName string) {
}

func getInstanceWithStorage(path string) schema.VwoInstance {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.ProcessSettingsFile(path); err != nil {
		log.Println("Error Processing Settings File: ", err)
	}
	settingsFileManager.Process()
	settingsFile := settingsFileManager.GetSettingsFile()

	logs := logger.Init(constants.SDKName, false, false, ioutil.Discard)
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

func getInstanceWithoutStorage(path string) schema.VwoInstance {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.ProcessSettingsFile(path); err != nil {
		log.Println("Error Processing Settings File: ", err)
	}
	settingsFileManager.Process()
	settingsFile := settingsFileManager.GetSettingsFile()

	logs := logger.Init(constants.SDKName, false, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settingsFile,
		UserStorage:       nil,
		Logger:            logs,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}

func TestBucketUserToVariation(t *testing.T) {
	vwoInstance := getInstanceWithoutStorage("./testdata/testBucket.json")

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
	vwoInstance := getInstanceWithoutStorage("./testdata/testBucket.json")

	variations := vwoInstance.SettingsFile.Campaigns[1].Variations
	bucketValue := 2345
	actual, _ := GetBucketerVariation(vwoInstance, variations, bucketValue, "", "")
	expected := variations[0]
	assert.Equal(t, expected, actual, "Expected Variation do not match with Actual")

	bucketValue = 0
	actual, _ = GetBucketerVariation(vwoInstance, variations, bucketValue, "", "")
	assert.Empty(t, actual, "Variation should be empty")

	bucketValue = 12345
	actual, _ = GetBucketerVariation(vwoInstance, variations, bucketValue, "", "")
	assert.Empty(t, actual, "Variation should be empty")
}

func TestIsUserPartOfCampaign(t *testing.T) {
	vwoInstance := getInstanceWithoutStorage("./testdata/testBucket.json")

	userID := "James"
	campaign := vwoInstance.SettingsFile.Campaigns[1]
	actual := IsUserPartOfCampaign(vwoInstance, userID, campaign)
	assert.True(t, actual, "User should be part of the campaign")

	userID = "Christy"
	campaign = vwoInstance.SettingsFile.Campaigns[0]
	actual = IsUserPartOfCampaign(vwoInstance, userID, campaign)
	assert.False(t, actual, "User should not be part of the campaign")
}

func TestGetBucketValueForUser(t *testing.T) {
	vwoInstance := getInstanceWithoutStorage("./testdata/testBucket.json")

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
