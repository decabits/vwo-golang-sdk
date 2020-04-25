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
