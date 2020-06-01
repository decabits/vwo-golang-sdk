/*
   Copyright 2020 Wingify Software Pvt. Ltd.

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
	"github.com/decabits/vwo-golang-sdk/pkg/logger"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsFeatureEnabled(t *testing.T) {
	assertOutput := assert.New(t)

	var userExpectation map[string][]TestCase
	data, err := ioutil.ReadFile("../testdata/userExpectations1.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &userExpectation); err != nil {
		logger.Info("Error: " + err.Error())
	}

	var settingsFiles map[string]schema.SettingsFile
	data, err = ioutil.ReadFile("../testdata/settings.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &settingsFiles); err != nil {
		logger.Info("Error: " + err.Error())
	}

	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()

	instance := VWOInstance{}
	instance.Logger = logs

	for settingsFileName, settingsFile := range settingsFiles {
		vwoInstance := schema.VwoInstance{
			Logger: logs,
		}
		settingsFile.Campaigns[0].Variations = utils.GetVariationAllocationRanges(vwoInstance, settingsFile.Campaigns[0].Variations)

		instance.SettingsFile = settingsFile

		testCases := userExpectation[settingsFileName]
		for i := range testCases {
			actual := instance.IsFeatureEnabled(settingsFile.Campaigns[0].Key, testCases[i].User, nil)
			expected := testCases[i].Variation
			if settingsFileName != "T_75_W_10_TIMES_10" && settingsFileName != "AB_T_100_W_33_33_33" && settingsFileName != "AB_T_50_W_50_50" && settingsFileName != "AB_T_100_W_0_100" && settingsFileName != "AB_T_20_W_10_90" {
				assertOutput.Equal(expected != "", actual, settingsFileName+" "+testCases[i].User)
			} else {
				assertOutput.Equal(expected == "", actual, settingsFileName+" "+testCases[i].User)
			}
		}
	}

	// vwoInstance, err := getInstance("./testdata/testIsFeatureEnabled.json")
	// assertOutput.Nil(err, "error fetching instance")

	// userID := ""
	// campaignKey := ""
	// value := vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	// assertOutput.False(value, "Invalid params")

	// userID = "USER_1"
	// campaignKey = "notPresent"
	// value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	// assertOutput.False(value, "Campaign does not exist")

	// userID = "USER_1"
	// campaignKey = "CAMPAIGN_1"
	// value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	// assertOutput.False(value, "Campaign Not running")

	// userID = "USER_3"
	// campaignKey = "CAMPAIGN_8"
	// value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	// assertOutput.False(value, "Campaign Not Valid")

	// userID = "Robbie"
	// campaignKey = "CAMPAIGN_2"
	// value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	// assertOutput.False(value, "No Variation from campaign Not alloted")

	// userID = "USER_8"
	// campaignKey = "CAMPAIGN_3"
	// value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	// assertOutput.True(value, "Campaign Feature Rollout")

	// userID = "USER_2"
	// campaignKey = "CAMPAIGN_4"
	// value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	// assertOutput.True(value, "Campaign Feature Test")

	// userID = "USER_8"
	// campaignKey = "CAMPAIGN_4"
	// value = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)
	// assertOutput.False(value, "Campaign Feature Rollout")
}
