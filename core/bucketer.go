package core

import (
	"fmt"
	"math"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/spaolacci/murmur3"
)

const (
	umax32Bit = 0xFFFFFFFF
	bucketer  = "bucketer.go"
)

// GetBucketerVariation function returns the Variation by checking the Start and End Bucket Allocations of each Variation
func GetBucketerVariation(variations []schema.Variation, bucketValue int) (schema.Variation, error) {
	/*
		Args:
			variations : list of variations (schema.Variation)
			bucketValue: the bucket value of the user

		Returns:
			schema.Variation: variation  allotted to the user
			error: if no variation found, else nil
	*/

	for _, variation := range variations {
		if variation.StartVariationAllocation <= bucketValue && variation.EndVariationAllocation >= bucketValue {
			return variation, nil
		}
	}
	return schema.Variation{}, fmt.Errorf(constants.ErrorMessageNoVariationForBucketValue, bucketValue)
}

// GetBucketValueForUser returns Bucket Value of the user by hashing the userId with murmur hash and scaling it down.
func GetBucketValueForUser(vwoInstance schema.VwoInstance, userID string, maxValue, multiplier float64) int {
	/*
		Args:
			vwoInstance: vwo Instance for logger implementation
			userID: the unique ID assigned to User
			maxValue: maximum value that can be alloted to the bucket value
			multiplier: value for distributing ranges slightly

		Returns:
			int: the bucket value allotted to User
			(between 1 to MAX_TRAFFIC_PERCENT)
	*/

	hashValue := hash(userID) & umax32Bit
	ratio := float64(hashValue) / math.Pow(2, 32)
	multipliedValue := (maxValue*ratio + 1) * multiplier
	bucketValue := int(math.Floor(multipliedValue))

	message := fmt.Sprintf(constants.InfoMessageUserHashBucketValue, userID, hashValue, bucketValue)
	utils.LogMessage(vwoInstance, constants.Info, bucketer, message)

	return bucketValue
}

// IsUserPartOfCampaign calculates if the provided userID should become part of the campaign or not
func IsUserPartOfCampaign(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign) bool {
	/*
		Args:
			userID: the unique ID assigned to a user
			campaign: for getting traffic allotted to the campaign

		Returns:
			bool: if User is a part of Campaign or not
	*/

	if len(campaign.Variations) == 0 {
		return false
	}
	valueAssignedToUser := GetBucketValueForUser(vwoInstance, userID, constants.MaxTrafficPercent, 1)
	isUserPart := valueAssignedToUser != 0 && valueAssignedToUser <= campaign.PercentTraffic

	message := fmt.Sprintf(constants.InfoMessageUserEligibilityForCampaign, userID, isUserPart)
	utils.LogMessage(vwoInstance, constants.Info, bucketer, message)

	return isUserPart
}

// BucketUserToVariation validates the User ID and returns Variation into which the User is bucketed in.
func BucketUserToVariation(vwoInstance schema.VwoInstance, userID string, campaign schema.Campaign) (schema.Variation, error) {
	/*
		Args:
		    userID: the unique ID assigned to User
		    campaign: the Campaign of which User is a part of

		Returns:
			schema.Variation: variation data into which user is bucketed in
			error: if no variation found, else nil
	*/

	if len(campaign.Variations) == 0 {
		return schema.Variation{}, fmt.Errorf(constants.ErrorMessageNoVariationInCampaign, campaign.Key)
	}
	multiplier := (float64(constants.MaxTrafficValue) / float64(campaign.PercentTraffic)) / 100
	bucketValue := GetBucketValueForUser(vwoInstance, userID, constants.MaxTrafficValue, multiplier)
	return GetBucketerVariation(campaign.Variations, bucketValue)
}

// hash function generates hash value for given string using murmur hash
func hash(s string) uint32 {
	hasher := murmur3.New32WithSeed(uint32(constants.SeedValue))
	hasher.Write([]byte(s))
	return hasher.Sum32()
}
