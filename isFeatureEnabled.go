package vwo

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// IsFeatureEnabled ...
func (vwo *VWOInstance) IsFeatureEnabled(campaignKey, userID string) bool {
	options := schema.Options{}
	return vwo.IsFeatureEnabledWithOptions(campaignKey, userID, options)
}

// IsFeatureEnabledWithOptions function
func (vwo *VWOInstance) IsFeatureEnabledWithOptions(campaignKey, userID string, options schema.Options) bool {
	if !utils.ValidateIsFeatureEnabled(campaignKey, userID) {
		return false
	}

	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		vwo.Logger.Error("Error geting campaign: ", err)
		return false
	}

	if campaign.Status != constants.StatusRunning {
		vwo.Logger.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return false
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		vwo.Logger.Error("ERROR_MESSAGES.INVALID_API")
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
		vwo.Logger.Error("INFO_MESSAGES.INVALID_VARIATION_KEY")
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
		vwo.Logger.Info("INFO_MESSAGES.FEATURE_ENABLED_FOR_USER")
	} else {
		vwo.Logger.Info("INFO_MESSAGES.FEATURE_NOT_ENABLED_FOR_USER")
	}

	return isFeatureEnabled
}
