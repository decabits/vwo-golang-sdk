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
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/utils"
)

const getFeatureVariableValue = "getFeatureVariableValue.go"

// GetFeatureVariableValue function
/*
This API method: Gets the value of the variable whose key is passed
1. Validates the arguments being passed
2. Finds the corresponding Campaign
3. Checks the Campaign Status
4. Validates the Campaign Type
5. Assigns the determinitic variation to the user(based on userId), if user becomes part of campaign
   If userStorageService is used, it will look into it for the variation and if found, no further processing is done
6. Gets the value of the variable depeneding upon the type of the campaign
7. Logs and returns the value
*/
func (vwo *VWOInstance) GetFeatureVariableValue(campaignKey, variableKey, userID string, option interface{}) interface{} {
	/*
		Args:
			campaignKey: Key of the running campaign
			variableKey: Key of variable whose value is to be found
			userID: Unique identification of user
			customVariables(In option): variables for pre-segmentation
			variationTargetingVariables(In option): variables for variation targeting
			revenueGoal(In option): Value of revenue for the goal if the goal is revenue tracking

		Returns:
			interrface{}: Value of the variable
	*/
	if !utils.ValidateGetFeatureVariableValue(campaignKey, variableKey, userID) {
		message := fmt.Sprintf(constants.ErrorMessagesGetFeatureVariableMissingParams)
		utils.LogMessage(vwo.Logger, constants.Error, getFeatureVariableValue, message)
		return nil
	}

	options := utils.ParseOptions(option)

	campaign, err := utils.GetCampaign(vwo.SettingsFile, campaignKey)
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessageCampaignNotFound+" \n %v", campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Error, getFeatureVariableValue, message)
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
		UserID:            userID,
		Campaign:          campaign,
	}
	variation, err := core.GetVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		message := fmt.Sprintf(constants.InfoMessageInvalidVariationKey+" \n %v", userID, campaignKey, err.Error())
		utils.LogMessage(vwo.Logger, constants.Info, getFeatureVariableValue, message)
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