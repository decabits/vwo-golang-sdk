package core

//TO BE COMPLETED
import (
	"errors"
	"math/rand"
	"strconv"

	log "github.com/golang/glog"

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
	targettedVariation, err := FindTargetedVariation(userID, campaign, options)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("INFO_MESSAGES.GOT_VARIATION_FOR_USER")
		return targettedVariation, nil
	}

	variationName, err := GetVariationFromUserStorage(vwoInstance.UserStorage, userID, campaign)
	if err != nil {
		return schema.Variation{}, err
	}

	log.Info("DEBUG_MESSAGES.GETTING_STORED_VARIATION")
	return utils.GetCampaignVariation(campaign, variationName)
}

//GetVariationOfCampiagnForUser ...
func GetVariationOfCampiagnForUser(userID string, campaign schema.Campaign) (schema.VariationResponse, error) {
	variation, err := BucketUserToVariation(userID, campaign)

	if err != nil {
		return schema.VariationResponse{}, errors.New("Variation Not Found")
	}
	var Response schema.VariationResponse
	Response.Variation = variation
	Response.VariationName = variation.Name
	Response.VariationID = variation.ID
	return Response, nil

}

// FindTargetedVariation function
func FindTargetedVariation(userID string, campaign schema.Campaign, options schema.Options) (schema.Variation, error) {
	if campaign.IsForcedVariation == false {
		return schema.Variation{}, errors.New("DEBUG_MESSAGES.WHITELISTING_SKIPPED")
	}
	whiteListedVariationsList := GetWhiteListedVariationsList(userID, campaign, options)
	whiteListedVariationsLength := len(whiteListedVariationsList)
	var targettedVariation schema.Variation
	if whiteListedVariationsLength == 0 {
		return schema.Variation{}, errors.New("No White listed variation")
	} else if whiteListedVariationsLength == 1 {
		targettedVariation = whiteListedVariationsList[0]
	} else {
		whiteListedVariationsList = utils.ScaleVariations(whiteListedVariationsList)
		variationAllocation := utils.GetVariationAllocationRanges(whiteListedVariationsList)
		for i, variation := range whiteListedVariationsList {
			variation.StartVariationAllocation = variationAllocation[i].StartVariationAllocation
			variation.EndVariationAllocation = variationAllocation[i].EndVariationAllocation
		}
		bucketValue := GetBucketValueForUser(userID, constants.MaxTrafficValue)
		var err error
		targettedVariation, err = GetBucketerVariation(whiteListedVariationsList, bucketValue)
		if err != nil {
			return schema.Variation{}, errors.New("No targetted variation")
		}
	}
	return targettedVariation, nil
}

//GetVariationFromUserStorage ...
func GetVariationFromUserStorage(UserStorage schema.UserStorage, userID string, campaign schema.Campaign) (string, error) {
	if !UserStorage.Exist() {
		return "", errors.New("DEBUG_MESSAGES.NO_USER_STORAGE_SERVICE_GET")
	}
	userStorageFetch, err := UserStorage.Get(userID, campaign.Key)
	if err != nil {
		return "", errors.New("ERROR_MESSAGES.GET_USER_STORAGE_SERVICE_FAILED")
	}
	log.Info("INFO_MESSAGES.GETTING_DATA_USER_STORAGE_SERVICE")
	return userStorageFetch.VariationName, nil
}

//GetWhiteListedVariationsList ...
func GetWhiteListedVariationsList(userID string, campaign schema.Campaign, options schema.Options) []schema.Variation {
	// check Validity of
	var whiteListedVariationsList []schema.Variation
	for _, variation := range campaign.Variations {
		if len(variation.Segments) == 0 {
			log.Warning("DEBUG_MESSAGES.SEGMENTATION_SKIPPED")
		}
		status := EvaluateSegmentation(variation.Segments, options)
		if status {
			whiteListedVariationsList = append(whiteListedVariationsList, variation)
		}
		log.Info("DEBUG_MESSAGES.SEGMENTATION_STATUS" + strconv.FormatBool(status))
	}
	return whiteListedVariationsList
}

// EvaluateSegmentation function
func EvaluateSegmentation(segments map[string]interface{}, options schema.Options) bool {
	//TO BE COMPLETED
	v := rand.Intn(1)
	return v == 1 
}

// // SetUserData ...
// func SetUserData(config schema.Config, campaign schema.Campaign, variationName, userID string) bool {
// 	UserServiceData := config.UserDatas
// 	if len(UserServiceData) == 0 {
// 		return false
// 	}
// 	for _, userData := range UserServiceData {
// 		if userData.UserID == userID && userData.CampaignKey == campaign.Key {
// 			userData.VariationName = variationName
// 			return true
// 		}
// 	}
// 	return false
// }