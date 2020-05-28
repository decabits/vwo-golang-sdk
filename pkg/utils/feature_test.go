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

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVariableValueForVariation(t *testing.T) {
	vwoInstance := getInstance()
	campaign := vwoInstance.SettingsFile.Campaigns[2]
	userID := ""

	variation := campaign.Variations[0]
	variableKey := "string1"
	variable := GetVariableValueForVariation(vwoInstance, campaign, variation, variableKey, userID)
	assert.Empty(t, variable, "Expected object should be empty")

	variation = campaign.Variations[0]
	variableKey = "int2"
	variable = GetVariableValueForVariation(vwoInstance, campaign, variation, variableKey, userID)
	assert.Equal(t, campaign.Variations[0].Variables[0], variable, "Expected and Actual IDs should be same")

	campaign = vwoInstance.SettingsFile.Campaigns[5]
	variation = campaign.Variations[1]
	variableKey = "int3"
	variable = GetVariableValueForVariation(vwoInstance, campaign, variation, variableKey, userID)
	assert.Equal(t, campaign.Variations[0].Variables[0], variable, "Expected and Actual IDs should be same")

}

func TestGetVariableForFeature(t *testing.T) {
	vwoInstance := getInstance()

	variables := vwoInstance.SettingsFile.Campaigns[2].Variables
	variableKey := "int1"
	variable := GetVariableForFeature(variables, variableKey)
	assert.Equal(t, variables[0], variable, "Expected and Actual IDs should be same")

	variableKey = "float2"
	variable = GetVariableForFeature(variables, variableKey)
	assert.Empty(t, variable, "Expected variable should be empty")
}
