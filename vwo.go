package vwo

import (
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
	Launch(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage interface{})
	LaunchWithLogger(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage interface{}, logger *logger.Logger)
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

// Launch function
func (vwo *VWOInstance) Launch(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage interface{}) {
	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()

	vwo.LaunchWithLogger(isDevelopmentMode, settingsFile, storage, logs)
}

// LaunchWithLogger Function
func (vwo *VWOInstance) LaunchWithLogger(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage interface{}, logger *logger.Logger) {
	vwo.SettingsFile = settingsFile
	vwo.UserStorage = storage
	vwo.Logger = logger
	vwo.IsDevelopmentMode = isDevelopmentMode
	message := fmt.Sprintf(constants.DebugMessagesDevelopmentMode, isDevelopmentMode)
	utils.LogMessage(vwo.Logger, constants.Debug, fileVWO, message)
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
