package utils

import (
	"github.com/decabits/vwo-go-sdk/constants"
	"github.com/decabits/vwo-go-sdk/schema"
)

// GetURLParams ...
func GetURLParams(settingsFile schema.SettingsFile, tagKey, tagValue, userID string) schema.Impression {
	//Implementation of Tags
	params := GetCommonProperties(userID, settingsFile)
	params.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsPush
	return params
}
