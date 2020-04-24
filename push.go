package vwo

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Push ...
func (vwo *VWOInstance) Push(tagKey, tagValue, userID string) bool {
	if len(tagKey) > constants.PushAPITagKeyLength {
		vwo.Logger.Error("ERROR_MESSAGES.TAG_KEY_LENGTH_EXCEEDED")
		return false
	}
	if len(tagValue) > constants.PushAPITagValueLength {
		vwo.Logger.Error("ERROR_MESSAGES.TAG_VALUE_LENGTH_EXCEEDED")
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
