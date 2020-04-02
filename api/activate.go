package api

import (
	log "github.com/golang/glog"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Activate function
func Activate(vwoInstance schema.VwoInstance, campaignKey, userID string) string {
	options := schema.Options{}
	return ActivateWithOptions(vwoInstance, campaignKey, userID, options)
}

func ActivateWithOptions(vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options) string {
	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		log.Error("Error geting campaign: ", err)
		return ""
	}

	if campaign.Status != constants.StatusRunning {
		log.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return ""
	}
	if !utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		log.Error("ERROR_MESSAGES.INVALID_API")
		return ""
	}

	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		log.Error("No Variation Found")
		return ""
	}

	impression := utils.CreateImpression(vwoInstance.SettingsFile, campaign.ID, variation.ID, userID)
	if event.Dispatch(impression) {
		return variation.Name
	}

	log.Info("Main Keys For Impression")
	return ""
}
