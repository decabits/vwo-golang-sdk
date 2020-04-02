package event

import (
	"fmt"
	"strconv"

	log "github.com/golang/glog"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Dispatch ...
func Dispatch(impression schema.Impression) bool {
	URL := impression.URL + "?" +
		"random=" + strconv.FormatFloat(float64(impression.Random), 'f', -1, 32) +
		"&sdk=" + impression.Sdk +
		"&sdk-v=" + impression.SdkV +
		"&ap=" + impression.Ap +
		"&sId=" + impression.SID +
		"&u=" + impression.U +
		"&account_id=" + strconv.Itoa(impression.AccountID) +
		"&uId=" + impression.UID +
		"&goal_id=" + strconv.Itoa(impression.GoalID) +
		"&experiment_id=" + strconv.Itoa(impression.ExperimentID) +
		"&combination=" + impression.Combination +
		"&r=" + strconv.Itoa(impression.R) +
		"&ed=" + string(impression.ED)

	response, err := utils.GetRequest(URL)
	if err != nil {
		log.Error("ERROR_MESSAGES.IMPRESSION_FAILED")
		return false
	}
	fmt.Println(response)
	return true
}
