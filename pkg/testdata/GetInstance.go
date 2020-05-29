package testdata

import (
	"encoding/json"
	"io/ioutil"

	"github.com/decabits/vwo-golang-sdk/pkg/constants"
	"github.com/decabits/vwo-golang-sdk/pkg/logger"
	"github.com/decabits/vwo-golang-sdk/pkg/schema"
)

//GetInstanceWithCustomSettings function
func GetInstanceWithCustomSettings(SettingsFileName string) schema.VwoInstance {
	logger := logger.Init(constants.SDKName, false, false, ioutil.Discard)
	defer logger.Close()

	var settingsFiles map[string]schema.SettingsFile
	data, err := ioutil.ReadFile("../testdata/customSettings.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &settingsFiles); err != nil {
		logger.Info("Error: " + err.Error())
	}

	settings := settingsFiles[SettingsFileName]
	

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settings,
		UserStorage:       nil,
		Logger:            logger,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}

//GetInstanceWithSettings function
func GetInstanceWithSettings(SettingsFileName string) schema.VwoInstance {
	logger := logger.Init(constants.SDKName, false, false, ioutil.Discard)
	defer logger.Close()

	var settingsFiles map[string]schema.SettingsFile
	data, err := ioutil.ReadFile("../testdata/settings.json")
	if err != nil {
		logger.Info("Error: " + err.Error())
	}

	if err = json.Unmarshal(data, &settingsFiles); err != nil {
		logger.Info("Error: " + err.Error())
	}

	settings := settingsFiles[SettingsFileName]
	

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settings,
		UserStorage:       nil,
		Logger:            logger,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}
