package api

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

//GetVariationName ...
func GetVariationName(vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options) string {
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

	// campaign.Variations = utils.GetVariationAllocationRanges(vwoInstance, campaign.Variations)
	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		vwoInstance.Logger.Error("INFO_MESSAGES.INVALID_VARIATION_KEY")
	}

	return variation.Name
}
