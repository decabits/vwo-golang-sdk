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
	// "encoding/json"
	// "io/ioutil"
	// "log"
	"testing"

	// "github.com/decabits/vwo-golang-sdk/pkg/constants"
	// "github.com/decabits/vwo-golang-sdk/pkg/logger"
	// "github.com/decabits/vwo-golang-sdk/pkg/schema"
	// "github.com/decabits/vwo-golang-sdk/pkg/testdata"
	// "github.com/decabits/vwo-golang-sdk/pkg/utils"
	// "github.com/stretchr/testify/assert"
)

type FeatureTestCase struct {

}
func TestGetFeatureVariableValue(t *testing.T) {
	// assertOutput := assert.New(t)

	// var userExpectation map[string][]TestCase
	// data, err := ioutil.ReadFile("../testdata/userExpectations2.json")
	// if err != nil {
	// 	logger.Info("Error: " + err.Error())
	// }

	// if err = json.Unmarshal(data, &userExpectation); err != nil {
	// 	logger.Info("Error: " + err.Error())
	// }

	// var settingsFiles map[string]schema.SettingsFile
	// data, err = ioutil.ReadFile("../testdata/settings.json")
	// if err != nil {
	// 	logger.Info("Error: " + err.Error())
	// }

	// if err = json.Unmarshal(data, &settingsFiles); err != nil {
	// 	logger.Info("Error: " + err.Error())
	// }

	// logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	// logger.SetFlags(log.LstdFlags)
	// defer logger.Close()

	// instance := VWOInstance{}
	// instance.SettingsFile = schema.SettingsFile{}
	// instance.Logger = logs

	// for settingsFileName, settingsFile := range settingsFiles {
	// 	vwoInstance := schema.VwoInstance{
	// 		Logger: logs,
	// 	}
	// 	settingsFile.Campaigns[0].Variations = utils.GetVariationAllocationRanges(vwoInstance, settingsFile.Campaigns[0].Variations)

	// 	instance.SettingsFile = settingsFile

	// 	testCases := userExpectation[settingsFileName]
	// 	for i := range testCases {
	// 		actual := instance.GetFeatureVariableValue(settingsFile.Campaigns[0].Key, variableKey, userID, nil)(settingsFile.Campaigns[0].Key, testCases[i].User, nil)
	// 		expected := testCases[i].Variation
	// 		assertOutput.Equal(expected, actual, settingsFileName+" "+testCases[i].User)
	// 	}
	// }
	// vwoInstance, err := getInstance("./testdata/testGetFeatureVariableValue.json")
	// assertOutput.Nil(err, "error fetching instance")

	// userID := ""
	// campaignKey := ""
	// variableKey := ""
	// value := vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	// assertOutput.Nil(value, "Invalid params")

	// userID = "USER_1"
	// campaignKey = "notPresent"
	// variableKey = "float2"
	// value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	// assertOutput.Nil(value, "Campaign does not exist")

	// userID = "USER_1"
	// campaignKey = "CAMPAIGN_1"
	// variableKey = "float2"
	// value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	// assertOutput.Nil(value, "Campaign Not running")

	// userID = "USER_3"
	// campaignKey = "CAMPAIGN_8"
	// variableKey = "float2"
	// value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	// assertOutput.Nil(value, "Campaign Not Valid")

	// userID = "USER_3"
	// campaignKey = "CAMPAIGN_2"
	// variableKey = "float2"
	// value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	// assertOutput.Nil(value, "Variation Not alloted as none exist")

	// userID = "USER_1"
	// campaignKey = "CAMPAIGN_3"
	// variableKey = "string1"
	// value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	// assertOutput.Nil(value, "No variable with name found")

	// userID = "USER_2"
	// campaignKey = "CAMPAIGN_3"
	// variableKey = "float2"
	// actual := vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	// expected := 10.67
	// assertOutput.Equal(expected, actual, "Value mismatch for variable")

	// userID = "USER_2"
	// campaignKey = "CAMPAIGN_4"
	// variableKey = "bool1"
	// value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	// expected1 := true
	// assertOutput.Equal(expected1, value, "Value mismatch for variable")

	// userID = "USER_2"
	// campaignKey = "CAMPAIGN_4"
	// variableKey = "int1"
	// value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	// expected2 := 301
	// assertOutput.Equal(float64(expected2), value, "Value mismatch for variable")

	// userID = "USER_2"
	// campaignKey = "CAMPAIGN_4"
	// variableKey = "string2"
	// value = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)
	// expected4 := "abcd"
	// assertOutput.Equal(expected4, value, "Value mismatch for variable")
}
