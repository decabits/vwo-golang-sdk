package utils

import (
	"github.com/decabits/vwo-golang-sdk/lib/constants"
	"github.com/decabits/vwo-golang-sdk/lib/schema"
)

//GetURLParams ...
func GetURLParams(settingsFile schema.SettingsFile, tagKey, tagValue, userID string) schema.Impression {
	url := constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsPush
	//Implementation of Tags
	params := GetCommonProperties(userID, settingsFile)
	params.URL = url
	return params
}
