package api

import (
	log "github.com/golang/glog"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Track function
func Track(vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options, goalIdentifier string) bool {
	if options.CustomVariables == nil || options.VariationTargetingVariables == nil {
		return false
	}

	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		log.Error("Error geting campaign: ", err)
		return false
	}

	if campaign.Status != constants.StatusRunning {
		log.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return false
	}
	if campaign.Type != constants.CampaignTypeFeatureRollout {
		log.Error("ERROR_MESSAGES.INVALID_API")
		return false
	}

	goal, err := utils.GetCampaignGoal(campaign, goalIdentifier)
	if err != nil {
		return false
	}

	if goal.Type == constants.GoalTypeRevenue && options.RevenueGoal > 0 {
		log.Error("ERROR_MESSAGES.TRACK_API_REVENUE_NOT_PASSED_FOR_REVENUE_GOAL")
		return false
	}

	variation, err := core.GetVariation(vwoInstance, userID, campaign)
	if err != nil {
		log.Error("No Variation Found")
		return false
	}

	impression := utils.CreateImpressionExtended(vwoInstance.SettingsFile, variation.ID, userID, campaign.ID, goal.ID, 5) //revenueValue
	if event.Dispatch(impression) {
		return true
	}

	log.Info("Ain't Keys For Impression")
	return false
}
