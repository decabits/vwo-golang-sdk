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
	"io/ioutil"
	"log"
	"testing"

	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/logger"
	"github.com/decabits/vwo-golang-sdk/pkg/testdata"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/service"
	"github.com/stretchr/testify/assert"
)

func getInstanceWithoutStorage(path string) schema.VwoInstance {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.ProcessSettingsFile(path); err != nil {
		log.Println("Error Processing Settings File: ", err)
	}
	settingsFileManager.Process()
	settingsFile := settingsFileManager.GetSettingsFile()

	logs := logger.Init(constants.SDKName, false, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settingsFile,
		UserStorage:       nil,
		Logger:            logs,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}

func TestPreEvaluateSegment(t *testing.T) {
	vwoInstance := testdata.GetInstanceWithSettings("AB_T_100_W_33_33_33")

	segments := vwoInstance.SettingsFile.Campaigns[0].Segments
	options := schema.Options{}
	value := PreEvaluateSegment(vwoInstance, segments, options, "")
	assert.False(t, value, "Expected False as no segments")

	// segments = vwoInstance.SettingsFile.Campaigns[0].Variations[0].Segments
	// options = schema.Options{
	// 	VariationTargetingVariables: map[string]interface{}{"_vwo_user_id": "USER_2"},
	// 	RevenueValue:                12,
	// }
	// value = PreEvaluateSegment(vwoInstance, segments, options, "")
	// assert.True(t, value, "Expected True")

	// options = schema.Options{
	// 	VariationTargetingVariables: map[string]interface{}{"_vwo_user_id": "USER_9"},
	// 	RevenueValue:                12,
	// }
	// value = PreEvaluateSegment(vwoInstance, segments, options, "")
	// assert.False(t, value, "Expected True")
}

func TestEvaluateSegment(t *testing.T) {
	vwoInstance := getInstanceWithoutStorage("./testdata/testVariation.json")

	segments := vwoInstance.SettingsFile.Campaigns[0].Segments
	options := schema.Options{}
	value := EvaluateSegment(vwoInstance, segments, options)
	assert.True(t, value, "Expected False as mismatch")

	segments = vwoInstance.SettingsFile.Campaigns[0].Variations[0].Segments
	options = schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "USER_1"},
		RevenueValue:    12,
	}
	value = EvaluateSegment(vwoInstance, segments, options)
	assert.True(t, value, "Expected True")

	options = schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "USER_3"},
		RevenueValue:    12,
	}
	value = EvaluateSegment(vwoInstance, segments, options)
	assert.False(t, value, "Expected True")
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
	expected := campaign.Variations
	assert.Equal(t, expected, actual, "No WhiteListed Variations Found")
}

func TestFindTargetedVariation(t *testing.T) {
	assertOutput := assert.New(t)

	vwoInstance := getInstanceWithoutStorage("./testdata/testVariation.json")
	options := schema.Options{
		VariationTargetingVariables: map[string]interface{}{"_vwo_user_id": "USER_1"},
		RevenueValue:                12,
	}

	userID := "USER_1"
	campaign := vwoInstance.SettingsFile.Campaigns[2]
	actual, err := FindTargetedVariation(vwoInstance, userID, campaign, options)
	assertOutput.Empty(actual, "Forced variation Disabled")
	assertOutput.NotNil(err, "Error Expected")

	userID = "USER_1"
	campaign = vwoInstance.SettingsFile.Campaigns[3]
	actual, err = FindTargetedVariation(vwoInstance, userID, campaign, options)
	assertOutput.Empty(actual, "Expected no variation")
	assertOutput.NotNil(err, "Error Expected")

	userID = "USER_1"
	campaign = vwoInstance.SettingsFile.Campaigns[4]
	actual, err = FindTargetedVariation(vwoInstance, userID, campaign, options)
	expected := campaign.Variations[0]
	assertOutput.Equal(expected, actual, "Expected single variation")
	assertOutput.Nil(err, "Error Expected")

	userID = "USER_1"
	campaign = vwoInstance.SettingsFile.Campaigns[5]
	actual, err = FindTargetedVariation(vwoInstance, userID, campaign, options)
	expected = campaign.Variations[1]
	assertOutput.Equal(expected, actual, "Expected single variation")
	assertOutput.Nil(err, "Error Expected")

	options = schema.Options{
		CustomVariables: map[string]interface{}{"_vwo_user_id": "USER_8"},
		RevenueValue:    12,
	}
	userID = "USER_8"
	campaign = vwoInstance.SettingsFile.Campaigns[6]
	actual, err = FindTargetedVariation(vwoInstance, userID, campaign, options)
	assertOutput.Empty(actual, "No Varaition Expected ")
	assertOutput.NotNil(err, "Error Expected")
}

func TestGetVariation(t *testing.T) {
	// assertOutput := assert.New(t)

	// vwoInstance := getInstanceWithoutStorage("./testdata/testVariation.json")
	// options := schema.Options{
	// 	CustomVariables: map[string]interface{}{"_vwo_user_id": "USER_1"},
	// 	RevenueValue:    12,
	// }

	// userID := "USER_1"
	// campaign := vwoInstance.SettingsFile.Campaigns[4]
	// actual, err := GetVariation(vwoInstance, userID, campaign, options)
	// expected := campaign.Variations[0]
	// assertOutput.Nil(err, "Variation mis match")
	// assertOutput.Equal(expected, actual, "Variation mis match")

	// userID = "USER_3"
	// campaign = vwoInstance.SettingsFile.Campaigns[0]
	// actual, err = GetVariation(vwoInstance, userID, campaign, options)
	// expected = campaign.Variations[0]
	// assertOutput.Nil(err, "Variation not found in userStorage")
	// assertOutput.Equal(expected, actual, "Variation not found in userStorage")

	// userID = "USER_8"
	// campaign = vwoInstance.SettingsFile.Campaigns[7]
	// actual, err = GetVariation(vwoInstance, userID, campaign, options)
	// assertOutput.NotNil(err, "Variation not in campaign")
	// assertOutput.Empty(actual, "Variation not in campaign")

	// userID = "Dummy_USER_8"
	// campaign = vwoInstance.SettingsFile.Campaigns[3]
	// actual, err = GetVariation(vwoInstance, userID, campaign, options)
	// assertOutput.NotNil(err, "User not eligible for campaign")
	// assertOutput.Empty(actual, "User not eligible for campaign")

	// vwoInstance = testdata.GetInstanceWithStorage("AB_T_50_W_50_50")
	// options = schema.Options{
	// 	CustomVariables: map[string]interface{}{"_vwo_user_id": "USER_1"},
	// 	RevenueValue:    12,
	// }

	// userID = "user1"
	// campaign = vwoInstance.SettingsFile.Campaigns[3]
	// expected = campaign.Variations[0]
	// actual, err = GetVariation(vwoInstance, userID, campaign, options)
	// assertOutput.Nil(err, "User not eligible for campaign")
	// assertOutput.Equal(expected, actual, "User not eligible for campaign")

	// userID = "user1"
	// campaign = vwoInstance.SettingsFile.Campaigns[0]
	// actual, err = GetVariation(vwoInstance, userID, campaign, options)
	// assertOutput.Nil(err, "Actual and Expected Variation Name mismatch")
	// assertOutput.NotEmpty(actual, "Actual and Expected Variation Name mismatch")
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
	userID = "DummyUser"
	expected := "DummyVariation"
	actual = GetVariationFromUserStorage(vwoInstance, userID, campaign)
	assert.Equal(t, expected, actual, "Actual and Expected Variation Name mismatch")
	
	campaign = vwoInstance.SettingsFile.Campaigns[0]
	userID = "DummyUser1"
	expected = ""
	actual = GetVariationFromUserStorage(vwoInstance, userID, campaign)
	assert.Equal(t, expected, actual, "Actual and Expected Variation Name mismatch")

}
