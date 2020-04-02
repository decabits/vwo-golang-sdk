package api

import (
	log "github.com/golang/glog"

	"github.com/Piyushhbhutoria/vwo-go-sdk/constants"
	"github.com/Piyushhbhutoria/vwo-go-sdk/core"
	"github.com/Piyushhbhutoria/vwo-go-sdk/event"
	"github.com/Piyushhbhutoria/vwo-go-sdk/schema"
	"github.com/Piyushhbhutoria/vwo-go-sdk/service"
	"github.com/Piyushhbhutoria/vwo-go-sdk/utils"
)

//IsFeatureEnabled ...
func IsFeatureEnabled(vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options) bool {
	settingsFileManager := service.SettingsFileManager{}
	vwoInstance.SettingsFile = settingsFileManager.GetSettingsFile()

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
	if campaign.Type == constants.CampaignTypeVisualAB {
		log.Error("ERROR_MESSAGES.INVALID_API")
		return false
	}

	variation, err := core.GetVariation(vwoInstance, userID, campaign)
	if err != nil {
		log.Error("No Variation Found")
		return false
	}

	if campaign.Type == constants.CampaignTypeFeatureTest {
		impression := utils.CreateImpression(vwoInstance.SettingsFile, campaign.ID, variation.ID, userID)
		if !event.Dispatch(impression) {
			return false
		}
		result := variation.IsFeatureEnabled
		if result {
			log.Info("Feature Enabled For User")
		} else {
			log.Info("Feature Not Enabled For User")
		}
		return result
	}
	return false
}
