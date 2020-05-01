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

package vwo

import (
	"github.com/decabits/vwo-golang-sdk/pkg/api"
	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/logger"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/service"
)

// VWOInstance struct to store params
type VWOInstance schema.VwoInstance

const fileVWO = "vwo.go"

// Init function to intialise sdk
func Init(settingsFile schema.SettingsFile, vwoOption ...api.VWOOption) (*api.VWOInstance, error) {
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
