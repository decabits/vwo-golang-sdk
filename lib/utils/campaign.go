package utils

import (
	"errors"
	"log"
	"math"

	"github.com/decabits/vwo-golang-sdk/lib/constants"
	"github.com/decabits/vwo-golang-sdk/lib/schema"
)

// SetVariationAllocation Sets variation allocation range in the provided campaign.
func SetVariationAllocation(campaigns schema.Campaign) {
	/*
		Args:
			campaign (schema.Campaign struct): Campaign object
	*/

	campaigns.Variations = GetVariationAllocationRanges(campaigns.Variations)
	for _, variation := range campaigns.Variations {
		//logger
		log.Println(variation)
	}
}

// GetVariationAllocationRanges Returns a list of variation allocation ranges.
func GetVariationAllocationRanges(variations []schema.Variation) []schema.Variation {
	/*
		Args:
			variations (list of variations i.e schema.Variation struct)
		Returns:
			list(list of variations i.e schema.Variation struct):
	*/

	var (
		currentAllocation         = 0
		variationAllocationRanges []schema.Variation
	)
	for _, variation := range variations {
		stepFactor := getVariationBucketingRange(variation.Weight)
		if stepFactor != 0 {
			variation.StartVariationAllocation = currentAllocation + 1
			variation.EndVariationAllocation = currentAllocation + stepFactor
			currentAllocation += stepFactor
		} else {
			variation.StartVariationAllocation = -1
			variation.EndVariationAllocation = -1
		}
		variationAllocationRanges = append(variationAllocationRanges, variation)
	}
	return variationAllocationRanges
}

//SetVariationAllocationFromRanges ...
func SetVariationAllocationFromRanges(variations []schema.Variation, variationAllocationRanges []schema.VariationAllocationRange) {
	for i, variation := range variations {
		variation.StartVariationAllocation = variationAllocationRanges[i].StartRange
		variation.EndVariationAllocation = variationAllocationRanges[i].EndRange
	}
}

//Returns the bucket size of variation.
func getVariationBucketingRange(weight int) int {
	/*
		Args:
			weight (int): weight of variation

		Returns:
			int: Bucket start range of Variation
	*/

	if weight == 0 {
		return 0
	}
	startRange := int(math.Ceil(float64(weight) * 100))
	return Min(startRange, constants.MaxTrafficValue)
}

// GetCampaign function finds and returns campaign from given campaign_key.
func GetCampaign(settingsFile schema.SettingsFile, campaignKey string) (schema.Campaign, error) {
	/*
		Args:
			settingsFile (dict): Settings file for the project
			campaignKey (string): Campaign identifier key
		Returns:
			schema.Campaign: Campaign object
	*/
	for _, campaign := range settingsFile.Campaigns {
		if campaign.Key == campaignKey {
			return campaign, nil
		}
	}
	return schema.Campaign{}, errors.New("Campaign not found")
}

// ScaleVariations function It extracts the weights from all the variations inside the campaign and scales them so that the total sum of eligible variations' weights become 100%
func ScaleVariations(variations []schema.Variation) []schema.Variation {
	/*
		Args:
			variations(list): list of variations(dict object) having weight as a property
	*/
	weightSum := 0
	for _, variation := range variations {
		weightSum += variation.Weight
	}
	if weightSum == 0 {
		normalizedWeight := 100 / len(variations)
		for _, variation := range variations {
			variation.Weight = normalizedWeight
		}
	} else {
		for _, variation := range variations {
			variation.Weight = (variation.Weight / weightSum) * 100
		}
	}
	return variations
}

//GetCampaignGoal returns goal from given campaign and Goal_identifier.
func GetCampaignGoal(campaign schema.Campaign, goalIdentifier string) (schema.Goal, error) {
	/*
		 Args:
			campaign (dict): The running campaign
			goalIdentifier (string): Goal identifier
		Returns:
			schema.Goal: Goal corresponding to goal_identifer in respective campaign
	*/
	goals := campaign.Goals
	for _, goal := range goals {
		if goal.Identifier == goalIdentifier {
			return goal, nil
		}
	}
	return schema.Goal{}, errors.New("Goal NOt Found")
}

// GetCampaignVariation returns variation from given campaign and variation_name.
func GetCampaignVariation(campaign schema.Campaign, variationName string) (schema.Variation, error) {
	/*
		 Args:
			campaign (dict): The running campaign
			variationName (string): Variation identifier
		Returns:
			schema.Variation: Variation corresponding to variation_name in respective campaign
	*/
	if len(campaign.Variations) == 0 {
		return schema.Variation{}, errors.New("Invalid Campaign")
	}
	for _, variation := range campaign.Variations {
		if variation.Name == variationName {
			return variation, nil
		}
	}
	return schema.Variation{}, errors.New("CampaignVariation not found")
}

//GetControlVariation Returns control variation from a given campaign
func GetControlVariation(campaign schema.Campaign) schema.Variation {
	/*
		Args:
			campaign (schema.Campaign): Running campaign
		Returns:
			variation (dischema.Variation): Control variation from the campaign, ie having id = 1
	*/

	variations := campaign.Variations
	for _, variation := range variations {
		if variation.ID == "1" {
			return variation
		}
	}
	return schema.Variation{}
}

// //GetSegments Returns segments from the campaign
// func GetSegments(campaign schema.Campaign) schema.Segment {
// 	/*
// 		Args:
// 			campaign(schema.Campaign): Running campaign
// 		Returns:
// 			segments(schema.Segment): a dsl of segments
// 	*/

// 	segments := campaign.Segments
// 	return segments
// }

// Min function
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max function
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
