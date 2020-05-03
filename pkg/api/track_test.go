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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrack(t *testing.T) {
	assertOutput := assert.New(t)

	vwoInstance, err := getInstance("./testdata/testTrack.json")
	assertOutput.Nil(err, "error fetching instance")

	userID := ""
	campaignKey := ""
	goalIdentifier := ""
	value := vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Invalid params")

	userID = "USER_1"
	campaignKey = "notPresent"
	goalIdentifier = "GOAL_2"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Campaign does not exist")

	userID = "USER_1"
	campaignKey = "CAMPAIGN_8"
	goalIdentifier = "GOAL_2"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Campaign Not running")

	userID = "USER_3"
	campaignKey = "CAMPAIGN_1"
	goalIdentifier = "GOAL_2"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Campaign Not Valid")

	userID = "USER_3"
	campaignKey = "CAMPAIGN_9"
	goalIdentifier = "test"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Goal Not Found")

	userID = "USER_3"
	campaignKey = "CAMPAIGN_10"
	goalIdentifier = "GOAL_1"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)
	assertOutput.False(value, "Invalid Goal type")

	userID = "USER_3"
	campaignKey = "CAMPAIGN_10"
	option := map[string]interface{}{
		"revenueGoal": 0,
	}
	goalIdentifier = "GOAL_1"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, option)
	assertOutput.False(value, "Revenue Not defined")

	userID = "USER_3"
	campaignKey = "CAMPAIGN_2"
	option = map[string]interface{}{
		"revenueGoal": 10,
	}
	goalIdentifier = "GOAL_2"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, option)
	assertOutput.False(value, "No Variation in Campaign")

	option = map[string]interface{}{
		"revenueGoal": 12,
	}
	userID = "USER_9"
	campaignKey = "CAMPAIGN_10"
	goalIdentifier = "GOAL_2"
	value = vwoInstance.Track(campaignKey, userID, goalIdentifier, option)
	assertOutput.True(value, "Variation should be defined")
}
