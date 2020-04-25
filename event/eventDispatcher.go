package event

import (
	"fmt"
	"strconv"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// Dispatch function dispatches the event represented by the impression object to our servers
func Dispatch(vwoInstance schema.VwoInstance, impression schema.Impression) {
	/*
		Args:
			impression: impression to be dispatched
		Returns:
	*/

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
	file := "eventDispatcher.go"
	if err != nil {
		message := fmt.Sprintf(constants.ErrorMessagesImpressionFailed, err)
		utils.LogMessage(vwoInstance,constants.Error, file, message)
	} else {
		message := fmt.Sprintf(constants.InfoMessageImpressionSuccess, impression)
		utils.LogMessage(vwoInstance, constants.Info, file, message)
	}
}

// DispatchTrackingGoal function dispatches the event with goal tracking represented by the impression object to our servers
func DispatchTrackingGoal(vwoInstance schema.VwoInstance, impression schema.Impression) {
	/*
		Args:
			impression: impression to be dispatched
		Returns:
	*/

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
		file := "eventDispatcher.go"
		message := fmt.Sprintf(constants.ErrorMessagesImpressionFailed, err)
		utils.LogMessage(vwoInstance,constants.Error, file, message)
	}
}
