package core

import (
	"fmt"
	"strconv"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
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
			userId (string): the unique ID assigned to User
			campaign (dischema.Campaignct): campaign in which user is participating
			customVariables(In schema.Options): variables for pre-segmentation, pass it through **kwargs as
			customVariables = {}
			variationTargetingVariables(In schema.Options): variables for variation targeting, pass it through **kwargs as
			variationTargetingVariables = {}
		Returns:
			variation (schema.Variation): Dict object containing the information regarding variation
			assigned else empty object
			error(error): Error message
	*/
	_, ok := options.VariationTargetingVariables["_vwo_user_id"]
	if !ok {
		options.VariationTargetingVariables = map[string]interface{}{"_vwo_user_id": userID}
	}


	targettedVariation, err := FindTargetedVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		utils.LogMessage(vwoInstance,constants.Error, variationDecider, err.Error())
	} else {
		message := fmt.Sprintf(constants.InfoMessageGotVariation, userID, campaign.Key, targettedVariation.Name)
		utils.LogMessage(vwoInstance, constants.Info, variationDecider, message)
		return targettedVariation, nil
	}

	variationName, err := GetVariationFromUserStorage(vwoInstance, userID, campaign)
	if err != nil {
		utils.LogMessage(vwoInstance,constants.Error, variationDecider, err.Error())
	}
	if variationName != "" {
		message := fmt.Sprintf(constants.InfoMessageGettingStoredVariation, userID, campaign.Key, variationName)
		utils.LogMessage(vwoInstance, constants.Info, variationDecider, message)
		return utils.GetCampaignVariation(campaign, variationName)
	}

	if EvaluateSegment(vwoInstance, campaign.Segments, options) && IsUserPartOfCampaign(vwoInstance, userID, campaign) {
		variation, err := BucketUserToVariation(vwoInstance, userID, campaign)
		if err != nil {
			return schema.Variation{}, fmt.Errorf(constants.InfoMessageNoVariationAllocated, userID, campaign.Key, campaign.Type, err.Error())
		}
		if vwoInstance.UserStorage.Exist() {
			vwoInstance.UserStorage.Set(userID, campaign.Key, variationName)
			message := fmt.Sprintf(constants.InfoMessageSettingDataUserStorageService, userID)
			utils.LogMessage(vwoInstance, constants.Info, variationDecider, message)
		}

		message := fmt.Sprintf(constants.InfoMessageGotVariationForUser, userID, campaign.Key, variation.Name, "GetVariation")
		utils.LogMessage(vwoInstance, constants.Info, variationDecider, message)
		return variation, nil
	}

	return schema.Variation{}, fmt.Errorf(constants.ErrorMessageNoVariationAlloted, userID, campaign.Key, campaign.Type)
}

// FindTargetedVariation function Identifies and retrives if there exists any targeted variation in the given campaign for given userID
func FindTargetedVariation(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign, options schema.Options) (schema.Variation, error) {
	/*
		Args:
			userId (string): the unique ID assigned to User
			campaign (dischema.Campaignct): campaign in which user is participating
			customVariables(In schema.Options): variables for pre-segmentation, pass it through **kwargs as
			customVariables = {}
			variationTargetingVariables(In schema.Options): variables for variation targeting, pass it through **kwargs as
			variationTargetingVariables = {}
		Returns:
			variation (schema.Variation): Dict object containing the information regarding variation
			assigned else empty object
			error(error): Error message
	*/

	if campaign.IsForcedVariation == false {
		return schema.Variation{}, fmt.Errorf(constants.InfoMessageWhitelistingSkipped, userID, campaign.Key)
	}
	whiteListedVariationsList := GetWhiteListedVariationsList(vwoInstance, userID, campaign, options)
	whiteListedVariationsLength := len(whiteListedVariationsList)
	var targettedVariation schema.Variation
	if whiteListedVariationsLength == 0 {
		return schema.Variation{}, fmt.Errorf(constants.InfoMessageNoWhitelistedVariation, campaign.Key)
	} else if whiteListedVariationsLength == 1 {
		targettedVariation = whiteListedVariationsList[0]
	} else {
		whiteListedVariationsList = utils.ScaleVariations(whiteListedVariationsList)
		whiteListedVariationsList = utils.GetVariationAllocationRanges(vwoInstance, whiteListedVariationsList)
		bucketValue := GetBucketValueForUser(vwoInstance, userID, constants.MaxTrafficValue, 1)
		var err error
		targettedVariation, err = GetBucketerVariation(whiteListedVariationsList, bucketValue)
		if err != nil {
			return schema.Variation{}, fmt.Errorf(constants.InfoMessageNoTargettedVariation, err.Error())
		}
	}
	return targettedVariation, nil
}

// GetVariationFromUserStorage function tries retrieving variation from user_storage
func GetVariationFromUserStorage(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign) (string, error) {
	/*
		Args:
			userId (string): the unique ID assigned to User
			campaign (dischema.Campaignct): campaign in which user is participating

		Returns:
			variationName: Name of the found varaition in the user storage
			error: Error message
	*/

	if !vwoInstance.UserStorage.Exist() {
		return "", fmt.Errorf(constants.InfoMessageNoUserStorageServiceGet)
	}
	userStorageFetch := vwoInstance.UserStorage.Get(userID, campaign.Key)

	message := fmt.Sprintf(constants.InfoMessageGettingDataUserStorageService, userID)
	utils.LogMessage(vwoInstance, constants.Info, variationDecider, message)
	return userStorageFetch.VariationName, nil
}

//GetWhiteListedVariationsList function identifies all forced variations which are targeted by variation_targeting_variables
func GetWhiteListedVariationsList(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign, options schema.Options) []schema.Variation {
	/*
		Args:
			userId (string): the unique ID assigned to User
			campaign (schema.Campaign): campaign in which user is participating
			customVariables(In schema.Options): variables for pre-segmentation, pass it through **kwargs as
			customVariables = {}
			variationTargetingVariables(In schema.Options): variables for variation targeting, pass it through **kwargs as
			variationTargetingVariables = {}
		Returns:
			variation (schema.Variation): Dict object containing the information regarding variation
			assigned else empty object
	*/

	var whiteListedVariationsList []schema.Variation
	for _, variation := range campaign.Variations {
		if len(variation.Segments) == 0 {
			message := fmt.Sprintf(constants.InfoMessageNoSegmentsInVariation, userID, campaign.Key, variation)
			utils.LogMessage(vwoInstance, constants.Info, variationDecider, message)
		}
		status := PreEvaluateSegment(vwoInstance, variation.Segments, options)
		if status {
			whiteListedVariationsList = append(whiteListedVariationsList, variation)
		}

		message := fmt.Sprintf(constants.InfoMessageSegmentationStatus, userID, campaign.Key, options.CustomVariables, strconv.FormatBool(status), variation)
		utils.LogMessage(vwoInstance, constants.Info, variationDecider, message)
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
		message := fmt.Sprintf(constants.InfoMessageSegmentationSkipped, segments, options.CustomVariables)
		utils.LogMessage(vwoInstance, constants.Info, variationDecider, message)
		vwoInstance.Logger.Info()

		return true
	}
	return SegmentEvaluator(segments, options)
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
		message := fmt.Sprintf(constants.InfoMessageSegmentationSkipped, segments, options.CustomVariables)
		utils.LogMessage(vwoInstance, constants.Info, variationDecider, message)
		vwoInstance.Logger.Info()

		return false
	}
	return SegmentEvaluator(segments, options)
}
