package core

//TO BE COMPLETED
import (
	"errors"
	"log"

	"github.com/decabits/vwo-golang-sdk/lib/schema"
	"github.com/decabits/vwo-golang-sdk/lib/utils"
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
func GetVariation(userID string, campaign schema.Campaign, options schema.Options) (schema.Variation, error) {
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

	customVariables = options.CustomVariables
	variationTargetingVariables = options.VariationTargetingVariables
	
	targettedVariation := FindTargetedVariation(userID,campaign,variationTargetingVariables)
	if (targettedVariation != nil){
		log.Println("Got Variation For User")
		return targettedVariation
	}

	variation, err := GetVariationFromUserStorage(userID, campaign)
	if err != nil {
		return nil, err
	}

	variation, err := GetVariationFromUserStorage(userID, campaign)
	if err != nil {
		return nil, err
	}

	log.Println("Got No Variation For User")
	return nil, errors.New("Got No Variation For User")
}

//FindTargetedVariation
func FindTargetedVariation(userID string, campaign schema.Campaign, variationTargetingVariables schema.Options.VariationTargetingVariables) (schema.Variation, error){
	if campaign.IsForcedVariation == false {
		log.Println("DEBUG_MESSAGES.WHITELISTING_SKIPPED")
		return nil
	} 
	else {
		whiteListedVariationsList := GetWhiteListedVariationsList(userID, campaign, variationTargetingVariables)
		whiteListedVariationsLength := len(whiteListedVariationsList)
		if whiteListedVariationsLength == 0{
			return nil,errors.New("No White listed variation")
		}else if whiteListedVariationsLength == 1{
			targetedVariation = whiteListedVariationsList[0]
		}else{
			whiteListedVariationsList = utils.ScaleVariations(whiteListedVariationsList)
			variationAllocation := utils.GetVariationAllocationRanges(whiteListedVariationsList)
			utils.SetVariationAllocationFromRanges(whiteListedVariationsList,variationAllocation)
			bucketValue := GetBucketValueForUser(userID , constants.MaxTrafficValue)
			targettedVariation, err := GetBucketerVariation(variations []schema.Variation, bucketValue)
			if (err != nil){
				return nil,errors.New("No targetted variation")
		}
	}
	return targettedVariation,nil
}

//GetVariationFromUserStorage ...
func GetVariationFromUserStorage(userId string, campaign schema.Campaign) (schema.Variation,error){
	userStorageData := GetUserStorageData(userId,campaign.Key)
	if (userStorageData.CampaignKey==campaign.Key){
		variationName := userStorageData.VariationName
		variation := utils.GetCampaignVariation(campaign,variationName)
		if(variation != nil){
			log.Println("Got Stored Variation")
			return variation,nil
		}
	}

	return nil,errors.New("No Stored Variation")
}

//GetWhiteListedVariationsList ...
func GetWhiteListedVariationsList(userID string, campaign schema.Campaign, variationTargetingVariables schema.Options) []schema.Variation {
	// check Validity of 
	var whiteListedVariationsList = []schema.Variation
	for _, variation := range campaign.Variations {
		//TO BE COMPLETED
	}
	return whiteListedVariationsList
}

//EvaluatePreSegmentation ...
func EvaluatePreSegmentation(userId string, campaign schema.Campaign, customVariables schema.Options.CustomVariables) bool{
	segment := campaign.Segments
	//TO BE COMPLETED
}

//GetUserStorageData ...
func GetUserStorageData(userID string, campaignKey string) {
	//TO BE COMPLETED
}

//SetUserStorageData ...
func SetUserStorageData(userID, campaignKey, variationName string) bool{
	/*
	If UserStorage is provided and variation was found,
    set the assigned variation in UserStorage.
	It creates bucket and then stores.
    Args:
        userID (string): Unique user identifier
        campaign_key (string): Unique campaign identifier
        variation_name (string): variation identifier
    Returns:
		bool: true if found otherwise false
	*/
	
	//TO BE COMPLETED

	// newUserStorageData = {
    //         "userId": userID,
    //         "campaignKey": campaignKey,
    //         "variationName": variationName,
    //     }
	
}
