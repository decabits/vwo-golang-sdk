package core

//TO BE COMPLETED
import (
	"errors"
	"log"

	"github.com/decabits/vwo-golang-sdk/lib/constants"
	"github.com/decabits/vwo-golang-sdk/lib/schema"
	"github.com/decabits/vwo-golang-sdk/lib/utils"
)

// VariationDecider struct
type VariationDecider struct {
	Bucketer         string
	SegmentEvaluator string
}

//UserStorageData ...
type UserStorageData interface {
	GetUserData(config schema.Config, campaignKey, userID string) []schema.UserData
	SetUserData(config schema.Config, campaign schema.Campaign, variationName, userID string) bool
}

//GetCampaignVariation ... TEMPORARY FUNCTION
func GetCampaignVariation(config schema.Config, settingsFile schema.SettingsFile, campaign schema.Campaign, userID string) (schema.Variation, error) {
	if campaign.IsForcedVariation {
		//WhiteListingChecks
	}
	storedVariation, err := GetStoredVariation(config, settingsFile, campaign.Key, userID)
	if err == nil {
		return storedVariation, nil
	}
	//SegmentChecks
	Response, err := GetVariation(config, userID, campaign)
	if err == nil {
		if SetUserData(config, campaign, Response.Name, userID) {
			log.Println("Data Stored")
		}
		return Response, nil
	}
	return schema.Variation{}, nil
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
func GetVariation(config schema.Config, userID string, campaign schema.Campaign) (schema.Variation, error) {
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
	//customVariables = options.CustomVariables
	//variationTargetingVariables = options.VariationTargetingVariables

	//targettedVariation := FindTargetedVariation(userID,campaign,variationTargetingVariables)
	//if (targettedVariation != nil){
	// 	log.Println("Got Variation For User")
	// 	return targettedVariation
	// }

	variation, err := GetVariationFromUserStorage(config, userID, campaign)
	if err != nil {
		return schema.Variation{}, err
	}

	return variation, nil
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

//FindTargetedVariation ...
func FindTargetedVariation(userID string, campaign schema.Campaign, variationTargetingVariables []int) (schema.Variation, error) {
	if campaign.IsForcedVariation == false {
		return schema.Variation{}, errors.New("DEBUG_MESSAGES.WHITELISTING_SKIPPED")
	}
	whiteListedVariationsList := GetWhiteListedVariationsList(userID, campaign, variationTargetingVariables)
	whiteListedVariationsLength := len(whiteListedVariationsList)
	var targettedVariation schema.Variation
	if whiteListedVariationsLength == 0 {
		return schema.Variation{}, errors.New("No White listed variation")
	} else if whiteListedVariationsLength == 1 {
		targettedVariation = whiteListedVariationsList[0]
	}
	whiteListedVariationsList = utils.ScaleVariations(whiteListedVariationsList)
	variationAllocation := utils.GetVariationAllocationRanges(whiteListedVariationsList)
	for i, variation := range whiteListedVariationsList {
		variation.StartVariationAllocation = variationAllocation[i].StartVariationAllocation
		variation.EndVariationAllocation = variationAllocation[i].EndVariationAllocation
	}
	//utils.SetVariationAllocationFromRanges(whiteListedVariationsList, variationAllocation)
	bucketValue := GetBucketValueForUser(userID, constants.MaxTrafficValue)
	targettedVariation, err := GetBucketerVariation(whiteListedVariationsList, bucketValue)
	if err != nil {
		return schema.Variation{}, errors.New("No targetted variation")
	}
	return targettedVariation, nil
}

//GetVariationFromUserStorage ...
func GetVariationFromUserStorage(config schema.Config, userID string, campaign schema.Campaign) (schema.Variation, error) {
	userStorageData := GetUserData(config, userID, campaign.Key)
	if userStorageData.CampaignKey == campaign.Key {
		variationName := userStorageData.VariationName
		variation, err := utils.GetCampaignVariation(campaign, variationName)
		if err == nil {
			log.Println("Got Stored Variation")
			return variation, nil
		}
	}

	return schema.Variation{}, errors.New("No Stored Variation")
}

//GetWhiteListedVariationsList ...
func GetWhiteListedVariationsList(userID string, campaign schema.Campaign, variationTargetingVariables []int) []schema.Variation {
	// check Validity of
	var whiteListedVariationsList []schema.Variation
	// for _, variation := range campaign.Variations {

	// }
	return whiteListedVariationsList
}

//EvaluatePreSegmentation ...
func EvaluatePreSegmentation(userID string, campaign schema.Campaign, customVariables []int) bool {
	// segment := campaign.Segments
	//TO BE COMPLETED
	return true

}

//GetUserData ...
func GetUserData(config schema.Config, campaignKey, userID string) schema.UserData {
	for _, userData := range config.UserDatas {
		if userData.UserID == userID && userData.CampaignKey == campaignKey {
			user := schema.UserData{
				UserID:        userID,
				CampaignKey:   campaignKey,
				VariationName: userData.VariationName,
			}
			return user
		}
	}
	return schema.UserData{}
}

//SetUserData ...
func SetUserData(config schema.Config, campaign schema.Campaign, variationName, userID string) bool {
	UserServiceData := config.UserDatas
	if len(UserServiceData) == 0 {
		return false
	}
	for _, userData := range UserServiceData {
		if userData.UserID == userID && userData.CampaignKey == campaign.Key {
			userData.VariationName = variationName
			return true
		}
	}
	return false
}

//GetStoredVariation ...
func GetStoredVariation(config schema.Config, settingsFile schema.SettingsFile, campaignKey, userID string) (schema.Variation, error) {
	return schema.Variation{}, nil
}
