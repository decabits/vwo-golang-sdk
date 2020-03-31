package api

import (
	"log"

	"github.com/decabits/vwo-golang-sdk/lib/event"
	"github.com/decabits/vwo-golang-sdk/lib/constants"
	"github.com/decabits/vwo-golang-sdk/lib/core"
	"github.com/decabits/vwo-golang-sdk/lib/schema"
	"github.com/decabits/vwo-golang-sdk/lib/service"
	"github.com/decabits/vwo-golang-sdk/lib/utils"
)

//IsFeatureEnabled ...
func IsFeatureEnabled(vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options) bool {
	config := vwoInstance.Config
	settingsFileManager := service.SettingsFileManager{}
	vwoInstance.SettingsFile = settingsFileManager.GetSettingsFile()

	if options.CustomVariables == nil || options.VariationTargetingVariables == nil {
		return false
	}

	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		log.Println("Error geting campaign: ", err)
		return false
	}

	if campaign.Status != constants.StatusRunning {
		log.Println("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return false
	}
	if campaign.Type == constants.CampaignTypeVisualAB {
		log.Println("ERROR_MESSAGES.INVALID_API")
		return false
	}

	variation, err := core.GetVariation(config, userID, campaign, options)
	if err != nil {
		log.Println("No Variation Found")
		return false
	}

	if campaign.Type == constants.CampaignTypeFeatureTest {
		impression := utils.CreateImpression(vwoInstance.SettingsFile, campaign.ID, variation.ID, userID)
		if !event.Dispatch(impression) {
			return false
		}
		result := variation.IsFeatureEnabled
		if result {
			log.Println("Feature Enabled For User")
		} else {
			log.Println("Feature Not Enabled For User")
		}
		return result
	}
	return false
}
