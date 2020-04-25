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
/*
This API method: Marks the conversion of the campaign for a particular goal
1. Validates the arguments being passed
2. Finds the corresponding Campaign
3. Checks the Campaign Status
4. Validates the Campaign Type
5. Gets passed Goal
6. Validates Revenue and Goal type
7. Assigns the determinitic variation to the user(based on userId), if user becomes part of campaign
   If userStorageService is used, it will look into it for the variation and if found, no further processing is done
8. If feature enabled, sends a call to VWO server for tracking visitor
*/
func (vwo *VWOInstance) Track(campaignKey, userID string, goalIdentifier string) bool {
	options := schema.Options{}
	return vwo.TrackWithOptions(campaignKey, userID, goalIdentifier, options)
}

// TrackWithOptions function
func (vwo *VWOInstance) TrackWithOptions(campaignKey, userID, goalIdentifier string, options schema.Options) bool {
	if !utils.ValidateTrack(campaignKey, userID, goalIdentifier) {
		message := fmt.Sprintf(constants.ErrorMessagesTrackAPIMissingParams)
		utils.LogMessage(vwo.Logger, constants.Error, track, message)
		return false
	}

	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotFound+" \n %v", campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Error, track, message)
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
		message := fmt.Sprintf(constants.ErrorMessagesTrackAPIGoalNotFound+" \n %v", goalIdentifier, campaignKey, userID, err.Error())
		utils.LogMessage(vwo.Logger, constants.Error, track, message)
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
		message := fmt.Sprintf(constants.InfoMessageInvalidVariationKey+" \n %v", userID, campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Info, track, message)
		return false
	}

	impression := utils.CreateImpressionTrackingGoal(vwoInstance, variation.ID, userID, campaign.ID, goal.ID, 5) // revenueValue = 5
	event.DispatchTrackingGoal(vwoInstance, impression)

	return true
}
