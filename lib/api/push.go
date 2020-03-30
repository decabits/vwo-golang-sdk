package api

import (
	"log"

	"github.com/decabits/vwo-golang-sdk/lib/constants"
	"github.com/decabits/vwo-golang-sdk/lib/core"
	"github.com/decabits/vwo-golang-sdk/lib/schema"
	"github.com/decabits/vwo-golang-sdk/lib/service"
	"github.com/decabits/vwo-golang-sdk/lib/utils"
)

//Push ...
func Push(vwoInstance schema.VwoInstance, tagKey, tagValue, userID string) bool{

	if	(len(tagKey)>constants.PushAPITagKeyLength){
		log.Println("Tag Key length exceeded")
		return false
	}
	if	(len(tagValue)>constants.PushAPITagValueLength){
		log.Println("Tag Value length exceeded")
		return false
	}

	impression = utils.GetURLParams(vwoInstance.SettingsFile, tagKey, tagValue, userID)
	if event.Dispatch(impression) {
		return true
	}//Check Tag Type for URL

	return false
}