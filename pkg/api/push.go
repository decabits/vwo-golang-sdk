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

package api

import (
	"fmt"

	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/event"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
	"github.com/decabits/vwo-golang-sdk/pkg/utils"
)

const push = "push.go"

// Push function
/*
This API method: Pushes the key-value tag pair for a particular user
1. Validates the arguments being passed
2. Checks the length of tag Key and Value
3. Sends a call to VWO push api
*/
func (vwo *VWOInstance) Push(tagKey, tagValue, userID string) bool {
	/*
		Args:
			tagKey: Key of the corresponding tag
			tagValue: Value of the corresponding tag
			userID: Unique identification of user

		Returns:
			bool: true if the push api call is done, else false
	*/
	if !utils.ValidatePush(tagKey, tagValue, userID) {
		message := fmt.Sprintf(constants.ErrorMessagesPushAPIMissingParams)
		utils.LogMessage(vwo.Logger, constants.Error, push, message)
		return false
	}

	if len(tagKey) > constants.PushAPITagKeyLength {
		message := fmt.Sprintf(constants.ErrorMessagesTagKeyLengthExceeded, tagKey, userID)
		utils.LogMessage(vwo.Logger, constants.Error, push, message)
		return false
	}
	if len(tagValue) > constants.PushAPITagValueLength {
		message := fmt.Sprintf(constants.ErrorMessagesTagValueLengthExceeded, tagValue, tagKey, userID)
		utils.LogMessage(vwo.Logger, constants.Error, push, message)
		return false
	}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      vwo.SettingsFile,
		UserStorage:       vwo.UserStorage,
		Logger:            vwo.Logger,
		IsDevelopmentMode: vwo.IsDevelopmentMode,
		UserID:            userID,
	}
	impression := utils.CreateImpressionForPush(vwoInstance, tagKey, tagValue, userID)
	event.Dispatch(vwoInstance, impression)

	return true
}