package vwo

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/service"
	"github.com/google/logger"
)

// VWOInstance struct
type VWOInstance schema.VwoInstance

// VWO interface
type VWO interface {
	Launch(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage schema.UserStorage)
	LaunchWithLogger(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage schema.UserStorage, logger *logger.Logger)
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
func (vwo *VWOInstance) Launch(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage schema.UserStorage) {
	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()

	vwo.LaunchWithLogger(isDevelopmentMode, settingsFile, storage, logs)
}

// LaunchWithLogger Function
func (vwo *VWOInstance) LaunchWithLogger(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage schema.UserStorage, logger *logger.Logger) {
	vwo.SettingsFile = settingsFile
	vwo.UserStorage = storage
	vwo.Logger = logger
	vwo.IsDevelopmentMode = isDevelopmentMode
}

// GetSettingsFile function to fetch settingsfile
func GetSettingsFile(accountID, SDKKey string) (schema.SettingsFile, error) {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.FetchSettingsFile(accountID, SDKKey); err != nil {
		return schema.SettingsFile{}, fmt.Errorf("Error Processing Settings File: %v", err)
	}
	settingsFileManager.Process()
	return settingsFileManager.GetSettingsFile(), nil
}
