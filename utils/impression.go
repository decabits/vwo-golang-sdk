package utils

import (
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
)

// CreateImpressionForPush function
func CreateImpressionForPush(vwoInstance schema.VwoInstance, tagKey, tagValue, userID string) schema.Impression {
	impression := getCommonProperties(vwoInstance, userID)
	impression.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsPush
	parameters := url.Values{}
	parameters.Add(tagKey, tagValue)
	impression.U = parameters.Encode()
	vwoInstance.Logger.Info("DEBUG_MESSAGES.IMPRESSION_FOR_PUSH ", impression)
	return impression
}

// CreateImpressionTrackingGoal ...
func CreateImpressionTrackingGoal(vwoInstance schema.VwoInstance, variationID int, userID string, campaignID, goalID, revenueGoal int) schema.Impression {
	impression := getCommonProperties(vwoInstance, userID)

	impression.ExperimentID = campaignID
	impression.Combination = variationID

	impression.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsTrackGoal
	impression.GoalID = goalID
	if revenueGoal > 0 {
		impression.R = revenueGoal
	}

	vwoInstance.Logger.Info("DEBUG_MESSAGES.IMPRESSION_FOR_GOAL_TRACK ", impression)
	return impression
}

// CreateImpressionTrackingUser function to track user
func CreateImpressionTrackingUser(vwoInstance schema.VwoInstance, campaignID int, variationID int, userID string) schema.Impression {
	impression := getCommonProperties(vwoInstance, userID)

	impression.ExperimentID = campaignID
	impression.Combination = variationID

	impression.ED = `{\"p\":\"` + constants.Platform + `\"}`
	impression.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsTrackUser

	vwoInstance.Logger.Info("DEBUG_MESSAGES.IMPRESSION_FOR_TRACK_USER ", impression)
	return impression
}

func getCommonProperties(vwoInstance schema.VwoInstance, userID string) schema.Impression {
	return schema.Impression{
		Random:    rand.Float32(),
		Sdk:       constants.SDKName,
		SdkV:      constants.SDKVersion,
		Ap:        constants.Platform,
		SID:       strconv.FormatInt(time.Now().Unix(), 10),
		U:         generateFor(vwoInstance, userID, vwoInstance.SettingsFile.AccountID),
		AccountID: vwoInstance.SettingsFile.AccountID,
		UID:       userID,
	}
}
