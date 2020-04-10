package vwo

import (
	"io/ioutil"
	"log"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/service"
	"github.com/google/logger"
)

// VWO struct
type VWO struct {
	vwoInstance schema.VwoInstance
}

// New Function
func New(settingsFileLocation string, storage schema.UserStorage) schema.VwoInstance {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.ProcessSettingsFile(settingsFileLocation); err != nil {
		log.Println("Error Processing Settings File: ", err)
	}
	settingsFileManager.Process()
	settingsFile := settingsFileManager.GetSettingsFile()

	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settingsFile,
		UserStorage:       storage,
		Logger:            logs,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}

// Default function
func Default(accountID, SDKKey string, storage schema.UserStorage) schema.VwoInstance {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.FetchSettingsFile(accountID, SDKKey); err != nil {
		log.Println("Error Processing Settings File: ", err)
	}
	settingsFileManager.Process()
	settingsFile := settingsFileManager.GetSettingsFile()

	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settingsFile,
		UserStorage:       storage,
		Logger:            logs,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}
