/*
   Copyright 2020 Wingify Software Pvt. Ltd.

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
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/utils"
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
func (vwo *VWOInstance) GetVariationName(campaignKey, userID string, option interface{}) string {
	/*
		Args:
			campaignKey: Key of the running campaign
			userID: Unique identification of user
			customVariables(In option): variables for pre-segmentation
			variationTargetingVariables(In option): variables for variation targeting
			revenueValue(In option): Value of revenue for the goal if the goal is revenue tracking

		Returns:
			string: Variation Name for user to corresponding camapign
	*/

	vwoInstance := schema.VwoInstance{
		SettingsFile:      vwo.SettingsFile,
		UserStorage:       vwo.UserStorage,
		Logger:            vwo.Logger,
		IsDevelopmentMode: vwo.IsDevelopmentMode,
		API:               "GetVariationName",
	}

	if !utils.ValidateGetVariationName(campaignKey, userID) {
		message := fmt.Sprintf(constants.ErrorMessageGetVariationAPIMissingParams, vwoInstance.API)
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, message)
		return ""
	}

	options := utils.ParseOptions(option)

	campaign, err := utils.GetCampaign(vwoInstance.API, vwo.SettingsFile, campaignKey)
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotFound, vwoInstance.API, campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, message)
		return ""
	}

	if campaign.Status != constants.StatusRunning {
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotRunning, vwoInstance.API, campaignKey)
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, message)
		return ""
	}
	if utils.CheckCampaignType(campaign, constants.CampaignTypeFeatureRollout) {
		message := fmt.Sprintf(constants.ErrorMessageInvalidAPI, vwoInstance.API, campaignKey, campaign.Type, userID)
		utils.LogMessage(vwo.Logger, constants.Error, getVariationName, message)
		return ""
	}

	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		message := fmt.Sprintf(constants.InfoMessageInvalidVariationKey, vwoInstance.API, userID, campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Info, getVariationName, message)
	}

	return variation.Name
}
