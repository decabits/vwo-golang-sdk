package utils

import (
	"math/rand"
	"net/url"
	"time"

	log "github.com/golang/glog"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
)

// CreateImpressionForPush function
func CreateImpressionForPush(settingsFile schema.SettingsFile, tagKey, tagValue, userID string) schema.Impression {
	impression := GetCommonProperties(userID, settingsFile)
	impression.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsPush
	parameters := url.Values{}
	parameters.Add(tagKey, tagValue)
	impression.U = parameters.Encode()
	return impression
}

// CreateImpressionExtended ...
func CreateImpressionExtended(settingsFile schema.SettingsFile, variationID, userID string, campaignID, goalID, revenueGoal int) schema.Impression {
	impression := GetCommonProperties(userID, settingsFile)

	impression.ExperimentID = campaignID
	impression.Combination = variationID

	impression.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsTrackGoal
	impression.GoalID = goalID
	if revenueGoal > 0 {
		impression.R = revenueGoal
		log.Info("DEBUG_MESSAGES.IMPRESSION_FOR_TRACK_USER")
	}

	return impression
}

//CreateImpression ...
func CreateImpression(settingsFile schema.SettingsFile, campaignID int, variationID, userID string) schema.Impression {
	impression := GetCommonProperties(userID, settingsFile)

	impression.ExperimentID = campaignID
	impression.Combination = variationID

	impression.ED = []byte(`{'p': constants.Platform}`)
	impression.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsTrackUser

	log.Info("DEBUG_MESSAGES.IMPRESSION_FOR_TRACK_USER")
	return impression
}

// GetCommonProperties function
func GetCommonProperties(userID string, settingsFile schema.SettingsFile) schema.Impression {
	return schema.Impression{
		Random:       rand.Float32(),
		Sdk:          constants.SDKName,
		SdkV:         constants.SDKVersion,
		Ap:           constants.Platform,
		SID:          string(time.Now().Unix()),
		U:            GenerateFor(userID, settingsFile.AccountID),
		AccountID:    settingsFile.AccountID,
		UID:          userID,
		URL:          "",
		GoalID:       0,
		ExperimentID: 0,
		Combination:  "",
		R:            0,
	}
}
