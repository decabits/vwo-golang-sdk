package vwo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/service"
	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/google/logger"
)

const fileVWO = "vwo.go"

// VWOInstance struct
type VWOInstance schema.VwoInstance

// VWO interface
type VWO interface {
	Launch(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage interface{}) error
	Activate(campaignKey, userID string, options interface{}) string
	GetFeatureVariableValue(campaignKey, variableKey, userID string, options interface{}) interface{}
	GetVariationName(campaignKey, userID string, options interface{}) string
	IsFeatureEnabled(campaignKey, userID string, options interface{}) bool
	Push(tagKey, tagValue, userID string) bool
	Track(campaignKey, userID string, goalIdentifier string, options interface{}) bool
}

// Launch function to launch sdk
func (vwo *VWOInstance) Launch(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage interface{}, logs interface{}) error {
	/*
		Args:
			isDevelopmentMode: turn this true to stop API calls to server
			settingsFile: settings file fetched from getsettingsfile
			storage: custom storage functions
			logs: custom logger if any
	*/
	if logs == nil {
		logs = logger.Init(constants.SDKName, true, false, ioutil.Discard)
		logger.SetFlags(log.LstdFlags)
		defer logger.Close()
	}

	if utils.ValidateStorage(storage) && utils.ValidateLogger(logs) {
		vwo.SettingsFile = settingsFile
		vwo.UserStorage = storage
		vwo.Logger = logs
		vwo.IsDevelopmentMode = isDevelopmentMode
		message := fmt.Sprintf(constants.DebugMessagesDevelopmentMode, isDevelopmentMode)
		utils.LogMessage(vwo.Logger, constants.Debug, fileVWO, message)
		return nil
	}
	return errors.New("Invalid storage object/Logger given. Refer documentation on how to pass custom storage.")
}

// GetSettingsFile function to fetch and parse settingsfile
func GetSettingsFile(accountID, SDKKey string) schema.SettingsFile {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.FetchSettingsFile(accountID, SDKKey); err != nil {
		logger.Warning("Error Processing Settings File: " + err.Error())
	}
	settingsFileManager.Process()
	logger.Warning(fileVWO + " : " + constants.DebugMessagesSettingsFileProcessed)
	return settingsFileManager.GetSettingsFile()
}
