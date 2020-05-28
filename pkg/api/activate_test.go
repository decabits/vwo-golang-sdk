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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActivate(t *testing.T) {
	assertOutput := assert.New(t)

	vwoInstance, err := getInstance("./testdata/testdata.json")
	assertOutput.Nil(err, "error fetching instance")

	userID := ""
	campaignKey := ""
	value := vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.Empty(value, "Invalid params")

	userID = "USER_1"
	campaignKey = "notPresent"
	value = vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign does not exist")

	userID = "USER_1"
	campaignKey = "CAMPAIGN_8"
	value = vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign Not running")

	userID = "USER_3"
	campaignKey = "CAMPAIGN_1"
	value = vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign Not Valid")

	userID = "USER_3"
	campaignKey = "CAMPAIGN_9"
	value = vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.Empty(value, "No Variation in Campaign")

	userID = "USER_3"
	campaignKey = "CAMPAIGN_10"
	actual := vwoInstance.Activate(campaignKey, userID, nil)
	assertOutput.NotEmpty(actual, "Variation should be found")
}
