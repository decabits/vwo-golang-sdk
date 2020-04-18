package event

import (
	"strconv"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Dispatch ...
func Dispatch(vwoInstance schema.VwoInstance, impression schema.Impression) {
	URL := impression.URL + "?" +
		"random=" + strconv.FormatFloat(float64(impression.Random), 'f', -1, 64) +
		"&sdk=" + impression.Sdk +
		"&sdk-v=" + impression.SdkV +
		"&ap=" + impression.Ap +
		"&sId=" + impression.SID +
		"&u=" + impression.U +
		"&account_id=" + strconv.Itoa(impression.AccountID) +
		"&uId=" + impression.UID +
		"&experiment_id=" + strconv.Itoa(impression.ExperimentID) +
		"&combination=" + strconv.Itoa(impression.Combination) +
		"&ed=" + impression.ED

	_, err := utils.GetRequest(URL)
	if err != nil {
		vwoInstance.Logger.Errorf("ERROR_MESSAGES.IMPRESSION_FAILED %+v", err)
	}
}

// DispatchTrackingGoal function
func DispatchTrackingGoal(vwoInstance schema.VwoInstance, impression schema.Impression) {
	URL := impression.URL + "?" +
		"random=" + strconv.FormatFloat(float64(impression.Random), 'f', -1, 64) +
		"&sdk=" + impression.Sdk +
		"&sdk-v=" + impression.SdkV +
		"&ap=" + impression.Ap +
		"&sId=" + impression.SID +
		"&u=" + impression.U +
		"&account_id=" + strconv.Itoa(impression.AccountID) +
		"&uId=" + impression.UID +
		"&experiment_id=" + strconv.Itoa(impression.ExperimentID) +
		"&combination=" + strconv.Itoa(impression.Combination) +
		"&goal_id=" + strconv.Itoa(impression.GoalID) +
		"&r=" + strconv.Itoa(impression.R)

	_, err := utils.GetRequest(URL)
	if err != nil {
		vwoInstance.Logger.Errorf("ERROR_MESSAGES.IMPRESSION_FAILED %+v", err)
	}
}
