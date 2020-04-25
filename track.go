package vwo

import (
	"fmt"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

const track = "track.go"

// Track function
func (vwo *VWOInstance) Track(campaignKey, userID string, goalIdentifier string) bool {
	options := schema.Options{}
	return vwo.TrackWithOptions(campaignKey, userID, goalIdentifier, options)
}

// TrackWithOptions function
func (vwo *VWOInstance) TrackWithOptions(campaignKey, userID, goalIdentifier string, options schema.Options) bool {
	if !utils.ValidateTrack(campaignKey, userID, goalIdentifier) {
		return false
	}

	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotFound, campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, track, message)
		utils.LogMessage(vwo.Logger, constants.Error, track, err.Error())
		return false
	}

	if campaign.Status != constants.StatusRunning {
		message := fmt.Sprintf(constants.ErrorMessagesCampaignNotRunning, "Track", campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, track, message)
		return false
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		message := fmt.Sprintf(constants.ErrorMessagesInvalidAPI, "Track", campaignKey, campaign.Type, userID)
		utils.LogMessage(vwo.Logger, constants.Error, track, message)
		return false
	}

	goal, err := utils.GetCampaignGoal(campaign, goalIdentifier)
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessagesTrackAPIGoalNotFound, goalIdentifier, campaignKey, userID)
		utils.LogMessage(vwo.Logger, constants.Error, track, message)
		utils.LogMessage(vwo.Logger, constants.Error, track, err.Error())
		return false
	}

	if goal.Type == constants.GoalTypeRevenue && options.RevenueGoal == 0 {
		message := fmt.Sprintf(constants.ErrorMessagesTrackAPIRevenueNotPassedForRevenueGoal, options.RevenueGoal, campaignKey, userID)
		utils.LogMessage(vwo.Logger, constants.Error, track, message)
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
		message := fmt.Sprintf(constants.InfoMessageInvalidVariationKey, userID, campaignKey)
		utils.LogMessage(vwo.Logger, constants.Info, track, message)
		utils.LogMessage(vwo.Logger, constants.Error, track, err.Error())
		return false
	}

	impression := utils.CreateImpressionTrackingGoal(vwoInstance, variation.ID, userID, campaign.ID, goal.ID, 5) // revenueValue = 5
	event.DispatchTrackingGoal(vwoInstance, impression)

	return true
}
