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
		vwoInstance.Logger.Info("Tag Key length exceeded")
		return false
	}
	if len(tagValue) > constants.PushAPITagValueLength {
		vwoInstance.Logger.Info("Tag Value length exceeded")
		return false
	}

	impression := utils.CreateImpressionForPush(vwoInstance, tagKey, tagValue, userID)
	if event.Dispatch(vwoInstance, impression) {
		return true
	}

	return false
}
