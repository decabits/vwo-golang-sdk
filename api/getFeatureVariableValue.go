package api

import (
	log "github.com/golang/glog"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// GetFeatureVariableValue ...
func GetFeatureVariableValue(vwoInstance schema.VwoInstance, campaignKey, variableKey, userID string, options schema.Options) schema.Variable {
	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		log.Error("Error geting campaign: ", err)
		return schema.Variable{}
	}

	if campaign.Status != constants.StatusRunning {
		log.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return schema.Variable{}
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		log.Error("ERROR_MESSAGES.INVALID_API")
		return schema.Variable{}
	}

	variation, err := core.GetVariation(vwoInstance, userID, campaign)
	if err != nil {
		log.Error("No Variation Found")
		return schema.Variable{}
	}

	var variable schema.Variable
	if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		variable = utils.GetVariableForFeature(campaign, variableKey)
	} else if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureTest) {
		variable = utils.GetVariableValueForVariation(campaign, variation, variableKey)
		if variable.Key != "" {
			if variation.IsFeatureEnabled {
				log.Info("INFO_MESSAGES.USER_RECEIVED_VARIABLE_VALUE")
			} else {
				log.Info("INFO_MESSAGES.VARIABLE_NOT_USED_RETURN_DEFAULT_VARIABLE_VALUE")
			}
		}
	}

	if variable.Key == "" {
		log.Error("ERROR_MESSAGES.VARIABLE_NOT_FOUND")
	}

	return variable
}
