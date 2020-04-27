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

package vwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVariationName(t *testing.T) {
	assertOutput := assert.New(t)

	vwoInstance := VWOInstance{}
	err := vwoInstance.getInstance("./testdata/testdata.json")
	assertOutput.Nil(err, "error fetching instance")

	userID := ""
	campaignKey := ""
	value := vwoInstance.GetVariationName(campaignKey, userID, nil)
	assertOutput.Empty(value, "Invalid params")

	userID = "Varun"
	campaignKey = "notPresent"
	value = vwoInstance.GetVariationName(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign does not exist")

	userID = "Varun"
	campaignKey = "phpab1"
	value = vwoInstance.GetVariationName(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign Not running")

	userID = "Liza"
	campaignKey = "php1"
	value = vwoInstance.GetVariationName(campaignKey, userID, nil)
	assertOutput.Empty(value, "Campaign Not Valid")

	userID = "Liza"
	campaignKey = "phpab2"
	value = vwoInstance.GetVariationName(campaignKey, userID, nil)
	assertOutput.Empty(value, "Variation Not found")

	userID = "Liza"
	campaignKey = "phpab3"
	actual := vwoInstance.GetVariationName(campaignKey, userID, nil)
	assertOutput.NotEmpty(actual, "Variation Name does not match")
}
