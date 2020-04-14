package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/google/logger"
)

// GetInstance function creates and return a temporary VWO instance for testing
func GetInstance(path string) schema.VwoInstance {
	logger := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	defer logger.Close()

	settingsFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var settings schema.SettingsFile
	if err = json.Unmarshal(settingsFile, &settings); err != nil {
		fmt.Println(err)
	}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settings,
		Logger:            logger,
		IsDevelopmentMode: true,
	}

	return vwoInstance
}
