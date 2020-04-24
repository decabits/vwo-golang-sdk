package vwo

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// GetVariationName ...
func (vwo *VWOInstance) GetVariationName(campaignKey, userID string) string {
	options := schema.Options{}
	return vwo.GetVariationNameWithOptions(campaignKey, userID, options)
}

// GetVariationNameWithOptions
func (vwo *VWOInstance) GetVariationNameWithOptions(campaignKey, userID string, options schema.Options) string {
	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		vwo.Logger.Error("Error geting campaign: ", err)
		return ""
	}

	if campaign.Status != constants.StatusRunning {
		vwo.Logger.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return ""
	}
	if !utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		vwo.Logger.Error("ERROR_MESSAGES.INVALID_API")
		return ""
	}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      vwo.SettingsFile,
		UserStorage:       vwo.UserStorage,
		Logger:            vwo.Logger,
		IsDevelopmentMode: vwo.IsDevelopmentMode,
	}
	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		vwo.Logger.Error("INFO_MESSAGES.INVALID_VARIATION_KEY")
	}

	return variation.Name
}
