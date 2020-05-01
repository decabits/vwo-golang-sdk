package vwo

import (
	"github.com/decabits/vwo-golang-sdk/pkg/api"
	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/service"
	"github.com/google/logger"
)

// VWOInstance struct
type VWOInstance schema.VwoInstance

const fileVWO = "vwo.go"

func Init(settingsFile schema.SettingsFile, vwoOption ...api.VWOOption) (*api.VWOInstance, error) {
	/*
		Args:
			isDevelopmentMode: turn this true to stop API calls to server
			settingsFile: settings file fetched from getsettingsfile
			storage: custom storage functions
			logs: custom logger if any

		Returns:
			error: nil if SDK is launched, else error encountered
	*/

	vwo := &api.VWOInstance{
		SettingsFile: settingsFile,
	}
	return vwo.Launch(vwoOption...)
}

// GetSettingsFile function to fetch and parse settingsfile
func GetSettingsFile(accountID, SDKKey string) schema.SettingsFile {
	/*
		Args:
			accountID: Config account ID
			SDKKey: Config SDK Key

		Returns:
			schema.SettingsFile: settings file fetched
	*/
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.FetchSettingsFile(accountID, SDKKey); err != nil {
		logger.Warning(constants.ErrorMessageCannotProcessSettingsFile + err.Error())
	}
	settingsFileManager.Process()
	logger.Warning(fileVWO + " : " + constants.DebugMessagesSettingsFileProcessed)
	return settingsFileManager.GetSettingsFile()
}
