package api

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Track function
func Track(vwoInstance schema.VwoInstance, campaignKey, userID string, goalIdentifier string) bool {
	options := schema.Options{}
	return TrackWithOptions(vwoInstance, campaignKey, userID, goalIdentifier, options)
}

// TrackWithOptions function
func TrackWithOptions(vwoInstance schema.VwoInstance, campaignKey, userID string, goalIdentifier string, options schema.Options) bool {
	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		vwoInstance.Logger.Error("Error geting campaign: %+v\n", err)
		return false
	}

	if campaign.Status != constants.StatusRunning {
		vwoInstance.Logger.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return false
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		vwoInstance.Logger.Error("ERROR_MESSAGES.INVALID_API")
		return false
	}

	goal, err := utils.GetCampaignGoal(campaign, goalIdentifier)
	if err != nil {
		vwoInstance.Logger.Errorf("ERROR_MESSAGES.TRACK_API_GOAL_NOT_FOUND: %+v\n", err)
		return false
	}

	if goal.Type == constants.GoalTypeRevenue && options.RevenueGoal == 0 {
		vwoInstance.Logger.Error("ERROR_MESSAGES.TRACK_API_REVENUE_NOT_PASSED_FOR_REVENUE_GOAL")
		return false
	}

	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		vwoInstance.Logger.Errorf("INFO_MESSAGES.INVALID_VARIATION_KEY %+v\n", err)
		return false
	}

	impression := utils.CreateImpressionTrackingGoal(vwoInstance, variation.ID, userID, campaign.ID, goal.ID, 5) // revenueValue = 5
	event.DispatchTrackingGoal(vwoInstance, impression) 

	return true
}
