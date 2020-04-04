package core

import (
	"errors"
	"math"
	"strconv"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
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

// GetBucketValueForUser function returns Bucket Value of the user by hashing the userId with murmur hash and scaling it down.
func GetBucketValueForUser(vwoInstance schema.VwoInstance, userID string, maxValue int) int {
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
	vwoInstance.Logger.Info("DEBUG_MESSAGES.USER_HASH_BUCKET_VALUE " + strconv.Itoa(bucketValue))
	return bucketValue
}

func hash(s string) uint32 {
	hasher := murmur3.New32WithSeed(uint32(constants.SeedValue))
	hasher.Write([]byte(s))
	return hasher.Sum32()
}
