package core

import (
	"errors"
	"math"

	log "github.com/golang/glog"

	"github.com/decabits/vwo-go-sdk/constants"
	"github.com/decabits/vwo-go-sdk/schema"
	"github.com/spaolacci/murmur3"
)

const (
	umax32Bit = 0xFFFFFFFF
)

//GetBucketerVariation function returns the Variation by checking the Start and End Bucket Allocations of each Variation
func GetBucketerVariation(variations []schema.Variation, bucketValue int) (schema.Variation, error) {
	/*
		Args:
			variations (list): list of variations
			bucketValue (int): the bucket value of the user
		Returns:
			(dict|None): variation data allotted to the user or None if not
	*/

	for _, variation := range variations {
		if variation.StartVariationAllocation <= bucketValue && variation.EndVariationAllocation >= bucketValue {
			return variation, nil
		}
	}
	return schema.Variation{}, errors.New("Variation Not Found")
}

//GetBucketValueForUser function returns Bucket Value of the user by hashing the userId with murmur hash and scaling it down.
func GetBucketValueForUser(userID string, maxValue int) int {
	/*
		Args:
			user_id (string): the unique ID assigned to User
			max_value(int): maximum value that can be alloted to the bucket value
			multiplier(int): value for distributing ranges slightly
		Returns:
			int: the bucket value allotted to User
			(between 1 to MAX_TRAFFIC_PERCENT)
	*/

	hashValue := int(hash(userID) & umax32Bit)
	ratio := float64(hashValue) / math.Pow(2, 32)
	multipliedValue := float64(maxValue)*ratio + 1
	bucketValue := int(multipliedValue)
	log.Info("User Hash Bucket Value")
	return bucketValue
}

// GetBucketValueForUserMultiplier returns Bucket Value of the user by hashing the userId with murmur hash and scaling it down.
func GetBucketValueForUserMultiplier(userID string, maxValue, multiplier float64) int {
	/*
		Args:
			user_id (string): the unique ID assigned to User
			max_value(int): maximum value that can be alloted to the bucket value
			multiplier(int): value for distributing ranges slightly
		Returns:
			int: the bucket value allotted to User
			(between 1 to MAX_TRAFFIC_PERCENT)
	*/

	hashValue := float64(hash(userID) & umax32Bit)
	ratio := hashValue / math.Pow(2, 32)
	multipliedValue := (float64(maxValue)*ratio + 1) * multiplier
	bucketValue := int(multipliedValue)
	log.Info("User Hash Bucket Value")
	return bucketValue
}

// IsUserPartOfCampaign calculates if the provided user_id should become part of the campaign or not
func IsUserPartOfCampaign(userID string, campaign schema.Campaign) bool {
	/*
		Args:
			user_id (strings): the unique ID assigned to a user
			campaign (dict): for getting traffic allotted to the campaign

		Returns:
			bool: if User is a part of Campaign or not
	*/

	if len(campaign.Variations) == 0 {
		return false
	}
	valueAssignedToUser := GetBucketValueForUser(userID, constants.MaxTrafficValue)
	if valueAssignedToUser > 0 && valueAssignedToUser <= campaign.PercentTraffic {
		log.Info("User Eligibile For Campaign")
		return true
	}
	return false
}

// BucketUserToVariation validates the User ID and returns Variation into which the User is bucketed in.
func BucketUserToVariation(userID string, campaign schema.Campaign) (schema.Variation, error) {
	/*
	   Args:
	       user_id (string): the unique ID assigned to User
	       campaign (dict): the Campaign of which User is a part of

	   Returns:
	       (dict|None): variation data into which user is bucketed in or None if not
	*/

	if len(campaign.Variations) == 0 {
		return schema.Variation{}, errors.New("Invalid Campaign")
	}
	normalize := float64(constants.MaxTrafficValue / campaign.PercentTraffic)
	multiplier := normalize / 100
	bucketValue := GetBucketValueForUserMultiplier(userID, constants.MaxTrafficValue, multiplier)
	return GetBucketerVariation(campaign.Variations, bucketValue)
}

func hash(s string) uint32 {
	hasher := murmur3.New32WithSeed(uint32(constants.SeedValue))
	hasher.Write([]byte(s))
	return hasher.Sum32()
}
