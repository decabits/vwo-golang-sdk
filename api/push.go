package api

import (
	log "github.com/golang/glog"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/event"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Push ...
func Push(vwoInstance schema.VwoInstance, tagKey, tagValue, userID string) bool {
	if len(tagKey) > constants.PushAPITagKeyLength {
		log.Info("Tag Key length exceeded")
		return false
	}
	if len(tagValue) > constants.PushAPITagValueLength {
		log.Info("Tag Value length exceeded")
		return false
	}

	impression := utils.CreateImpressionForPush(vwoInstance.SettingsFile, tagKey, tagValue, userID)
	if event.Dispatch(impression) {
		return true
	}

	return false
}
