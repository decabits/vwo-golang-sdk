/*
   Copyright 2019-2020 Wingify Software Pvt. Ltd.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package api

import (
	"fmt"

	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/core"
	"github.com/decabits/vwo-golang-sdk/pkg/event"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/utils"
)

const fileIsFeatureEnabled = "isFeatureEnabled.go"

// IsFeatureEnabled function
/*
This API method: Whether a feature is enabled or not for the given user
1. Validates the arguments being passed
2. Finds the corresponding Campaign
3. Checks the Campaign Status
4. Validates the Campaign Type
5. Assigns the determinitic variation to the user(based on userId), if user becomes part of campaign
   If userStorageService is used, it will look into it for the variation and if found, no further processing is done
6. If feature enabled, sends a call to VWO server for tracking visitor
*/
func (vwo *VWOInstance) IsFeatureEnabled(campaignKey, userID string, option interface{}) bool {
	/*
		Args:
			campaignKey: Key of the running campaign
			userID: Unique identification of user
			customVariables(In option): variables for pre-segmentation
			variationTargetingVariables(In option): variables for variation targeting
			revenueGoal(In option): Value of revenue for the goal if the goal is revenue tracking

		Returns:
			bool: True if the user the feature is enambled for the user, else false
	*/
	if !utils.ValidateIsFeatureEnabled(campaignKey, userID) {
		message := fmt.Sprintf(constants.ErrorMessagesIsFeatureEnabledAPIMissingParams)
		utils.LogMessage(vwo.Logger, constants.Error, fileIsFeatureEnabled, message)
		return false
	}

	options := utils.ParseOptions(option)

	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotFound+" \n %v", campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Error, fileIsFeatureEnabled, message)
		return false
	}

	if campaign.Status != constants.StatusRunning {
		message := fmt.Sprintf(constants.ErrorMessagesCampaignNotRunning, "IsFeatureEnabled", campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, fileIsFeatureEnabled, message)
		return false
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		message := fmt.Sprintf(constants.ErrorMessagesInvalidAPI, "IsFeatureEnabled", campaignKey, campaign.Type, userID)
		utils.LogMessage(vwo.Logger, constants.Error, fileIsFeatureEnabled, message)
		return false
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
		utils.LogMessage(vwo.Logger, constants.Info, fileIsFeatureEnabled, message)
		return false
	}

	isFeatureEnabled := false
	if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureTest) {
		isFeatureEnabled = variation.IsFeatureEnabled
		impression := utils.CreateImpressionTrackingUser(vwoInstance, campaign.ID, variation.ID, userID)
		event.Dispatch(vwoInstance, impression)
	} else if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		isFeatureEnabled = true
	}

	if isFeatureEnabled {
		message := fmt.Sprintf(constants.InfoMessageFeatureEnabledForUser, campaignKey, userID)
		utils.LogMessage(vwo.Logger, constants.Info, fileIsFeatureEnabled, message)
	} else {
		message := fmt.Sprintf(constants.InfoMessageFeatureNotEnabledForUser, campaignKey, userID)
		utils.LogMessage(vwo.Logger, constants.Info, fileIsFeatureEnabled, message)
	}

	return isFeatureEnabled
}