package vwo

import (
	"fmt"
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

const activate = "activate.go"

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
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotFound, campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, activate, message)
		utils.LogMessage(vwo.Logger, constants.Error, activate, err.Error())
		return ""
	}

	if campaign.Status != constants.StatusRunning {
		message := fmt.Sprintf(constants.ErrorMessagesCampaignNotRunning, "Activate", campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, activate, message)
		return ""
	}
	if !utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		message := fmt.Sprintf(constants.ErrorMessagesInvalidAPI, "Activate", campaignKey, campaign.Type, userID)
		utils.LogMessage(vwo.Logger, constants.Error, activate, message)
		return ""
	}

	vwoInstance.Campaign = campaign

	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		message := fmt.Sprintf(constants.InfoMessageInvalidVariationKey, userID, campaignKey)
		utils.LogMessage(vwo.Logger, constants.Info, activate, message)
		utils.LogMessage(vwo.Logger, constants.Error, activate, err.Error())
		return ""
	}

	impression := utils.CreateImpressionTrackingUser(vwoInstance, campaign.ID, variation.ID, userID)
	event.Dispatch(vwoInstance, impression)

	return variation.Name
}
