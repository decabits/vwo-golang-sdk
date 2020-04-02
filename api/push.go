package api

import (
	log "github.com/golang/glog"

	"github.com/decabits/vwo-go-sdk/constants"
	"github.com/decabits/vwo-go-sdk/event"
	"github.com/decabits/vwo-go-sdk/schema"
	"github.com/decabits/vwo-go-sdk/utils"
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

	impression := utils.GetURLParams(vwoInstance.SettingsFile, tagKey, tagValue, userID)
	if event.Dispatch(impression) {
		return true
	}

	return false
}
