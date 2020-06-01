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

package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	"github.com/decabits/vwo-golang-sdk/pkg/testdata"
	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/logger"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/utils"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	User      string `json:"user"`
	Variation string `json:"variation"`
}

func TestPreEvaluateSegment(t *testing.T) {
	vwoInstance := testdata.GetInstanceWithSettings("AB_T_100_W_33_33_33")

	segments := vwoInstance.SettingsFile.Campaigns[0].Segments
	options := schema.Options{
		VariationTargetingVariables: nil,
	}
	value := PreEvaluateSegment(vwoInstance, segments, options, "")
	assert.False(t, value, "Expected False as no segments")
}

func TestEvaluateSegment(t *testing.T) {
	vwoInstance := testdata.GetInstanceWithSettings("T_50_W_50_50_WS")

	segments := vwoInstance.SettingsFile.Campaigns[0].Segments
	options := schema.Options{
		CustomVariables: map[string]interface{}{"a":"123", "hello":"world"},
	}
	value := EvaluateSegment(vwoInstance, segments, options)
	assert.True(t, value, "Expected True as mismatch")
}

func TestGetWhiteListedVariationsList(t *testing.T) {
	vwoInstance := testdata.GetInstanceWithSettings("AB_T_100_W_0_100")

	options := schema.Options{}
	userID := testdata.GetRandomUser()
	campaign := vwoInstance.SettingsFile.Campaigns[0]
	actual := GetWhiteListedVariationsList(vwoInstance, userID, campaign, options)
	assert.Empty(t, actual, "No WhiteListed Variations Found")

	vwoInstance = testdata.GetInstanceWithCustomSettings("SettingsFile3")
	options = schema.Options{
		VariationTargetingVariables: map[string]interface{}{"a": "123"},
		RevenueValue:                12,
	}
	userID = testdata.GetRandomUser()
	campaign = vwoInstance.SettingsFile.Campaigns[0]
	actual = GetWhiteListedVariationsList(vwoInstance, userID, campaign, options)
	expected := campaign.Variations[0:2]
	assert.Equal(t, expected, actual, "No WhiteListed Variations Found")
}

func TestFindTargetedVariation(t *testing.T) {
	assertOutput := assert.New(t)

	// CORNER CASES

	instance := testdata.GetInstanceWithCustomSettings("SettingsFile3")

	campaign := instance.SettingsFile.Campaigns[0]
	options := schema.Options {
		VariationTargetingVariables: map[string]interface{}{"a":"789"},
	}
	actual, _ := FindTargetedVariation(instance, testdata.ValidUser, campaign, options)
	assertOutput.Equal("", actual.Name, "Variations should match")
}

func TestGetVariation(t *testing.T) {
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

	
	instance := schema.VwoInstance {
		Logger : logs,
	}

	for settingsFileName, settingsFile := range settingsFiles {
		vwoInstance := schema.VwoInstance{
			Logger: logs,
		}
		settingsFile.Campaigns[0].Variations = utils.GetVariationAllocationRanges(vwoInstance, settingsFile.Campaigns[0].Variations)

		instance.SettingsFile = settingsFile

		testCases := userExpectation[settingsFileName]
		for i := range testCases {
			campaign, err := utils.GetCampaign("", instance.SettingsFile, settingsFile.Campaigns[0].Key)
			assertOutput.Nil(err, "Incorrect Get Campaign Call")
			actual, _ := GetVariation(instance, testCases[i].User, campaign, schema.Options{})
			expected := testCases[i].Variation
			assertOutput.Equal(expected, actual.Name, settingsFileName+" "+testCases[i].User)
		}
	}

	// CORNER CASES

	var customSettingsFiles map[string]schema.SettingsFile
	data, err = ioutil.ReadFile("../testdata/customSettings.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &customSettingsFiles); err != nil {
		logger.Info("Error: " + err.Error())
	}

	settings := customSettingsFiles["SettingsFile3"]
	instance.SettingsFile = settings

	campaign := instance.SettingsFile.Campaigns[0]
	options := schema.Options {
		VariationTargetingVariables: map[string]interface{}{"a":"123"},
	}
	actual, _ := GetVariation(instance, testdata.ValidUser, campaign, options)
	expected := testdata.ValidVariationControl
	assertOutput.Equal(expected, actual.Name, "Variations should match")

	options = schema.Options {
		VariationTargetingVariables: map[string]interface{}{"b":"456"},
	}
	actual, _ = GetVariation(instance, testdata.ValidUser, campaign, options)
	expected = testdata.ValidVariationVariation2
	assertOutput.Equal(expected, actual.Name, "Variations should match")

	instance = testdata.GetInstanceWithStorage("AB_T_50_W_50_50")
	actual, _ = GetVariation(instance, testdata.TempUser, instance.SettingsFile.Campaigns[0], schema.Options{})
	expected = instance.SettingsFile.Campaigns[0].Variations[0].Name
	assertOutput.Equal(expected, actual.Name, "Variations should match")

	instance = testdata.GetInstanceWithStorage("AB_T_100_W_20_80")
	userID := testdata.GetRandomUser()
	actual, err = GetVariation(instance, userID, instance.SettingsFile.Campaigns[0], schema.Options{})
	assertOutput.NotNil(err, "No Variation Will Be Allcoated")
	assertOutput.Empty(actual, "Variations should be empty : " + userID)

	instance.SettingsFile.Campaigns[0].Variations = utils.GetVariationAllocationRanges(instance, instance.SettingsFile.Campaigns[0].Variations)
	userID = testdata.GetRandomUser()
	actual, err = GetVariation(instance, userID, instance.SettingsFile.Campaigns[0], schema.Options{})
	assertOutput.Equal(nil, err, "No error expected")
	assertOutput.NotEmpty(actual, "Variations should match : " + userID)

	instance = testdata.GetInstanceWithIncorrectStorage("AB_T_100_W_20_80")
	instance.SettingsFile.Campaigns[0].Variations = utils.GetVariationAllocationRanges(instance, instance.SettingsFile.Campaigns[0].Variations)
	userID = testdata.GetRandomUser()
	actual, err = GetVariation(instance, userID, instance.SettingsFile.Campaigns[0], schema.Options{})
	assertOutput.Equal(nil, err, "No error expected")
	assertOutput.NotEmpty(actual, "Variations should match : " + userID)
}

func TestGetVariationFromUserStorage(t *testing.T) {
	assertOutput := assert.New(t)
	vwoInstance := testdata.GetInstanceWithSettings("AB_T_50_W_50_50")

	campaign := vwoInstance.SettingsFile.Campaigns[0]
	userID := testdata.ValidUser
	actual := GetVariationFromUserStorage(vwoInstance, userID, campaign)
	assertOutput.Empty(actual, "Actual and Expected Variation Name mismatch")

	vwoInstance = testdata.GetInstanceWithStorage("AB_T_50_W_50_50")

	campaign = vwoInstance.SettingsFile.Campaigns[0]
	userID = testdata.ValidUser
	expected := testdata.DummyVariation
	actual = GetVariationFromUserStorage(vwoInstance, userID, campaign)
	assertOutput.Equal(expected, actual, "Actual and Expected Variation Name mismatch")

	campaign = vwoInstance.SettingsFile.Campaigns[0]
	userID = testdata.InvalidUser
	expected = ""
	actual = GetVariationFromUserStorage(vwoInstance, userID, campaign)
	assertOutput.Equal(expected, actual, "Actual and Expected Variation Name mismatch")

}
