package api

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// GetFeatureVariableValue ...
func GetFeatureVariableValue(vwoInstance schema.VwoInstance, campaignKey, variableKey, userID string, options schema.Options) interface{} {
	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		vwoInstance.Logger.Error("Error geting campaign: ", err)
		return nil
	}

	if campaign.Status != constants.StatusRunning {
		vwoInstance.Logger.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return nil
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		vwoInstance.Logger.Error("ERROR_MESSAGES.INVALID_API")
		return nil
	}

	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		vwoInstance.Logger.Error("INFO_MESSAGES.FEATURE_NOT_ENABLED_FOR_USER ", err)
		return nil
	}

	var variable schema.Variable
	if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		variable = utils.GetVariableForFeature(campaign, variableKey)
	} else if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureTest) {
		variable = utils.GetVariableValueForVariation(campaign, variation, variableKey)
		if variable.Key != "" {
			if variation.IsFeatureEnabled {
				vwoInstance.Logger.Info("INFO_MESSAGES.USER_RECEIVED_VARIABLE_VALUE")
			} else {
				vwoInstance.Logger.Info("INFO_MESSAGES.VARIABLE_NOT_USED_RETURN_DEFAULT_VARIABLE_VALUE")
			}
		}
	}

	if variable.Key == "" {
		vwoInstance.Logger.Error("ERROR_MESSAGES.VARIABLE_NOT_FOUND")
	}

	return variable.Value
}
