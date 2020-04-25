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
	LaunchWithLogger(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage interface{}, logger *logger.Logger) error
	Activate(campaignKey, userID string) string
	ActivateWithOptions(campaignKey, userID string, options schema.Options) string
	GetFeatureVariableValue(campaignKey, variableKey, userID string) interface{}
	GetFeatureVariableValueWithOptions(campaignKey, variableKey, userID string, options schema.Options) interface{}
	GetVariationName(campaignKey, userID string) string
	GetVariationNameWithOptions(campaignKey, userID string, options schema.Options) string
	IsFeatureEnabled(campaignKey, userID string) bool
	IsFeatureEnabledWithOptions(campaignKey, userID string, options schema.Options) bool
	Push(tagKey, tagValue, userID string) bool
	Track(campaignKey, userID string, goalIdentifier string) bool
	TrackWithOptions(campaignKey, userID string, goalIdentifier string, options schema.Options) bool
}

// Launch function to launch sdk
func (vwo *VWOInstance) Launch(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage interface{}) error {
	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()

	if utils.ValidateStorage(storage) {
		return vwo.LaunchWithLogger(isDevelopmentMode, settingsFile, storage, logs)
	}
	return errors.New("Invalid storage object given. Refer documentation on how to pass custom storage.")
}

// LaunchWithLogger function to launch sdk with custom logger
func (vwo *VWOInstance) LaunchWithLogger(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage interface{}, logs interface{}) error {
	if utils.ValidateLogger(logs) {
		vwo.SettingsFile = settingsFile
		vwo.UserStorage = storage
		vwo.Logger = logs
		vwo.IsDevelopmentMode = isDevelopmentMode
		message := fmt.Sprintf(constants.DebugMessagesDevelopmentMode, isDevelopmentMode)
		utils.LogMessage(vwo.Logger, constants.Debug, fileVWO, message)
		return nil
	}
	return errors.New("Invalid storage object given. Refer documentation on how to pass custom storage.")
}

// GetSettingsFile function to fetch settingsfile
func GetSettingsFile(accountID, SDKKey string) schema.SettingsFile {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.FetchSettingsFile(accountID, SDKKey); err != nil {
		logger.Warning("Error Processing Settings File: " + err.Error())
	}
	settingsFileManager.Process()
	logger.Warning(fileVWO + " : " + constants.DebugMessagesSettingsFileProcessed)
	return settingsFileManager.GetSettingsFile()
}
