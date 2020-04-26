package vwo

import (
	"log"
	"testing"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/service"
	"github.com/stretchr/testify/assert"
)

// GetInstance function creates and return a temporary VWO instance for testing
func (vwo *VWOInstance) getInstance(path string) error {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.ProcessSettingsFile(path); err != nil {
		log.Println("Error Processing Settings File: ", err)
	}
	settingsFileManager.Process()
	settingsFile := settingsFileManager.GetSettingsFile()

	if err := vwo.Launch(true, settingsFile, nil, nil); err != nil {
		return err
	}
	return nil
}

type WUserStorage interface {
	Getter(userID, campaignKey string) schema.UserData
	Setter(userID, campaignKey, variationName string)
}
type WUserStorageData struct{}

func (us *WUserStorageData) Getter(userID, campaignKey string) schema.UserData {
	return schema.UserData{}
}
func (us *WUserStorageData) Setter(userID, campaignKey, variationName string) {}

func TestLaunch(t *testing.T) {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.ProcessSettingsFile("./testdata/testdata.json"); err != nil {
		log.Println("Error Processing Settings File: ", err)
	}
	settingsFileManager.Process()
	settingsFile := settingsFileManager.GetSettingsFile()

	vwoInstance := VWOInstance{}
	storage := &WUserStorageData{}
	err := vwoInstance.Launch(true, settingsFile, storage, nil)
	assert.NotNil(t, err)

	vwoInstance = VWOInstance{}
	err = vwoInstance.Launch(true, settingsFile, nil, nil)
	assert.Nil(t, err)
}
