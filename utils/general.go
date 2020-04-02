package utils

import (
	"github.com/decabits/vwo-golang-sdk/schema"
)

// GetKeyValue function
func GetKeyValue(segments []schema.Segment) (string, []schema.Segment) {
	// To be done
	return "", []schema.Segment{}
}

// CheckCampaignType matched campaign type
func CheckCampaignType(campaign schema.Campaign, campaignType string) bool {
	return campaign.Type == campaignType
}
