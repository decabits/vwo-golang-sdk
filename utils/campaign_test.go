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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/google/logger"
	"github.com/stretchr/testify/assert"
)

// getInstance function creates and return a temporary VWO instance for testing
func getInstance() schema.VwoInstance {
	settingsFile, err := ioutil.ReadFile("./testData/testUtils.json")
	if err != nil {
		fmt.Println(err)
	}

	var settings schema.SettingsFile
	if err = json.Unmarshal(settingsFile, &settings); err != nil {
		fmt.Println(err)
	}

	logger := logger.Init(constants.SDKName, false, false, ioutil.Discard)
	defer logger.Close()

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settings,
		UserStorage:       nil,
		Logger:            logger,
		IsDevelopmentMode: true,
	}

	return vwoInstance
}

func TestGetVariationBucketingRange(t *testing.T) {
	var weight float64

	weight = 0
	actual := GetVariationBucketingRange(weight)
	expected := 0
	assert.Equal(t, expected, actual, "Expected and Actual Ranges should be same")

	weight = 33.333
	actual = GetVariationBucketingRange(weight)
	expected = 3334
	assert.Equal(t, expected, actual, "Expected and Actual Ranges should be same")

	weight = 102
	actual = GetVariationBucketingRange(weight)
	expected = constants.MaxTrafficValue
	assert.Equal(t, expected, actual, "Expected and Actual Ranges should be same")
}

func TestGetCampaign(t *testing.T) {
	vwoInstance := getInstance()

	campaignKey := "phpab1"
	campaign, err := GetCampaign(vwoInstance.SettingsFile, campaignKey)
	assert.Nil(t, err)
	assert.Equal(t, vwoInstance.SettingsFile.Campaigns[0], campaign, "Expected and Actual Campaign IDs should be same")

	campaignKey = "notAvailable"
	campaign, err = GetCampaign(vwoInstance.SettingsFile, campaignKey)
	assert.NotNil(t, err)
	assert.Empty(t, campaign, "Expected campaign should be empty")
}

func TestGetCampaignVariation(t *testing.T) {
	vwoInstance := getInstance()
	campaign := vwoInstance.SettingsFile.Campaigns[1]

	variationName := "Control"
	variation, err:= GetCampaignVariation(campaign, variationName)
	assert.Nil(t, err)
	assert.Equal(t, campaign.Variations[0], variation, "Expected and Actual Variation IDs should be same")

	variationName = "Variation-3"
	variation, err= GetCampaignVariation(campaign, variationName)
	assert.NotNil(t, err)
	assert.Empty(t, variation, "Expected Variation should be empty")

	campaign = vwoInstance.SettingsFile.Campaigns[0]
	variationName = "Control"
	variation, err= GetCampaignVariation(campaign, variationName)
	assert.NotNil(t, err)
	assert.Empty(t, variation, "Expected and Actual Variation IDs should be same")
}

func TestGetCampaignGoal(t *testing.T) {
	vwoInstance := getInstance()
	campaign := vwoInstance.SettingsFile.Campaigns[1]

	goalName := "rev"
	goal, err:= GetCampaignGoal(campaign, goalName)
	assert.Nil(t, err)
	assert.Equal(t, campaign.Goals[0], goal, "Expected and Actual Goal IDs should be same")

	goalName = "demo"
	goal, err= GetCampaignGoal(campaign, goalName)
	assert.NotNil(t, err)
	assert.Empty(t, goal, "Expected Goal should be empty")
}

func TestGetControlVariation(t *testing.T) {
	vwoInstance := getInstance()

	campaign := vwoInstance.SettingsFile.Campaigns[1]
	variation := GetControlVariation(campaign)
	assert.Equal(t, campaign.Variations[0], variation, "Expected variation should be present in the campaign")

	campaign = vwoInstance.SettingsFile.Campaigns[2]
	variation = GetControlVariation(campaign)
	assert.Empty(t, variation, "Expected variation should be empty")
}

func TestScaleVariations(t *testing.T) {
	vwoInstance := getInstance()

	variations := vwoInstance.SettingsFile.Campaigns[3].Variations
	variations = ScaleVariations(variations)
	assert.Equal(t, vwoInstance.SettingsFile.Campaigns[3].Variations, variations, "List of variations did not match")
	assert.Equal(t, 100.0, variations[0].Weight, "Variation weight did not match")

	variations = vwoInstance.SettingsFile.Campaigns[4].Variations
	variations = ScaleVariations(variations)
	assert.Equal(t, vwoInstance.SettingsFile.Campaigns[4].Variations, variations, "List of variations did not match")
}

func TestGetVariationAllocationRanges(t *testing.T) {
	vwoInstance := getInstance()

	variations := vwoInstance.SettingsFile.Campaigns[3].Variations
	assert.NotEmpty(t, variations, "No Variations recieved")
	variations = GetVariationAllocationRanges(vwoInstance, variations)
	assert.Equal(t, 1, variations[0].StartVariationAllocation, "Value Mismatch")
	assert.Equal(t, 10000, variations[0].EndVariationAllocation, "Value Mismatch")

	variations = vwoInstance.SettingsFile.Campaigns[4].Variations
	assert.NotEmpty(t, variations, "No Variations recieved")
	variations = GetVariationAllocationRanges(vwoInstance, variations)
	assert.Equal(t, -1, variations[0].StartVariationAllocation, "Start Allocation range failed to match")
	assert.Equal(t, -1, variations[0].EndVariationAllocation, "End Allocation range failed to match")

}

func TestMin(t *testing.T) {
	assert.Equal(t, 10, min(10, 20), "Incorrect")
	assert.Equal(t, 10, min(20, 10), "Incorrect")
	assert.NotEqual(t, 12, min(10, 20), "Incorrect")
}

func TestMax(t *testing.T) {
	assert.Equal(t, 20, max(10, 20), "Incorrect")
	assert.Equal(t, 20, max(20, 10), "Incorrect")
	assert.NotEqual(t, 12, max(10, 20), "Incorrect")
}
