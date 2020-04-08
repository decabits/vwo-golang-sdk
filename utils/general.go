package utils

import (
	"github.com/decabits/vwo-golang-sdk/schema"
)

// CheckCampaignType matches campaign type
func CheckCampaignType(campaign schema.Campaign, campaignType string) bool {
	/*
		Args:
			campaign : Campaign object
			campaignType : Type of campaign to be matched
		Return:
			bool: true if the type matches else false
	*/
	return campaign.Type == campaignType
}

// GetKeyValue returns first key value pair of the given map
func GetKeyValue(obj map[string]interface{}) (string, interface{}) {
	/*
		Args:
			obj: map whose firsr key value pair is needed
		
		Return:
			string: Key
			interface: value
	*/
	for k, v := range obj {
		return k, v
	}
	return "", nil
}
