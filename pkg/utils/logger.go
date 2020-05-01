/*
   Copyright 2019-2020 Wingify Software Pvt. Ltd.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package utils

import (
	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/google/logger"
)

// LogMessage function generates Log messages and logs them into the logger, logger can be defined by the user itself too
func LogMessage(logs interface{}, level, file, message string) {
	/*
		Args:
			file: Name of file from where the function is called
			message: Message to be logged
			level: level of logging
	*/

	formattedMessage := string(file) + " : " + message

	if customlog, ok := logs.(interface {
		CustomLog(a, b string)
	}); ok {
		customlog.CustomLog(level, formattedMessage)
		return
	}

	log := logs.(*logger.Logger)

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
