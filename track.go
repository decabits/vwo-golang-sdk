package vwo

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Track function
func (vwo *VWOInstance) Track(campaignKey, userID string, goalIdentifier string) bool {
	options := schema.Options{}
	return vwo.TrackWithOptions(campaignKey, userID, goalIdentifier, options)
}

// TrackWithOptions function
func (vwo *VWOInstance) TrackWithOptions(campaignKey, userID string, goalIdentifier string, options schema.Options) bool {
	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		vwo.Logger.Errorf("Error geting campaign: %+v\n", err)
		return false
	}

	if campaign.Status != constants.StatusRunning {
		vwo.Logger.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return false
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		vwo.Logger.Error("ERROR_MESSAGES.INVALID_API")
		return false
	}

	goal, err := utils.GetCampaignGoal(campaign, goalIdentifier)
	if err != nil {
		vwo.Logger.Errorf("ERROR_MESSAGES.TRACK_API_GOAL_NOT_FOUND: %+v\n", err)
		return false
	}

	if goal.Type == constants.GoalTypeRevenue && options.RevenueGoal == 0 {
		vwo.Logger.Error("ERROR_MESSAGES.TRACK_API_REVENUE_NOT_PASSED_FOR_REVENUE_GOAL")
		return false
	}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      vwo.SettingsFile,
		UserStorage:       vwo.UserStorage,
		Logger:            vwo.Logger,
		IsDevelopmentMode: vwo.IsDevelopmentMode,
	}
	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		vwo.Logger.Errorf("INFO_MESSAGES.INVALID_VARIATION_KEY %+v\n", err)
		return false
	}

	impression := utils.CreateImpressionTrackingGoal(vwoInstance, variation.ID, userID, campaign.ID, goal.ID, 5) // revenueValue = 5
	event.DispatchTrackingGoal(vwoInstance, impression)

	return true
}
