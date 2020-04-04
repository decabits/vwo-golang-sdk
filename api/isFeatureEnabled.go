package api

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/service"
	"github.com/decabits/vwo-golang-sdk/utils"
)

//IsFeatureEnabled ...
func IsFeatureEnabled(vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options) bool {
	settingsFileManager := service.SettingsFileManager{}
	vwoInstance.SettingsFile = settingsFileManager.GetSettingsFile()

	if options.CustomVariables == nil || options.VariationTargetingVariables == nil {
		return false
	}

	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		vwoInstance.Logger.Error("Error geting campaign: ", err)
		return false
	}

	if campaign.Status != constants.StatusRunning {
		vwoInstance.Logger.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return false
	}
	if campaign.Type == constants.CampaignTypeVisualAB {
		vwoInstance.Logger.Error("ERROR_MESSAGES.INVALID_API")
		return false
	}

	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		vwoInstance.Logger.Error("No Variation Found")
		return false
	}

	if campaign.Type == constants.CampaignTypeFeatureTest {
		impression := utils.CreateImpressionTrackingUser(vwoInstance, campaign.ID, variation.ID, userID)
		if !event.Dispatch(vwoInstance, impression) {
			return false
		}
		result := variation.IsFeatureEnabled
		if result {
			vwoInstance.Logger.Info("Feature Enabled For User")
		} else {
			vwoInstance.Logger.Info("Feature Not Enabled For User")
		}
		return result
	}
	return false
}
