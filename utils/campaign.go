package utils

import (
	"errors"
	"math"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
)

// GetVariationAllocationRanges returns a list of variation with set allocation ranges.
func GetVariationAllocationRanges(vwoInstance schema.VwoInstance, variations []schema.Variation) []schema.Variation {
	/*
		Args:
			variations: list of variations(schema.Variation)
		Returns:
			variations: list of variations(schema.Variation)
	*/

	var (
		currentAllocation         = 0
		variationAllocationRanges []schema.Variation
	)
	for _, variation := range variations {
		stepFactor := GetVariationBucketingRange(variation.Weight)
		if stepFactor != 0 {
			variation.StartVariationAllocation = currentAllocation + 1
			variation.EndVariationAllocation = currentAllocation + stepFactor
			currentAllocation += stepFactor
		} else {
			variation.StartVariationAllocation = -1
			variation.EndVariationAllocation = -1
		}
		vwoInstance.Logger.Infof("Variation: %+v with weight: %+v got range as: ( %+v - %+v ))", variation.Name, variation.Weight, variation.StartVariationAllocation, variation.EndVariationAllocation)
		variationAllocationRanges = append(variationAllocationRanges, variation)
	}
	return variationAllocationRanges
}

// GetVariationBucketingRange Returns the bucket size of variation.
func GetVariationBucketingRange(weight float64) int {
	/*
		Args:
			weight: weight of variation
		Returns:
			int: Bucket start range of Variation
	*/

	if weight == 0 {
		return 0
	}
	startRange := int(math.Ceil(weight * 100))
	return min(startRange, constants.MaxTrafficValue)
}

// GetCampaign function finds and returns campaign from given campaign_key.
func GetCampaign(settingsFile schema.SettingsFile, campaignKey string) (schema.Campaign, error) {
	/*
		Args:
			settingsFile  : Settings file for the project
			campaignKey: Campaign identifier key
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
func  ScaleVariations(variations []schema.Variation) []schema.Variation {
	/*
		Args:
			variations: list of variations(schema.Variartion) having weight as a property

		Return:
			variations: list of variations(schema.Variartion)
	*/
	weightSum := 0.0
	for _, variation := range variations {
		weightSum += variation.Weight
	}
	if weightSum == 0 {
		normalizedWeight := 100.0 / float64(len(variations))
		for i := range variations {
			variations[i].Weight = normalizedWeight
		}
	} else {
		for _, variation := range variations {
			variation.Weight = (variation.Weight / weightSum) * 100
		}
	}
	return variations
}

// GetCampaignGoal returns goal from given campaign and Goal_identifier.
func GetCampaignGoal(campaign schema.Campaign, goalIdentifier string) (schema.Goal, error) {
	/*
		 Args:
			campaign: The running campaign
			goalIdentifier: Goal identifier
		Returns:
			schema.Goal: Goal corresponding to goal_identifer in respective campaign
	*/
	goals := campaign.Goals
	for _, goal := range goals {
		if goal.Identifier == goalIdentifier {
			return goal, nil
		}
	}
	return schema.Goal{}, errors.New("Goal Not Found")
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
		return schema.Variation{}, errors.New("No Campaign available")
	}
	for _, variation := range campaign.Variations {
		if variation.Name == variationName {
			return variation, nil
		}
	}
	return schema.Variation{}, errors.New("CampaignVariation not found")
}

// GetControlVariation Returns control variation from a given campaign
func GetControlVariation(campaign schema.Campaign) schema.Variation {
	/*
		Args:
			campaign: Running campaign
		Returns:
			schema.Variation: Control variation from the campaign, ie having id = 1
	*/

	for _, variation := range campaign.Variations {
		if variation.ID == 1 {
			return variation
		}
	}
	return schema.Variation{}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
