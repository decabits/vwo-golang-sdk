package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchSettingsFile(t *testing.T) {
	settingsFileManager := SettingsFileManager{}
	err := settingsFileManager.FetchSettingsFile("accountID", "SDKKey")
	assert.Error(t, err, "No settingsFile processed")

	err = settingsFileManager.FetchSettingsFile("", "")
	assert.Error(t, err, "No settingsFile processed")

	err = settingsFileManager.FetchSettingsFile("accountID", "")
	assert.Error(t, err, "No settingsFile processed")
}

func TestProcessSettingsFile(t *testing.T) {
	settingsFileManager := SettingsFileManager{}
	err := settingsFileManager.ProcessSettingsFile("./testData/settingsFile2.json")
	assert.NoError(t, err, "No settingsFile processed")

	settingsFileManager.Process()
	settingsFile := settingsFileManager.GetSettingsFile()
	assert.NotEmpty(t, settingsFile, "No settingsFile processed")

	err = settingsFileManager.ProcessSettingsFile("./testData/settingsFile.json")
	assert.Error(t, err, "No settingsFile processed")

	err = settingsFileManager.ProcessSettingsFile("./testData/settingsFile3.json")
	assert.Error(t, err, "No settingsFile processed")
}
