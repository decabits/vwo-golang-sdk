package utils

import (
	"github.com/decabits/vwo-golang-sdk/schema"
)

// GetVariableForFeature function
func GetVariableForFeature(variables []schema.Variable, variableKey string) schema.Variable {
	for _, variable := range variables {
		if variable.Key == variableKey {
			return variable
		}
	}
	return schema.Variable{}
}

// GetVariableValueForVariation function
func GetVariableValueForVariation(vwoInstance schema.VwoInstance, campaign schema.Campaign, variation schema.Variation, variableKey string) schema.Variable {
	if !variation.IsFeatureEnabled {
		vwoInstance.Logger.Info("INFO_MESSAGES.FEATURE_NOT_ENABLED_FOR_USER")
		variation = GetControlVariation(campaign)
		vwoInstance.Logger.Info("INFO_MESSAGES_NEW_VARIATION", variation)
	}
	vwoInstance.Logger.Info("INFO_MESSAGES.FEATURE_ENABLED_FOR_USER")
	return GetVariableForFeature(variation.Variables, variableKey)
}
