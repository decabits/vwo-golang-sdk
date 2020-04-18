package api

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
func TestActivate(t *testing.T) {
	vwoInstance := GetInstance("./testData/testdata.json")

	userID := "Varun"
	campaignKey := "notPresent"
	value := Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab1"
	value = Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "php1"
	value = Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "phpab2"
	value = Activate(vwoInstance, campaignKey, userID)
	assert.Empty(t, value, "No Variation in Campaign")

	userID = "Liza"
	campaignKey = "phpab3"
	actual := Activate(vwoInstance, campaignKey, userID)
	assert.NotEmpty(t, actual, "Variation should be found")
}
