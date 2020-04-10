package core

import (
	"errors"

	"strconv"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

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
			assigned else None
			error(error): Error message
	*/
	options.VWOUserID = userID

	targettedVariation, err := FindTargetedVariation(vwoInstance, userID, campaign, options)
	if err != nil {
		vwoInstance.Logger.Error(err)
	} else {
		vwoInstance.Logger.Info("INFO_MESSAGES.GOT_VARIATION_FOR_USER", targettedVariation)
		return targettedVariation, nil
	}

	variationName, err := GetVariationFromUserStorage(vwoInstance, userID, campaign)
	if err != nil {
		vwoInstance.Logger.Error(err)
	}
	if variationName != "" {
		vwoInstance.Logger.Info("DEBUG_MESSAGES.GETTING_STORED_VARIATION")
		return utils.GetCampaignVariation(campaign, variationName)
	}

	if EvaluateSegment(vwoInstance, campaign.Segments, options) && IsUserPartOfCampaign(vwoInstance, userID, campaign) {
		variation, err := BucketUserToVariation(vwoInstance, userID, campaign)
		if err != nil {
			vwoInstance.Logger.Info("DEBUG_MESSAGES.VARIATION_NOT_FOUND")
			return schema.Variation{}, nil
		}
		if vwoInstance.UserStorage.Exist() {
			vwoInstance.UserStorage.Set(userID, campaign.Key, variationName)
		}
		vwoInstance.Logger.Info("INFO_MESSAGES.GOT_VARIATION_FOR_USER", variation)
		return variation, nil
	}

	return schema.Variation{}, nil
}

// FindTargetedVariation function
func FindTargetedVariation(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign, options schema.Options) (schema.Variation, error) {
	if campaign.IsForcedVariation == false {
		return schema.Variation{}, errors.New("DEBUG_MESSAGES.WHITELISTING_SKIPPED")
	}
	whiteListedVariationsList := GetWhiteListedVariationsList(vwoInstance, userID, campaign, options)
	whiteListedVariationsLength := len(whiteListedVariationsList)
	var targettedVariation schema.Variation
	if whiteListedVariationsLength == 0 {
		return schema.Variation{}, errors.New("No White listed variation")
	} else if whiteListedVariationsLength == 1 {
		targettedVariation = whiteListedVariationsList[0]
	} else {
		whiteListedVariationsList = utils.ScaleVariations(whiteListedVariationsList)
		whiteListedVariationsList = utils.GetVariationAllocationRanges(vwoInstance, whiteListedVariationsList)
		bucketValue := GetBucketValueForUser(vwoInstance, userID, constants.MaxTrafficValue, 1)
		var err error
		targettedVariation, err = GetBucketerVariation(whiteListedVariationsList, bucketValue)
		if err != nil {
			return schema.Variation{}, errors.New("No targetted variation: " + err.Error())
		}
	}
	return targettedVariation, nil
}

// GetVariationFromUserStorage ...
func GetVariationFromUserStorage(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign) (string, error) {
	if !vwoInstance.UserStorage.Exist() {
		return "", errors.New("DEBUG_MESSAGES.NO_USER_STORAGE_SERVICE_GET")
	}
	userStorageFetch := vwoInstance.UserStorage.Get(userID, campaign.Key)

	vwoInstance.Logger.Info("INFO_MESSAGES.GETTING_DATA_USER_STORAGE_SERVICE")
	return userStorageFetch.VariationName, nil
}

//GetWhiteListedVariationsList ...
func GetWhiteListedVariationsList(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign, options schema.Options) []schema.Variation {
	// check Validity of
	var whiteListedVariationsList []schema.Variation
	for _, variation := range campaign.Variations {
		if len(variation.Segments) == 0 {
			vwoInstance.Logger.Info("DEBUG_MESSAGES.SEGMENTATION_SKIPPED")
		}
		status := PreEvaluateSegment(vwoInstance, variation.Segments, options)
		if status {
			whiteListedVariationsList = append(whiteListedVariationsList, variation)
		}
		vwoInstance.Logger.Info("DEBUG_MESSAGES.SEGMENTATION_STATUS " + strconv.FormatBool(status))
	}
	return whiteListedVariationsList
}

// EvaluateSegment function
func EvaluateSegment(vwoInstance schema.VwoInstance, segments map[string]interface{}, options schema.Options) bool {
	if len(segments) == 0 {
		vwoInstance.Logger.Info("DEBUG_MESSAGES.SEGMENTATION_SKIPPED")
		return true
	}
	return SegmentEvaluator(segments, options)
}

// PreEvaluateSegment ...
func PreEvaluateSegment(vwoInstance schema.VwoInstance, segments map[string]interface{}, options schema.Options) bool {
	if len(segments) == 0 {
		vwoInstance.Logger.Info("DEBUG_MESSAGES.SEGMENTATION_SKIPPED")
		return false
	}
	return SegmentEvaluator(segments, options)
}
