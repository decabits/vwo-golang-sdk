package utils

import (
	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/google/logger"
)

// LogMessage function generates Log messages and logs them into the logger, logger can be defined by the user itself too
func LogMessage(log *logger.Logger, level, file, message string) {
	/*
		Args:
			file: Name of file from where the function is called
			message: Message to be logged
			level: level of logging
	*/

	formattedMessage := string(file) + " : " + message

	switch level {
	case constants.Info:
		log.Info(formattedMessage)
	case constants.Debug:
		log.Warning(formattedMessage)
	case constants.Error:
		log.Error(formattedMessage)
	default:
		log.Error("Invalid Logger Level")
	}
}