package event

import (
	"strconv"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Dispatch ...
func Dispatch(vwoInstance schema.VwoInstance, impression schema.Impression) bool {
	URL := impression.URL + "?" +
		"random=" + strconv.FormatFloat(float64(impression.Random), 'f', -1, 32) +
		"&sdk=" + impression.Sdk +
		"&sdk-v=" + impression.SdkV +
		"&ap=" + impression.Ap +
		"&sId=" + impression.SID +
		"&u=" + impression.U +
		"&account_id=" + strconv.Itoa(impression.AccountID) +
		"&uId=" + string(impression.UID) +
		"&experiment_id=" + strconv.Itoa(impression.ExperimentID) +
		"&combination=" + strconv.Itoa(impression.Combination) +
		"&ed=" + string(impression.ED)

	_, err := utils.GetRequest(URL)
	if err != nil {
		vwoInstance.Logger.Error("ERROR_MESSAGES.IMPRESSION_FAILED %+v", err)
		return false
	}
	return true
}

// DispatchTrackingGoal function
func DispatchTrackingGoal(vwoInstance schema.VwoInstance, impression schema.Impression) bool {
	URL := impression.URL + "?" +
		"random=" + strconv.FormatFloat(float64(impression.Random), 'f', -1, 32) +
		"&sdk=" + impression.Sdk +
		"&sdk-v=" + impression.SdkV +
		"&ap=" + impression.Ap +
		"&sId=" + impression.SID +
		"&u=" + impression.U +
		"&account_id=" + strconv.Itoa(impression.AccountID) +
		"&uId=" + string(impression.UID) +
		"&experiment_id=" + strconv.Itoa(impression.ExperimentID) +
		"&combination=" + strconv.Itoa(impression.Combination) +
		"&goal_id=" + strconv.Itoa(impression.GoalID) +
		"&r=" + strconv.Itoa(impression.R)

	_, err := utils.GetRequest(URL)
	if err != nil {
		vwoInstance.Logger.Error("ERROR_MESSAGES.IMPRESSION_FAILED %+v", err)
		return false
	}
	return true
}
