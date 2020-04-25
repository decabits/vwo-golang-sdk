package vwo

import (
	"fmt"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

const push = "push.go"

// Push ...
func (vwo *VWOInstance) Push(tagKey, tagValue, userID string) bool {
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
	}
	impression := utils.CreateImpressionForPush(vwoInstance, tagKey, tagValue, userID)
	event.Dispatch(vwoInstance, impression)

	return true
}
