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

package vwo

import (
	"fmt"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/core"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

const activate = "activate.go"

// Activate function
/*
This API method: Gets the variation assigned for the user for the campaign and send the metrics to VWO server
1. Validates the arguments being passed
2. Finds the corresponding Campaign
3. Checks the Campaign Status
4. Validates the Campaign Type
5. Assigns the determinitic variation to the user(based on userId), if user becomes part of campaign
   If userStorageService is used, it will look into it for the variation and if found, no further processing is done
6. Sends an impression call to VWO server to track user
*/
func (vwo *VWOInstance) Activate(campaignKey, userID string, option interface{}) string {
	/*
		Args:
			campaignKey: Key of the running campaign
			userID: Unique identification of user
			customVariables(In option): variables for pre-segmentation
			variationTargetingVariables(In option): variables for variation targeting
			revenueGoal(In option): Value of revenue for the goal if the goal is revenue tracking
			
		Returns:
			string: Variation Name for user to corresponding camapign
	*/

	options := utils.ParseOptions(option)

	vwoInstance := schema.VwoInstance{
		SettingsFile:      vwo.SettingsFile,
		UserStorage:       vwo.UserStorage,
		Logger:            vwo.Logger,
		IsDevelopmentMode: vwo.IsDevelopmentMode,
		UserID:            userID,
	}

	if !utils.ValidateActivate(campaignKey, userID) {
		message := fmt.Sprintf(constants.ErrorMessagesActivateAPIMissingParams)
		utils.LogMessage(vwo.Logger, constants.Error, activate, message)
		return ""
	}

	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotFound+" \n %v", campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Error, activate, message)
		return ""
	}

	if campaign.Status != constants.StatusRunning {
		message := fmt.Sprintf(constants.ErrorMessagesCampaignNotRunning, "Activate", campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, activate, message)
		return ""
	}
	if !utils.CheckCampaignType(campaign, constants.CampaignTypeVisualAB) {
		message := fmt.Sprintf(constants.ErrorMessagesInvalidAPI, "Activate", campaignKey, campaign.Type, userID)
		utils.LogMessage(vwo.Logger, constants.Error, activate, message)
		return ""
	}

	vwoInstance.Campaign = campaign

	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		message := fmt.Sprintf(constants.InfoMessageInvalidVariationKey+" \n %v", userID, campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Info, activate, message)
		return ""
	}

	impression := utils.CreateImpressionTrackingUser(vwoInstance, campaign.ID, variation.ID, userID)
	event.Dispatch(vwoInstance, impression)

	return variation.Name
}
