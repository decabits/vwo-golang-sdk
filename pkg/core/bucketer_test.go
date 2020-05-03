/*
   Copyright 2019-2020 Wingify Software Pvt. Ltd.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/logger"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/service"
	"github.com/stretchr/testify/assert"
)

// UserStorage interface
type UserStorage interface {
	Get(userID, campaignKey string) schema.UserData
	Set(userID, campaignKey, variationName string)
}

// UserStorageData struct
type UserStorageData struct{}

// BucketTestCase struct
type BucketTestCase struct {
	User        string `json:"user"`
	BucketValue int    `json:"bucket_value"`
}

// data is an example of how data is stored
var data = `{
    "CAMPAIGN_1": [{
            "UserID": "user1",
            "CampaignKey": "CAMPAIGN_1",
            "VariationName": "Control"
        },
        {
            "UserID": "user2",
            "CampaignKey": "CAMPAIGN_1",
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
	assertOutput := assert.New(t)
	vwoInstance := getInstanceWithoutStorage("./testdata/testBucket.json")

	campaign := vwoInstance.SettingsFile.Campaigns[1]
	userID := "Linda"
	expected := campaign.Variations[2]
	actual, err := BucketUserToVariation(vwoInstance, userID, campaign)
	assertOutput.Nil(err, "Variations did not match")
	assertOutput.Equal(expected, actual, "Variations did not match")

	campaign = vwoInstance.SettingsFile.Campaigns[0]
	userID = "Linda"
	actual, err = BucketUserToVariation(vwoInstance, userID, campaign)
	assertOutput.NotNil(err, "Variation expected to be empty")
	assertOutput.Empty(actual, "Variation expected to be empty")
}

func TestGetBucketerVariation(t *testing.T) {
	assertOutput := assert.New(t)
	vwoInstance := getInstanceWithoutStorage("./testdata/testBucket.json")

	variations := vwoInstance.SettingsFile.Campaigns[1].Variations
	bucketValue := 2345
	actual, err := GetBucketerVariation(vwoInstance, variations, bucketValue, "", "")
	expected := variations[0]
	assertOutput.Nil(err, "Expected Variation do not match with Actual")
	assertOutput.Equal(expected, actual, "Expected Variation do not match with Actual")

	bucketValue = 0
	actual, err = GetBucketerVariation(vwoInstance, variations, bucketValue, "", "")
	assertOutput.NotNil(err, "Variation should be empty")
	assertOutput.Empty(actual, "Variation should be empty")

	bucketValue = 12345
	actual, err = GetBucketerVariation(vwoInstance, variations, bucketValue, "", "")
	assertOutput.NotNil(err, "Variation should be empty")
	assertOutput.Empty(actual, "Variation should be empty")
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
	var settings map[string][]BucketTestCase
	data, err := ioutil.ReadFile("./testdata/bucketValueExpectations.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &settings); err != nil {
		logger.Info("Error: " + err.Error())
	}

	TestCases := settings["USER_AND_BUCKET_VALUES"]

	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()

	vwoInstance := schema.VwoInstance{
		Logger: logs,
	}

	for _, testCase := range TestCases {
		expected := testCase.BucketValue
		_, actual := GetBucketValueForUser(vwoInstance, testCase.User, 10000, 1)
		assert.Equal(t, expected, actual, "Failed for: "+ testCase.User)
	}
}

func TestHash(t *testing.T) {
	actual := hash("Robert")
	expected := uint32(1150261924)
	assert.Equal(t, expected, actual, "Hash values do not match")

	actual = hash("12345")
	expected = uint32(1377935000)
	assert.Equal(t, expected, actual, "Hash values do not match")
}
