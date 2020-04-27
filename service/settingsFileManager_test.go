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
