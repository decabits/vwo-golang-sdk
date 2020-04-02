package utils

import (
	"github.com/decabits/vwo-golang-sdk/schema"
)

// GetKeyValue function
func GetKeyValue(segments map[string]interface{}) (string, map[string]interface{}) {
	// To be done
	return "", make(map[string]interface{})
}

// CheckCampaignType matched campaign type
func CheckCampaignType(campaign schema.Campaign, campaignType string) bool {
	return campaign.Type == campaignType
}
