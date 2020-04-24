package vwo

import (
	"io/ioutil"
	"log"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/service"
	"github.com/google/logger"
	"github.com/stretchr/testify/mock"
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

type MockRunner struct {
	mock.Mock
}

// GetInstance function creates and return a temporary VWO instance for testing
func (vwo *VWOInstance) getInstance(path string) {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.ProcessSettingsFile(path); err != nil {
		log.Println("Error Processing Settings File: ", err)
	}
	settingsFileManager.Process()
	settingsFile := settingsFileManager.GetSettingsFile()

	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	// defer logger.Close()

	storage := &UserStorageData{}

	vwo.LaunchWithLogger(true, settingsFile, storage, logs)
}

// func TestLaunch(t *testing.T) {
// 	mockRunner := MockRunner{}
// 	vwoInstance := schema.VwoInstance{
// 		SettingsFile:      settingsFile,
// 		UserStorage:       storage,
// 		Logger:            logs,
// 		IsDevelopmentMode: isDevelopmentMode,
// 	}

// 	mockRunner.On("FetchSettingsFile", "accountID", "SDKKey").Return(errors.New("there was an error"))
// 	mockRunner.AssertExpectations(t)
// }
