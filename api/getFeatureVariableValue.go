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
		vwoInstance.Logger.Error("INFO_MESSAGES.INVALID_VARIATION_KEY ", err)
		return nil
	}

	var variable schema.Variable
	if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		variable = utils.GetVariableForFeature(campaign.Variables, variableKey)
	} else if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureTest) {
		variable = utils.GetVariableValueForVariation(vwoInstance, campaign, variation, variableKey)
	}

	if variable.Key == "" {
		vwoInstance.Logger.Info("INFO_MESSAGES.VARIABLE_NOT_FOUND")
	} else {
		vwoInstance.Logger.Info("INFO_MESSAGES.VARIABLE_FOUND", variable)
	}

	return variable.Value
}
