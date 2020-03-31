package api

import (
	"log"

	"github.com/decabits/vwo-golang-sdk/lib/constants"
	"github.com/decabits/vwo-golang-sdk/lib/core"
	"github.com/decabits/vwo-golang-sdk/lib/schema"
	"github.com/decabits/vwo-golang-sdk/lib/service"
	"github.com/decabits/vwo-golang-sdk/lib/utils"
)

// ActivateWithOptions function
func ActivateWithOptions(config schema.Config, vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options) string {
	if options.CustomVariables == nil || options.VariationTargetingVariables == nil {
		return ""
	}

	settingsFileManager := service.SettingsFileManager{}
	vwoInstance.SettingsFile = settingsFileManager.GetSettingsFile()

	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		log.Println("Error geting campaign: ", err)
		return ""
	}

	if campaign.Status != constants.StatusRunning {
		log.Println("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return ""
	}
	if campaign.Type != constants.CampaignTypeVisualAB {
		log.Println("ERROR_MESSAGES.INVALID_API")
		return ""
	}

	variation, err := core.GetVariation(config, userID, campaign, options)
	if err != nil {
		log.Println("No Variation Found")
		return ""
	}// Segmentation issue in VarDecider 

	impression := utils.CreateImpression(vwoInstance.SettingsFile, campaign.ID, variation.ID, userID) //TO BE COMPLETED
	if event.Dispatch(impression) {
		return variation.Name
	} // Gsearch Url with params 

	//log.Println("ain't Keys For Impression")
	return ""
}
