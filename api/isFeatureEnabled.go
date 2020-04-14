package api

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

//IsFeatureEnabled ...
func IsFeatureEnabled(vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options) bool {
	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		vwoInstance.Logger.Error("Error geting campaign: ", err)
		return false
	}

	if campaign.Status != constants.StatusRunning {
		vwoInstance.Logger.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return false
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		vwoInstance.Logger.Error("ERROR_MESSAGES.INVALID_API")
		return false
	}

	// campaign.Variations = utils.GetVariationAllocationRanges(vwoInstance, campaign.Variations)
	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		vwoInstance.Logger.Error("INFO_MESSAGES.INVALID_VARIATION_KEY")
		return false
	}

	isFeatureEnabled := false
	if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureTest) {
		isFeatureEnabled = variation.IsFeatureEnabled
		impression := utils.CreateImpressionTrackingUser(vwoInstance, campaign.ID, variation.ID, userID)
		event.Dispatch(vwoInstance, impression)
	} else if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		isFeatureEnabled = true
	}

	if isFeatureEnabled {
		vwoInstance.Logger.Info("INFO_MESSAGES.FEATURE_ENABLED_FOR_USER")
	} else {
		vwoInstance.Logger.Info("INFO_MESSAGES.FEATURE_NOT_ENABLED_FOR_USER")
	}

	return isFeatureEnabled
}
