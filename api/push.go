package api

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Push ...
func Push(vwoInstance schema.VwoInstance, tagKey, tagValue, userID string) bool {
	if len(tagKey) > constants.PushAPITagKeyLength {
		vwoInstance.Logger.Error("ERROR_MESSAGES.TAG_KEY_LENGTH_EXCEEDED")
		return false
	}
	if len(tagValue) > constants.PushAPITagValueLength {
		vwoInstance.Logger.Error("ERROR_MESSAGES.TAG_VALUE_LENGTH_EXCEEDED")
		return false
	}

	impression := utils.CreateImpressionForPush(vwoInstance, tagKey, tagValue, userID)
	event.Dispatch(vwoInstance, impression)

	return true
}
