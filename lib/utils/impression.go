package utils

import (
	"log"
	"math/rand"
	"time"

	"github.com/decabits/vwo-golang-sdk/lib/constants"
	"github.com/decabits/vwo-golang-sdk/lib/schema"
)

//CreateImpressionExtended ...
func CreateImpressionExtended(settingsFile schema.SettingsFile, variationID, userID string, campaignID, goalID, revenueGoal int) schema.Impression {
	impression := GetCommonProperties(userID, settingsFile)

	impression.ExperimentID = campaignID
	impression.Combination = variationID

	URL := constants.HTTPSProtocol + constants.EndPointsBaseURL

	impression.URL = URL + constants.EndPointsTrackGoal
	impression.GoalID = goalID
	if revenueGoal > 0 {
		impression.R = revenueGoal
		log.Println("Imression For Track User")
	}

	return impression

}

//CreateImpression ...
func CreateImpression(settingsFile schema.SettingsFile, campaignID int, variationID string, userID string) schema.Impression {
	impression := GetCommonProperties(userID, settingsFile)

	impression.ExperimentID = campaignID
	impression.Combination = variationID

	URL := constants.HTTPSProtocol + constants.EndPointsBaseURL

	// impression.ED=json.dumps({'p': constants.PLATFORM}))
	impression.URL = URL + constants.EndPointsTrackUser
	log.Println("Imression For Track User")

	return impression

}

//GetCommonProperties ...
func GetCommonProperties(userID string, settingsFile schema.SettingsFile) schema.Impression {
	accountID := settingsFile.AccountID

	var properties schema.Impression
	properties.Random = rand.Float32()
	properties.Sdk = constants.SDKName
	properties.SdkV = constants.SDKVersion
	properties.Ap = constants.Platform
	properties.SID = string(time.Now().Unix())
	properties.U = GenerateFor(userID, accountID)
	properties.AccountID = settingsFile.AccountID
	properties.UID = userID
	properties.URL = ""
	properties.GoalID = 0
	properties.ExperimentID = 0
	properties.Combination = ""
	properties.R = 0

	return properties
}
