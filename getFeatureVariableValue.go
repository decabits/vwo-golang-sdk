package vwo

import (
	"fmt"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

const getFeatureVariableValue = "getFeatureVariableValue.go"

// GetFeatureVariableValue function
func (vwo *VWOInstance) GetFeatureVariableValue(campaignKey, variableKey, userID string) interface{} {
	options := schema.Options{}
	return vwo.GetFeatureVariableValueWithOptions(campaignKey, variableKey, userID, options)
}

// GetFeatureVariableValueWithOptions function
func (vwo *VWOInstance) GetFeatureVariableValueWithOptions(campaignKey, variableKey, userID string, options schema.Options) interface{} {
	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotFound, campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, getFeatureVariableValue, message)
		utils.LogMessage(vwo.Logger, constants.Error, getFeatureVariableValue, err.Error())
		return nil
	}

	if campaign.Status != constants.StatusRunning {
		message := fmt.Sprintf(constants.ErrorMessagesCampaignNotRunning, "GetFeatureVariableValue", campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, getFeatureVariableValue, message)
		return nil
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		message := fmt.Sprintf(constants.ErrorMessagesInvalidAPI, "GetFeatureVariableValue", campaignKey, campaign.Type, userID)
		utils.LogMessage(vwo.Logger, constants.Error, getFeatureVariableValue, message)
		return nil
	}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      vwo.SettingsFile,
		UserStorage:       vwo.UserStorage,
		Logger:            vwo.Logger,
		IsDevelopmentMode: vwo.IsDevelopmentMode,
	}
	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		message := fmt.Sprintf(constants.InfoMessageInvalidVariationKey, userID, campaignKey)
		utils.LogMessage(vwo.Logger, constants.Info, getFeatureVariableValue, message)
		utils.LogMessage(vwo.Logger, constants.Error, getFeatureVariableValue, err.Error())
		return nil
	}

	var variable schema.Variable
	if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		variable = utils.GetVariableForFeature(campaign.Variables, variableKey)
	} else if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureTest) {
		variable = utils.GetVariableValueForVariation(vwoInstance, campaign, variation, variableKey, userID)
	}

	if variable.Key == "" {
		message := fmt.Sprintf(constants.ErrorMessagesvariableNotFound, variable.Key, userID)
		utils.LogMessage(vwo.Logger, constants.Error, getFeatureVariableValue, message)
	} else {
		message := fmt.Sprintf(constants.InfoMessageUserRecievedVariableValue, variable.Key, campaignKey, variable, userID)
		utils.LogMessage(vwo.Logger, constants.Error, getFeatureVariableValue, message)
	}

	return variable.Value
}
