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

package utils

import (
	"testing"

	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/stretchr/testify/assert"
)

func TestCreateImpressionForPush(t *testing.T) {
	vwoInstance := getInstance()
	URL := "https://dev.visualwebsiteoptimizer.com/server-side/push"

	userID := "Chris"
	tagValue := ""
	tagKey := ""
	DemoImpression := CreateImpressionForPush(vwoInstance, tagKey, tagValue, userID)
	assert.Equal(t, userID, DemoImpression.UID, "Non Matching UIDs")
	assert.Equal(t, vwoInstance.SettingsFile.AccountID, DemoImpression.AccountID, "Non Matching Account IDs")
	assert.Equal(t, URL, DemoImpression.URL, "Non Matching URLs")
	assert.Equal(t, "=", DemoImpression.U, "Non Matching Parameters")

	userID = "Lizzie"
	tagValue = "testVal"
	tagKey = "testKey"
	DemoImpression = CreateImpressionForPush(vwoInstance, tagKey, tagValue, userID)
	assert.Equal(t, userID, DemoImpression.UID, "Non Matching UIDs")
	assert.Equal(t, vwoInstance.SettingsFile.AccountID, DemoImpression.AccountID, "Non Matching Account IDs")
	assert.Equal(t, URL, DemoImpression.URL, "Non Matching URLs")
	assert.Equal(t, "testKey=testVal", DemoImpression.U, "Non Matching Parameters")
}

func TestCreateImpressionTrackingUser(t *testing.T) {
	vwoInstance := getInstance()
	variationID := 1
	campaignID := 283
	userID := "Chris"
	URL := "https://dev.visualwebsiteoptimizer.com/server-side/track-user"

	DemoImpression := CreateImpressionTrackingUser(vwoInstance, campaignID, variationID, userID)
	assert.Equal(t, userID, DemoImpression.UID, "Non Matching UIDs")
	assert.Equal(t, variationID, DemoImpression.Combination, "Non Matching VariationIDs")
	assert.Equal(t, campaignID, DemoImpression.ExperimentID, "Non Matching CampaignIDs")
	assert.Equal(t, vwoInstance.SettingsFile.AccountID, DemoImpression.AccountID, "Non Matching Account IDs")
	assert.Equal(t, URL, DemoImpression.URL, "Non Matching URLs")

}
func TestCreateImpressionTrackingGoal(t *testing.T) {
	vwoInstance := getInstance()
	variationID := 1
	campaignID := 283
	goalID := 281
	goalType:= "rev"
	revenueValue := 1000
	userID := "Chris"
	URL := "https://dev.visualwebsiteoptimizer.com/server-side/track-goal"

	DemoImpression := CreateImpressionTrackingGoal(vwoInstance, variationID, userID, goalType, campaignID, goalID, revenueValue)
	assert.Equal(t, userID, DemoImpression.UID, "Non Matching UIDs")
	assert.Equal(t, variationID, DemoImpression.Combination, "Non Matching VariationIDs")
	assert.Equal(t, goalID, DemoImpression.GoalID, "Non Matching GoalIDs")
	assert.Equal(t, campaignID, DemoImpression.ExperimentID, "Non Matching CampaignIDs")
	assert.Equal(t, revenueValue, DemoImpression.R, "Non Matching Revenues")
	assert.Equal(t, vwoInstance.SettingsFile.AccountID, DemoImpression.AccountID, "Non Matching Account IDs")
	assert.Equal(t, URL, DemoImpression.URL, "Non Matching URLs")
	var temp1 string
	assert.IsType(t, temp1, DemoImpression.SID, "Incorrect SID type")
	var temp2 float32
	assert.IsType(t, temp2, DemoImpression.Random, "Incorrect Random type")
}

func TestGetCommonProperties(t *testing.T) {
	userID := "USER_8"
	vwoInstance := getInstance()
	impression := getCommonProperties(vwoInstance, userID)

	assert.Equal(t, userID, impression.UID, "Non Matching UIDs")
	assert.Equal(t, vwoInstance.SettingsFile.AccountID, impression.AccountID, "Non Matching AccountIDs")
	assert.Equal(t, constants.Platform, impression.Ap, "Non Matching Platforms")
	assert.Equal(t, constants.SDKName, impression.Sdk, "Non Matching SDK")
	assert.Equal(t, constants.SDKVersion, impression.SdkV, "Non Matching SDK Version")
	var temp1 float32
	assert.IsType(t, temp1, impression.Random, "Random number should be of type float32")
	var temp2 string
	assert.IsType(t, temp2, impression.SID, "Random number should be of type float32")

}
