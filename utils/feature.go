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
	"fmt"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
)

const feature = "feature.go"

// GetVariableForFeature gets the variable from the list of variables in the campaign that matches the variableKey
func GetVariableForFeature(variables []schema.Variable, variableKey string) schema.Variable {
	/*
		Args:
			campaign : campaign object
			variableKey: variable Key identifier

		Returns:
			schema.Variable: first variable with the matching variable Key as needed
	*/
	for _, variable := range variables {
		if variable.Key == variableKey {
			return variable
		}
	}
	return schema.Variable{}
}

// GetVariableValueForVariation gets the variable from the list of variables in the variation that matches the variableKey
func GetVariableValueForVariation(vwoInstance schema.VwoInstance, campaign schema.Campaign, variation schema.Variation, variableKey, userID string) schema.Variable {
	/*
		Args:
			campaign : campaign object
			variableKey: variable Key identifier
			variation: variation object

		Returns:
			schema.Variable: first variable with the matching variable Key as needed
	*/

	if !variation.IsFeatureEnabled {
		message := fmt.Sprintf(constants.InfoMessageFeatureEnabledForUser, campaign.Key, userID)
		LogMessage(vwoInstance.Logger, constants.Info, feature, message)
		variation = GetControlVariation(campaign)
		message = fmt.Sprintf(constants.InfoMessageNewVariation, variation)
		LogMessage(vwoInstance.Logger, constants.Info, feature, message)
	}
	message := fmt.Sprintf(constants.InfoMessageFeatureEnabledForUser, campaign.Key, userID)
	LogMessage(vwoInstance.Logger, constants.Info, feature, message)
	return GetVariableForFeature(variation.Variables, variableKey)
}
