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

type VWOOption func(*VWOInstance)

// Launch function to launch SDK
func (vwo VWOInstance) Launch(vwoOption ...VWOOption) (*VWOInstance, error) {
	/*
		Args:
			isDevelopmentMode: turn this true to stop API calls to server
			settingsFile: settings file fetched from getsettingsfile
			storage: custom storage functions
			logs: custom logger if any

		Returns:
			error: nil if SDK is launched, else error encountered
	*/

	for _, option := range vwoOption {
		option(&vwo)
	}

	if vwo.Logger != nil {
		utils.LogMessage(vwo.Logger, constants.Debug, fileVWO, constants.DebugMessageCustomLoggerFound)
	}

	if vwo.Logger == nil {
		logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
		logger.SetFlags(log.LstdFlags)
		utils.LogMessage(logs, constants.Debug, fileVWO, constants.DebugMessageNoCustomLoggerFound)
		vwo.Logger = logs
		defer logger.Close()
	}

	message := fmt.Sprintf(constants.DebugMessagesDevelopmentMode+constants.DebugMessagesSDKInitialized, vwo.IsDevelopmentMode)
	utils.LogMessage(vwo.Logger, constants.Debug, fileVWO, message)

	if !utils.ValidateStorage(vwo.UserStorage) || !utils.ValidateLogger(vwo.Logger) {
		return &vwo, fmt.Errorf(constants.ErrorMessageInvalidLoggerStorage)
	}
	return &vwo, nil
}

// WithStorage sets user storage
func WithStorage(storage interface{}) VWOOption {
	return func(vwo *VWOInstance) {
		vwo.UserStorage = storage
	}
}

// WithLogger sets user storage
func WithLogger(logger interface{}) VWOOption {
	return func(vwo *VWOInstance) {
		vwo.Logger = logger
	}
}

// WithDevelopmentMode sets user storage
func WithDevelopmentMode() VWOOption {
	return func(vwo *VWOInstance) {
		vwo.IsDevelopmentMode = true
	}
}
