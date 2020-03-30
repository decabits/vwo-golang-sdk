package api

import (
	"log"

	"github.com/decabits/vwo-golang-sdk/lib/constants"
	"github.com/decabits/vwo-golang-sdk/lib/core"
	"github.com/decabits/vwo-golang-sdk/lib/schema"
	"github.com/decabits/vwo-golang-sdk/lib/service"
	"github.com/decabits/vwo-golang-sdk/lib/utils"
)

//Track ...
func Track(config schema.Config, vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options, goalIdentifier string) bool {
	if options.CustomVariables == nil || options.VariationTargetingVariables == nil {
		return false
	}

	settingsFileManager := service.SettingsFileManager{}
	vwoInstance.SettingsFile = settingsFileManager.GetSettingsFile()

	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		log.Println("Error geting campaign: ", err)
		return false
	}

	if campaign.Status != constants.StatusRunning {
		log.Println("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return false
	}
	if campaign.Type != constants.CampaignTypeFeatureRollout {
		log.Println("ERROR_MESSAGES.INVALID_API")
		return false
	}

	goal, err = utils.GetCampaignGoal(campaign, goalIdentifier)
	if err != nil {
		return false
	}
	revenueValue = options.RevenueGoal
	if goal.Type == constants.GoalTypeRevenue && revenueGoal > 0 {
		log.Println("ERROR_MESSAGES.TRACK_API_REVENUE_NOT_PASSED_FOR_REVENUE_GOAL")
		return false
	}

	variation, err := core.GetVariation(config, userID, campaign, options)
	if err != nil {
		log.Println("No Variation Found")
		return false
	}// Segmentation issue in VarDecider 

	impression := utils.CreateImpressionExtended(vwoInstance.SettingsFile, variation.ID, userID, campaign.ID, goal.ID, revenueValue)
	if event.Dispatch(impression) {
		return true
	}// Gsearch Url with params 

	log.Println("ain't Keys For Impression")
	return false
}
