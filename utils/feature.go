package utils

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
)

// GetVariableForFeature gets the variable from the list of variables in the campaign that matches the variableKey
func GetVariableForFeature(campaign schema.Campaign, variableKey string) schema.Variable {
	/*
		Args: 
			campaign : campaign object
			variableKey: variable Key identifier

		Returns:
			schema.Variable: first variable with the matching variable Key as needed
	*/
	if CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		variables := campaign.Variables
		for _, variable := range variables {
			if variable.Key == variableKey {
				return variable
			}
		}
	}
	return schema.Variable{}
}

// GetVariableValueForVariation gets the variable from the list of variables in the variation that matches the variableKey
func GetVariableValueForVariation(campaign schema.Campaign, variation schema.Variation, variableKey string) schema.Variable {
	/*
		Args:
			campaign : campaign object
			variableKey: variable Key identifier
			variation: variation object
			
		Returns:
			schema.Variable: first variable with the matching variable Key as needed
	*/
	if CheckCampaignType(campaign, constants.CampaignTypeFeatureTest) {
		if !variation.IsFeatureEnabled {
			variation = GetControlVariation(campaign)
		}
		if len(variation.Variables) == 0 {
			return schema.Variable{}
		}
		for _, variable := range variation.Variables {
			if variable.Key == variableKey {
				return variable
			}
		}
	}
	return schema.Variable{}
}

//GetTypeCastedFeatureValue ...
// func GetTypeCastedFeatureValue(variableValue interface{}, variableType string) interface{}{

// }
