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
func ActivateWithOptions(vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options) *string {
	if options.CustomVariables == nil || options.VariationTargetingVariables == nil {
		return nil
	}

	settingsFileManager := service.SettingsFileManager{}
	vwoInstance.SettingsFile = settingsFileManager.GetSettingsFile()

	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		log.Println("Error geting campaign: ", err)
		return nil
	}

	if campaign.Status != constants.StatusRunning {
		log.Println("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return nil
	}
	if campaign.Type != constants.CampaignTypeVisualAB {
		log.Println("ERROR_MESSAGES.INVALID_API")
		return nil
	}

	variation, err := core.GetVariation(userID, campaign, options)
	if err != nil {
		log.Println("No Variation Found")
		return nil
	}

	// impression := utils.CreateImpression(vwoInstance.SettingsFile, campaign.ID, variation.ID, userID) //TO BE COMPLETED
	// event.dispatch(impression)                                                                        //TO BE COMPLETED

	// log.Println("ain Keys For Impression")

	return variation.Name
}
