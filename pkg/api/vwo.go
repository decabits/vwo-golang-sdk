/*
   Copyright 2019-2020 Wingify Software Pvt. Ltd.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package api

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/utils"
	"github.com/google/logger"
)

const fileVWO = "vwo.go"

// VWOInstance struct
type VWOInstance schema.VwoInstance

// VWO interface
type VWO interface {
	Launch(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage interface{}) (*VWOInstance, error)
	Activate(campaignKey, userID string, options interface{}) string
	GetFeatureVariableValue(campaignKey, variableKey, userID string, options interface{}) interface{}
	GetVariationName(campaignKey, userID string, options interface{}) string
	IsFeatureEnabled(campaignKey, userID string, options interface{}) bool
	Push(tagKey, tagValue, userID string) bool
	Track(campaignKey, userID string, goalIdentifier string, options interface{}) bool
}

// Launch function to launch SDK
func (vwo *VWOInstance) Launch(isDevelopmentMode bool, settingsFile schema.SettingsFile, storage interface{}, logs interface{}) (*VWOInstance, error) {
	/*
		Args:
			isDevelopmentMode: turn this true to stop API calls to server
			settingsFile: settings file fetched from getsettingsfile
			storage: custom storage functions
			logs: custom logger if any

		Returns:
			error: nil if SDK is launched, else error encountered
	*/

	if logs != nil {
		utils.LogMessage(logs, constants.Debug, fileVWO, constants.DebugMessageCustomLoggerFound)
	}

	if logs == nil {
		logs = logger.Init(constants.SDKName, true, false, ioutil.Discard)
		logger.SetFlags(log.LstdFlags)
		utils.LogMessage(logs, constants.Debug, fileVWO, constants.DebugMessageNoCustomLoggerFound)
		defer logger.Close()
	}

	if utils.ValidateStorage(storage) && utils.ValidateLogger(logs) {
		vwo.SettingsFile = settingsFile
		vwo.UserStorage = storage
		vwo.Logger = logs
		vwo.IsDevelopmentMode = isDevelopmentMode
		message := fmt.Sprintf(constants.DebugMessagesDevelopmentMode+constants.DebugMessagesSDKInitialized, isDevelopmentMode)
		utils.LogMessage(vwo.Logger, constants.Debug, fileVWO, message)
		return vwo, nil
	}
	return vwo, fmt.Errorf(constants.ErrorMessageInvalidLoggerStorage)
}
