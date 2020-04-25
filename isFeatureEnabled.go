package vwo

import (
	"fmt"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

const fileIsFeatureEnabled = "isFeatureEnabled.go"

// IsFeatureEnabled function
/*
This API method: Whether a feature is enabled or not for the given user
1. Validates the arguments being passed
2. Finds the corresponding Campaign
3. Checks the Campaign Status
4. Validates the Campaign Type
5. Assigns the determinitic variation to the user(based on userId), if user becomes part of campaign
   If userStorageService is used, it will look into it for the variation and if found, no further processing is done
6. If feature enabled, sends a call to VWO server for tracking visitor
*/
func (vwo *VWOInstance) IsFeatureEnabled(campaignKey, userID string) bool {
	options := schema.Options{}
	return vwo.IsFeatureEnabledWithOptions(campaignKey, userID, options)
}

// IsFeatureEnabledWithOptions function
func (vwo *VWOInstance) IsFeatureEnabledWithOptions(campaignKey, userID string, options schema.Options) bool {
	if !utils.ValidateIsFeatureEnabled(campaignKey, userID) {
		message := fmt.Sprintf(constants.ErrorMessagesIsFeatureEnabledAPIMissingParams)
		utils.LogMessage(vwo.Logger, constants.Error, fileIsFeatureEnabled, message)
		return false
	}

	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotFound+" \n %v", campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Error, fileIsFeatureEnabled, message)
		return false
	}

	if campaign.Status != constants.StatusRunning {
		message := fmt.Sprintf(constants.ErrorMessagesCampaignNotRunning, "IsFeatureEnabled", campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, fileIsFeatureEnabled, message)
		return false
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		message := fmt.Sprintf(constants.ErrorMessagesInvalidAPI, "IsFeatureEnabled", campaignKey, campaign.Type, userID)
		utils.LogMessage(vwo.Logger, constants.Error, fileIsFeatureEnabled, message)
		return false
	}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      vwo.SettingsFile,
		UserStorage:       vwo.UserStorage,
		Logger:            vwo.Logger,
		IsDevelopmentMode: vwo.IsDevelopmentMode,
		UserID:            userID,
		Campaign:          campaign,
	}
	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		message := fmt.Sprintf(constants.InfoMessageInvalidVariationKey+" \n %v", userID, campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Info, fileIsFeatureEnabled, message)
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
		message := fmt.Sprintf(constants.InfoMessageFeatureEnabledForUser, campaignKey, userID)
		utils.LogMessage(vwo.Logger, constants.Info, fileIsFeatureEnabled, message)
	} else {
		message := fmt.Sprintf(constants.InfoMessageFeatureNotEnabledForUser, campaignKey, userID)
		utils.LogMessage(vwo.Logger, constants.Info, fileIsFeatureEnabled, message)
	}

	return isFeatureEnabled
}
