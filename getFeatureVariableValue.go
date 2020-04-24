package vwo

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// GetFeatureVariableValue function
func (vwo *VWOInstance) GetFeatureVariableValue(campaignKey, variableKey, userID string) interface{} {
	options := schema.Options{}
	return vwo.GetFeatureVariableValueWithOptions(campaignKey, variableKey, userID, options)
}

// GetFeatureVariableValueWithOptions function
func (vwo *VWOInstance) GetFeatureVariableValueWithOptions(campaignKey, variableKey, userID string, options schema.Options) interface{} {
	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		vwo.Logger.Error("Error geting campaign: ", err)
		return nil
	}

	if campaign.Status != constants.StatusRunning {
		vwo.Logger.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return nil
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		vwo.Logger.Error("ERROR_MESSAGES.INVALID_API")
		return nil
	}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      vwo.SettingsFile,
		UserStorage:       vwo.UserStorage,
		Logger:            vwo.Logger,
		IsDevelopmentMode: vwo.IsDevelopmentMode,
	}
	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		vwo.Logger.Error("INFO_MESSAGES.INVALID_VARIATION_KEY ", err)
		return nil
	}

	var variable schema.Variable
	if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		variable = utils.GetVariableForFeature(campaign.Variables, variableKey)
	} else if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureTest) {
		variable = utils.GetVariableValueForVariation(vwoInstance, campaign, variation, variableKey)
	}

	if variable.Key == "" {
		vwo.Logger.Info("INFO_MESSAGES.VARIABLE_NOT_FOUND")
	} else {
		vwo.Logger.Info("INFO_MESSAGES.VARIABLE_FOUND", variable)
	}

	return variable.Value
}
