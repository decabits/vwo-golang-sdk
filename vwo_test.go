package vwo

import (
	"log"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/service"
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

	storage := &UserStorageData{}

	vwo.Launch(true, settingsFile, storage)
}
