package vwo

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

const activate = "activate.go"

// Activate function
func (vwo *VWOInstance) Activate(campaignKey, userID string) string {
	options := schema.Options{}
	return vwo.ActivateWithOptions(campaignKey, userID, options)
}

// ActivateWithOptions ...
func (vwo *VWOInstance) ActivateWithOptions(campaignKey, userID string, options schema.Options) string {
	if !utils.ValidateActivate(campaignKey, userID) {
		return ""
	}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      vwo.SettingsFile,
		UserStorage:       vwo.UserStorage,
		Logger:            vwo.Logger,
		IsDevelopmentMode: vwo.IsDevelopmentMode,
		UserID:            userID,
	}

	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		utils.LogMessage(vwoInstance, constants.Error, activate, "Error getting campaign: "+err.Error())
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

	vwoInstance.Campaign = campaign

	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		vwo.Logger.Error("INFO_MESSAGES.INVALID_VARIATION_KEY")
		return ""
	}

	impression := utils.CreateImpressionTrackingUser(vwoInstance, campaign.ID, variation.ID, userID)
	event.Dispatch(vwoInstance, impression)

	return variation.Name
}
