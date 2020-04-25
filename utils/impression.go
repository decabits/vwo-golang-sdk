package utils

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
)

// CreateImpressionForPush creates the impression from the arguments passed to push
func CreateImpressionForPush(vwoInstance schema.VwoInstance, tagKey, tagValue, userID string) schema.Impression {
	/*
		Args:
			tagKey : Campaign identifier
			tagValue : Variation identifier
			userId : User identifier

		Returns:
			schema.Impression: Imression struct with required values
	*/
	impression := getCommonProperties(vwoInstance, userID)
	impression.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsPush
	parameters := url.Values{}
	parameters.Add(tagKey, tagValue)
	impression.U = parameters.Encode()

	file := "impression.go"
	message := fmt.Sprintf(constants.InfoMessageImpressionForPush, impression)
	LogMessage(vwoInstance, constants.Info, file, message)

	return impression
}

// CreateImpressionTrackingGoal creates the impression from the arguments passed to track goal
func CreateImpressionTrackingGoal(vwoInstance schema.VwoInstance, variationID int, userID string, campaignID, goalID, revenueGoal int) schema.Impression {
	/*
		Args:
		    variationID : Variation identifier
			userID : User identifier
			campaignID : Campaign identifier
		    goalID : Goal identifier
		    revenueGoal : Revenue goal for the campaign

		Returns:
			schema.Impression: Imression struct with required values
	*/
	impression := getCommonProperties(vwoInstance, userID)

	impression.ExperimentID = campaignID
	impression.Combination = variationID

	impression.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsTrackGoal
	impression.GoalID = goalID
	if revenueGoal > 0 {
		impression.R = revenueGoal
	}

	file := "impression.go"
	message := fmt.Sprintf(constants.InfoMessageImpressionForTrackGoal, impression)
	LogMessage(vwoInstance, constants.Info, file, message)

	return impression
}

// CreateImpressionTrackingUser creates the impression from the arguments passed to track user
func CreateImpressionTrackingUser(vwoInstance schema.VwoInstance, campaignID int, variationID int, userID string) schema.Impression {
	/*
		Args:
			variationID : Variation identifier
			userID : User identifier
			campaignID : Campaign identifier

		Returns:
			schema.Impression: Imression struct with required values
	*/
	impression := getCommonProperties(vwoInstance, userID)

	impression.ExperimentID = campaignID
	impression.Combination = variationID

	impression.ED = `{\"p\":\"` + constants.Platform + `\"}`
	impression.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsTrackUser

	file := "impression.go"
	message := fmt.Sprintf(constants.InfoMessageImpressionForTrackUser, impression)
	LogMessage(vwoInstance, constants.Info, file, message)

	return impression
}

// getCommonProperties returns commonly used params for making requests to our servers.
func getCommonProperties(vwoInstance schema.VwoInstance, userID string) schema.Impression {
	/*
		Args:
			userID : Unique identification of user

		Returns:
			schema.Impression: commonly used params for making call to our servers
	*/
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
