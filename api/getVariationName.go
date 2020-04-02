package api

import (
	log "github.com/golang/glog"

	"github.com/Piyushhbhutoria/vwo-go-sdk/constants"
	"github.com/Piyushhbhutoria/vwo-go-sdk/core"
	"github.com/Piyushhbhutoria/vwo-go-sdk/schema"
	"github.com/Piyushhbhutoria/vwo-go-sdk/service"
	"github.com/Piyushhbhutoria/vwo-go-sdk/utils"
)

//GetVariationName ...
func GetVariationName(vwoInstance schema.VwoInstance, campaignKey, userID string, options schema.Options) string {
	settingsFileManager := service.SettingsFileManager{}
	vwoInstance.SettingsFile = settingsFileManager.GetSettingsFile()

	if options.CustomVariables == nil || options.VariationTargetingVariables == nil {
		return ""
	}

	campaign, err := utils.GetCampaign(vwoInstance.SettingsFile, campaignKey)
	if err != nil {
		log.Error("Error geting campaign: ", err)
		return ""
	}

	if campaign.Status != constants.StatusRunning {
		log.Error("ERROR_MESSAGES.CAMPAIGN_NOT_RUNNING")
		return ""
	}
	if campaign.Type != constants.CampaignTypeVisualAB {
		log.Error("ERROR_MESSAGES.INVALID_API")
		return ""
	}

	variation, err := core.GetVariation(vwoInstance, userID, campaign)
	if err != nil {
		log.Error("No Variation Found")
		return ""
	}

	return variation.Name
}
