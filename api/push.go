package api

import (
	log "github.com/golang/glog"

	"github.com/Piyushhbhutoria/vwo-go-sdk/constants"
	"github.com/Piyushhbhutoria/vwo-go-sdk/event"
	"github.com/Piyushhbhutoria/vwo-go-sdk/schema"
	"github.com/Piyushhbhutoria/vwo-go-sdk/utils"
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
