package utils

import (
	"github.com/decabits/vwo-golang-sdk/schema"
)

// CheckCampaignType matched campaign type
func CheckCampaignType(campaign schema.Campaign, campaignType string) bool {
	return campaign.Type == campaignType
}

func GetKeyValue(obj map[string]interface{}) (string, interface{}) {
	for k, v := range obj {
		return k, v
	}
	return "", nil
}
