/*
   Copyright 2019-2020 Wingify Software Pvt. Ltd.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package utils

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
)

const impressions = "impression.go"

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

	message := fmt.Sprintf(constants.DebugMessageImpressionForPush, vwoInstance.API, impression)
	LogMessage(vwoInstance.Logger, constants.Debug, impressions, message)

	return impression
}

// CreateImpressionTrackingGoal creates the impression from the arguments passed to track goal
func CreateImpressionTrackingGoal(vwoInstance schema.VwoInstance, variationID int, userID string, campaignID, goalID, revenueValue int) schema.Impression {
	/*
		Args:
		    variationID : Variation identifier
			userID : User identifier
			campaignID : Campaign identifier
		    goalID : Goal identifier
		    revenueValue : Revenue goal for the campaign

		Returns:
			schema.Impression: Imression struct with required values
	*/
	impression := getCommonProperties(vwoInstance, userID)

	impression.ExperimentID = campaignID
	impression.Combination = variationID

	impression.URL = constants.HTTPSProtocol + constants.EndPointsBaseURL + constants.EndPointsTrackGoal
	impression.GoalID = goalID
	if revenueValue > 0 {
		impression.R = revenueValue
	}

	message := fmt.Sprintf(constants.DebugMessageImpressionForTrackGoal, vwoInstance.API, impression)
	LogMessage(vwoInstance.Logger, constants.Debug, impressions, message)

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

	message := fmt.Sprintf(constants.DebugMessageImpressionForTrackUser, vwoInstance.API, impression)
	LogMessage(vwoInstance.Logger, constants.Debug, impressions, message)

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
		UID:       url.QueryEscape(userID),
	}
}
