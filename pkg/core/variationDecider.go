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

package core

import (
	"fmt"
	"strconv"

	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/utils"
)

const variationDecider = "variationDecider.go"

// VariationDecider struct
type VariationDecider struct {
	Bucketer         string
	SegmentEvaluator string
}

// GetVariation function
/*	Returns variation for the user for given campaign
    This method achieves the variation assignment in the following way:
    1. First get variation from UserStorage, if variation is found in user_storage_data,
    return from there
    2. Evaluates white listing users for each variation, and find a targeted variation.
    3. If no targeted variation is found, evaluate pre-segmentation result
    4. Evaluate percent traffic
    5. If user becomes part of campaign assign a variation.
	6. Store the variation found in the user_storage
*/
func GetVariation(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign, options schema.Options) (schema.Variation, error) {
	/*
		Args:
			userId: the unique ID assigned to User
			campaign: campaign in which user is participating
			customVariables(In option): variables for pre-segmentation
			variationTargetingVariables(In option): variables for variation targeting
			revenueGoal(In option): Value of revenue for the goal if the goal is revenue tracking

		Returns:
			schema.Variation: Struct object containing the information regarding variation assigned else empty object
			error: Error message
	*/

	_, ok := options.VariationTargetingVariables["_vwo_user_id"]
	if !ok {
		options.VariationTargetingVariables = map[string]interface{}{"_vwo_user_id": userID}
	}

	targettedVariation, err := FindTargetedVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		utils.LogMessage(vwoInstance.Logger, constants.Error, variationDecider, err.Error())
	} else {
		message := fmt.Sprintf(constants.DebugMessageGotVariation, userID, campaign.Key, targettedVariation.Name)
		utils.LogMessage(vwoInstance.Logger, constants.Info, variationDecider, message)
		return targettedVariation, nil
	}

	variationName, err := GetVariationFromUserStorage(vwoInstance, userID, campaign)
	if err != nil {
		utils.LogMessage(vwoInstance.Logger, constants.Error, variationDecider, err.Error())
	}
	if variationName != "" {
		message := fmt.Sprintf(constants.DebugMessageGettingStoredVariation, userID, campaign.Key, variationName)
		utils.LogMessage(vwoInstance.Logger, constants.Debug, variationDecider, message)
		return utils.GetCampaignVariation(campaign, variationName)
	}

	if EvaluateSegment(vwoInstance, campaign.Segments, options) && IsUserPartOfCampaign(vwoInstance, userID, campaign) {
		variation, err := BucketUserToVariation(vwoInstance, userID, campaign)
		if err != nil {
			return schema.Variation{}, fmt.Errorf(constants.DebugMessageNoVariationAllocated, userID, campaign.Key, campaign.Type, err.Error())
		}
		if vwoInstance.UserStorage != nil {
			if storage, ok := vwoInstance.UserStorage.(interface{ Set(a, b, c string) }); ok {
				storage.Set(userID, campaign.Key, variationName)
				message := fmt.Sprintf(constants.InfoMessageSettingDataUserStorageService, userID)
				utils.LogMessage(vwoInstance.Logger, constants.Info, variationDecider, message)
			}
			utils.LogMessage(vwoInstance.Logger, constants.Debug, variationDecider, constants.DebugMessagesNoUserStorageServiceSet)
		}

		message := fmt.Sprintf(constants.InfoMessageVariationAllocated, userID, campaign.Key, variation.Name)
		utils.LogMessage(vwoInstance.Logger, constants.Info, variationDecider, message)
		return variation, nil
	}

	return schema.Variation{}, fmt.Errorf(constants.ErrorMessageNoVariationAlloted, userID, campaign.Key, campaign.Type)
}

// FindTargetedVariation function Identifies and retrives if there exists any targeted
// variation in the given campaign for given userID
func FindTargetedVariation(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign, options schema.Options) (schema.Variation, error) {
	/*
		Args:
			userId: the unique ID assigned to User
			campaign: campaign in which user is participating
			customVariables(In option): variables for pre-segmentation
			variationTargetingVariables(In option): variables for variation targeting
			revenueGoal(In option): Value of revenue for the goal if the goal is revenue tracking

		Returns:
			schema.Variation: Struct object containing the information regarding variation assigned else empty object
			error: Error message
	*/

	if campaign.IsForcedVariation == false {
		return schema.Variation{}, fmt.Errorf(constants.DebugMessageWhitelistingSkipped, userID, campaign.Key)
	}
	whiteListedVariationsList := GetWhiteListedVariationsList(vwoInstance, userID, campaign, options)
	whiteListedVariationsLength := len(whiteListedVariationsList)
	var targettedVariation schema.Variation
	if whiteListedVariationsLength == 0 {
		return schema.Variation{}, fmt.Errorf(constants.DebugMessageNoWhitelistedVariation, campaign.Key)
	} else if whiteListedVariationsLength == 1 {
		targettedVariation = whiteListedVariationsList[0]
	} else {
		whiteListedVariationsList = utils.ScaleVariations(whiteListedVariationsList)
		whiteListedVariationsList = utils.GetVariationAllocationRanges(vwoInstance, whiteListedVariationsList)
		bucketValue := GetBucketValueForUser(vwoInstance, userID, constants.MaxTrafficValue, 1)
		var err error
		targettedVariation, err = GetBucketerVariation(vwoInstance, whiteListedVariationsList, bucketValue, userID, campaign.Key)
		if err != nil {
			return schema.Variation{}, fmt.Errorf(constants.DebugMessageNoTargettedVariation, err.Error())
		}
	}
	return targettedVariation, nil
}

// GetVariationFromUserStorage function tries retrieving variation from user_storage
func GetVariationFromUserStorage(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign) (string, error) {
	/*
		Args:
			userId: the unique ID assigned to User
			campaign: campaign in which user is participating

		Returns:
			variationName: Name of the found varaition in the user storage
			error: Error message
	*/

	if vwoInstance.UserStorage == nil {
		return "", fmt.Errorf(constants.DebugMessageNoUserStorageServiceGet)
	}
	if storage, ok := vwoInstance.UserStorage.(interface {
		Get(a, b string) schema.UserData
	}); ok {
		userStorageFetch := storage.Get(userID, campaign.Key)
		message := fmt.Sprintf(constants.DebugMessageGettingStoredVariation, userID, campaign.Key, userStorageFetch.VariationName)
		utils.LogMessage(vwoInstance.Logger, constants.Debug, variationDecider, message)
		if userStorageFetch.VariationName == "" {
			message := fmt.Sprintf(constants.DebugMessagesNoStoredVariation, userID, campaign.Key)
			utils.LogMessage(vwoInstance.Logger, constants.Debug, variationDecider, message)
		}
		return userStorageFetch.VariationName, nil
	}
	return "", fmt.Errorf(constants.ErrorMessagesGetUserStorageServiceFailed)
}

//GetWhiteListedVariationsList function identifies all forced variations which are targeted by variation_targeting_variables
func GetWhiteListedVariationsList(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign, options schema.Options) []schema.Variation {
	/*
		Args:
			userId: the unique ID assigned to User
			campaign: campaign in which user is participating
			customVariables(In option): variables for pre-segmentation
			variationTargetingVariables(In option): variables for variation targeting
			revenueGoal(In option): Value of revenue for the goal if the goal is revenue tracking

		Returns:
			schema.Variation: Struct object containing the information regarding variation assigned else empty object
	*/

	var whiteListedVariationsList []schema.Variation
	for _, variation := range campaign.Variations {
		if len(variation.Segments) == 0 {
			message := fmt.Sprintf(constants.DebugMessageNoSegmentsInVariation, userID, campaign.Key, variation)
			utils.LogMessage(vwoInstance.Logger, constants.Info, variationDecider, message)
		}
		status := PreEvaluateSegment(vwoInstance, variation.Segments, options)
		if status {
			whiteListedVariationsList = append(whiteListedVariationsList, variation)
		}

		message := fmt.Sprintf(constants.InfoMessageSegmentationStatus, userID, campaign.Key, options.CustomVariables, strconv.FormatBool(status), variation)
		utils.LogMessage(vwoInstance.Logger, constants.Info, variationDecider, message)
	}
	return whiteListedVariationsList
}

// EvaluateSegment function evaluates segmentation for the userID against the segments found inside the campaign.
func EvaluateSegment(vwoInstance schema.VwoInstance, segments map[string]interface{}, options schema.Options) bool {
	/*
		Args:
			segments: segments from campaign or variation
			options: options object containing CustomVariables, VariationTargertting variables and Revenue Goal

		Returns:
			bool: if the options falls in the segments criteria
	*/

	if len(segments) == 0 {
		message := fmt.Sprintf(constants.DebugMessageSegmentationSkipped, segments, options.CustomVariables)
		utils.LogMessage(vwoInstance.Logger, constants.Info, variationDecider, message)

		return true
	}
	return SegmentEvaluator(segments, options.CustomVariables)
}

// PreEvaluateSegment function evaluates segmentation for the userID against the segments found inside the campaign.
func PreEvaluateSegment(vwoInstance schema.VwoInstance, segments map[string]interface{}, options schema.Options) bool {
	/*
		Args:
			segments: segments from campaign or variation
			options: options object containing CustomVariables, VariationTargertting variables and Revenue Goal

		Returns:
			bool: if the options falls in the segments criteria
	*/

	if len(segments) == 0 {
		message := fmt.Sprintf(constants.DebugMessageSegmentationSkipped, segments, options.CustomVariables)
		utils.LogMessage(vwoInstance.Logger, constants.Info, variationDecider, message)

		return false
	}
	return SegmentEvaluator(segments, options.VariationTargetingVariables)
}