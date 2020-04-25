package vwo

import (
	"fmt"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

const getVariationName = "getVariationName.go"

// GetVariationName function
/*
This API method: Gets the variation assigned for the user for the campaign
1. Validates the arguments being passed
2. Finds the corresponding Campaign
3. Checks the Campaign Status
4. Validates the Campaign Type
5. Assigns the determinitic variation to the user(based on userId), if user becomes part of campaign
   If userStorageService is used, it will look into it for the variation and if found, no further processing is done
*/
func (vwo *VWOInstance) GetVariationName(campaignKey, userID string) string {
	/*
		Args:
			campaignKey: Key of the running campaign 
			userID: Unique identification of user
		Returns:
			string: Variation Name for user to corresponding camapign
	*/
	options := schema.Options{}
	return vwo.GetVariationNameWithOptions(campaignKey, userID, options)
}

// GetVariationNameWithOptions ...
func (vwo *VWOInstance) GetVariationNameWithOptions(campaignKey, userID string, options schema.Options) string {
	/*
		Args:
			campaignKey: Key of the running campaign 
			userID: Unique identification of user
			customVariables(In schema.Options): variables for pre-segmentation, pass it through **kwargs as
			customVariables = {}
			variationTargetingVariables(In schema.Options): variables for variation targeting, pass it through **kwargs as
			variationTargetingVariables = {}
		Returns:
			string: Variation Name for user to corresponding camapign
	*/
	if !utils.ValidateGetVariationName(campaignKey, userID) {
		message := fmt.Sprintf(constants.ErrorMessagesGetVariationAPIMissingParams)
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, message)
		return ""
	}

	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotFound+" \n %v", campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, message)
		return ""
	}

	if campaign.Status != constants.StatusRunning {
		message := fmt.Sprintf(constants.ErrorMessagesCampaignNotRunning, "GetVariationName", campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, message)
		return ""
	}
	if !utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		message := fmt.Sprintf(constants.ErrorMessagesInvalidAPI, "GetVariationName", campaignKey, campaign.Type, userID)
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, message)
		return ""
	}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      vwo.SettingsFile,
		UserStorage:       vwo.UserStorage,
		Logger:            vwo.Logger,
		IsDevelopmentMode: vwo.IsDevelopmentMode,
		UserID:            userID,
		Campaign:          campaign,
	}
	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		message := fmt.Sprintf(constants.InfoMessageInvalidVariationKey+" \n %v", userID, campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Info, getVariationName, message)
	}

	return variation.Name
}
