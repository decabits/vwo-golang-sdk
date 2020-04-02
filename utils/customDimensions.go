package utils

import (
	"github.com/Piyushhbhutoria/vwo-go-sdk/schema"
	"github.com/Piyushhbhutoria/vwo-go-sdk/constants"
)

// GetURLParams ...
func GetURLParams(settingsFile schema.SettingsFile, tagKey, tagValue, userID string) schema.Impression {
	//Implementation of Tags
	params := GetCommonProperties(userID, settingsFile)
	params.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsPush
	return params
}

