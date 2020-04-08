package utils

import (
	"github.com/decabits/vwo-golang-sdk/schema"
)

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
func GetVariableValueForVariation(vwoInstance schema.VwoInstance, campaign schema.Campaign, variation schema.Variation, variableKey string) schema.Variable {
	/*
		Args:
			campaign : campaign object
			variableKey: variable Key identifier
			variation: variation object

		Returns:
			schema.Variable: first variable with the matching variable Key as needed
	*/
	if !variation.IsFeatureEnabled {
		vwoInstance.Logger.Info("INFO_MESSAGES.FEATURE_NOT_ENABLED_FOR_USER")
		variation = GetControlVariation(campaign)
		vwoInstance.Logger.Info("INFO_MESSAGES_NEW_VARIATION", variation)
	}
	vwoInstance.Logger.Info("INFO_MESSAGES.FEATURE_ENABLED_FOR_USER")
	return GetVariableForFeature(variation.Variables, variableKey)
}
