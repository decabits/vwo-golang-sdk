package testdata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/logger"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
)

// UserStorage interface
type UserStorage interface {
	Get(userID, campaignKey string) schema.UserData
	Set(userID, campaignKey, variationName, GoalIdentifier string)
}

// IncorrectUserStorage interface
type IncorrectUserStorage interface {
	Get(userID, campaignKey string) schema.UserData
	IncorrectSet(campaignKey, variationName string)
}

// UserStorageData struct
type UserStorageData struct{}

// IncorrectUserStorageData struct
type IncorrectUserStorageData struct{}

// data is an example of how data is stored
var data = `{
    "AB_T_50_W_50_50": [{
			"UserID": "DummyUser",
			"CampaignKey": "AB_T_50_W_50_50",
			"VariationName": "DummyVariation",
			"GoalIdentifier": "DummyGoal"
		},
		{
			"UserID": "TempUser",
			"CampaignKey": "AB_T_50_W_50_50",
			"VariationName": "Control",
			"GoalIdentifier": "DummyGoal"
		}],

		"AB_T_100_W_33_33_33": [{
			"UserID": "DummyUser",
			"CampaignKey": "AB_T_100_W_33_33_33",
			"VariationName": "Control",
			"GoalIdentifier": "DummyGoal_vwo_GOAL_2"
		},
		{
			"UserID": "TempUser",
			"CampaignKey": "AB_T_100_W_33_33_33",
			"VariationName": "Control",
			"GoalIdentifier": "DummyGoal_vwo_GOAL_1"
		}]
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

// Get function is used to get the data from user storage
func (us *IncorrectUserStorageData) Get(userID, campaignKey string) schema.UserData {
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
func (us *UserStorageData) Set(userID, campaignKey, variationName, GoalIdentifier string) {
}

// IncorrectSet function
func (us *UserStorageData) IncorrectSet(campaignKey, variationName string) {
}

//GetInstanceWithIncorrectStorage function
func GetInstanceWithIncorrectStorage(SettingsFileName string) schema.VwoInstance {
	logger := logger.Init(constants.SDKName, false, false, ioutil.Discard)
	defer logger.Close()

	var settingsFiles map[string]schema.SettingsFile
	data, err := ioutil.ReadFile("../testdata/settings.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &settingsFiles); err != nil {
		logger.Info("Error: " + err.Error())
	}

	settings := settingsFiles[SettingsFileName]

	storage := &IncorrectUserStorageData{}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settings,
		UserStorage:       storage,
		Logger:            logger,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}

//GetInstanceWithStorage function
func GetInstanceWithStorage(SettingsFileName string) schema.VwoInstance {
	logger := logger.Init(constants.SDKName, false, false, ioutil.Discard)
	defer logger.Close()

	var settingsFiles map[string]schema.SettingsFile
	data, err := ioutil.ReadFile("../testdata/settings.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &settingsFiles); err != nil {
		logger.Info("Error: " + err.Error())
	}

	settings := settingsFiles[SettingsFileName]

	storage := &UserStorageData{}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settings,
		UserStorage:       storage,
		Logger:            logger,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}

//GetInstanceWithCustomSettings function
func GetInstanceWithCustomSettings(SettingsFileName string) schema.VwoInstance {
	logger := logger.Init(constants.SDKName, false, false, ioutil.Discard)
	defer logger.Close()

	var settingsFiles map[string]schema.SettingsFile
	data, err := ioutil.ReadFile("../testdata/custom_settings.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &settingsFiles); err != nil {
		logger.Info("Error: " + err.Error())
	}

	settings := settingsFiles[SettingsFileName]

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settings,
		UserStorage:       nil,
		Logger:            logger,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}

//GetInstanceWithSettings function
func GetInstanceWithSettings(SettingsFileName string) schema.VwoInstance {
	logger := logger.Init(constants.SDKName, false, false, ioutil.Discard)
	defer logger.Close()

	var settingsFiles map[string]schema.SettingsFile
	data, err := ioutil.ReadFile("../testdata/settings.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &settingsFiles); err != nil {
		logger.Info("Error: " + err.Error())
	}

	settings := settingsFiles[SettingsFileName]

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settings,
		UserStorage:       nil,
		Logger:            logger,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}
