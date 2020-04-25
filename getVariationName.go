package vwo

import (
	"fmt"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

const getVariationName = "getVariationName.go"

// GetVariationName ...
func (vwo *VWOInstance) GetVariationName(campaignKey, userID string) string {
	options := schema.Options{}
	return vwo.GetVariationNameWithOptions(campaignKey, userID, options)
}

// GetVariationNameWithOptions ...
func (vwo *VWOInstance) GetVariationNameWithOptions(campaignKey, userID string, options schema.Options) string {
	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotFound, campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, message)
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, err.Error())
		return ""
	}

	if campaign.Status != constants.StatusRunning {
		message := fmt.Sprintf(constants.ErrorMessagesCampaignNotRunning, "GetVariationName", campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, message)
		return ""
	}
	if !utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		message := fmt.Sprintf(constants.ErrorMessagesInvalidAPI, "GetVariationName", campaignKey, campaign.Type, userID)
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, message)
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
		message := fmt.Sprintf(constants.InfoMessageInvalidVariationKey, userID, campaignKey)
		utils.LogMessage(vwo.Logger, constants.Info, getVariationName, message)
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, err.Error())
	}

	return variation.Name
}
