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
}
