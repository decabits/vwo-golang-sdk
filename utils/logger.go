package utils

import (
   "github.com/decabits/vwo-golang-sdk/schema"
   "github.com/decabits/vwo-golang-sdk/constants"
)

// LogMessage function generates Log messages and logs them into the logger, logger can be defined by the user itself too
func LogMessage(vwoInstance schema.VwoInstance, level, file, message string) {
	/*
		Args:
			file: Name of file from where the function is called
			message: Message to be logged
			level: level of logging
	*/

	var log string
	log = string(file) + " : " + message

	switch level {
	case constants.Info:
		vwoInstance.Logger.Info(log)
	case constants.Error:
		vwoInstance.Logger.Error(log)
	default:
		vwoInstance.Logger.Error("INVALID LEVEL TYPE")
	}
}
