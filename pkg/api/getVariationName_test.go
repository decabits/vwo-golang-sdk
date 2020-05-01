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
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/utils"
	"github.com/google/logger"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	User      string `json:"user"`
	Variation string `json:"variation"`
}

func TestGetVariationName(t *testing.T) {
	assertOutput := assert.New(t)

	var userExpectation map[string][]TestCase
	data, err := ioutil.ReadFile("./testdata/userExpectations1.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &userExpectation); err != nil {
		logger.Info("Error: " + err.Error())
	}

	var settingsFiles map[string]schema.SettingsFile
	data, err = ioutil.ReadFile("./testdata/settings.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &settingsFiles); err != nil {
		logger.Info("Error: " + err.Error())
	}

	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()

	for settingsFileName, settingsFile := range settingsFiles {
		vwoInstance := schema.VwoInstance{
			Logger: logs,
		}
		settingsFile.Campaigns[0].Variations = utils.GetVariationAllocationRanges(vwoInstance, settingsFile.Campaigns[0].Variations)

		instance := VWOInstance{}
		instance.SettingsFile = schema.SettingsFile{}
		instance.SettingsFile = settingsFile
		instance.Logger = logs

		testCases := userExpectation[settingsFileName]
		for i := range testCases {
			actual := instance.GetVariationName(settingsFile.Campaigns[0].Key, testCases[i].User, nil)
			expected := testCases[i].Variation
			assert.Equal(t, expected, actual, settingsFileName+" "+testCases[i].User)
		}
	}

	// CORNER CASES
	vwoInstance, err := getInstance("./testdata/testdata.json")
	assertOutput.Nil(err, "error fetching instance")

	userID := ""
	campaignKey := ""
	value := vwoInstance.GetVariationName(campaignKey, userID, nil)
	assertOutput.Empty(value, "Invalid params")

	userID = "USER_1"
	campaignKey = "notPresent"
	value = vwoInstance.GetVariationName(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign does not exist")

	userID = "USER_1"
	campaignKey = "CAMPAIGN_8"
	value = vwoInstance.GetVariationName(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign Not running")

	userID = "USER_3"
	campaignKey = "CAMPAIGN_1"
	value = vwoInstance.GetVariationName(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign Not Valid")

	userID = "USER_3"
	campaignKey = "CAMPAIGN_9"
	value = vwoInstance.GetVariationName(campaignKey, userID, nil)
	assertOutput.Empty(value, "Variation Not found")
}
