package api

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Activate function
func Activate(vwoInstance schema.VwoInstance, campaignKey, userID string) string {
	options := schema.Options{}
	return ActivateWithOptions(vwoInstance, campaignKey, userID, options)
}

// ActivateWithOptions ...
func ActivateWithOptions(vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options) string {
	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		vwoInstance.Logger.Error("Error geting campaign: ", err)
		return ""
	}

	if campaign.Status != constants.StatusRunning {
		vwoInstance.Logger.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return ""
	}
	if !utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		vwoInstance.Logger.Error("ERROR_MESSAGES.INVALID_API")
		return ""
	}

	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		vwoInstance.Logger.Error("INFO_MESSAGES.INVALID_VARIATION_KEY")
		return ""
	}

	impression := utils.CreateImpressionTrackingUser(vwoInstance, campaign.ID, variation.ID, userID)
	event.Dispatch(vwoInstance, impression)

	return variation.Name
}
