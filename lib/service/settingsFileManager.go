package service

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/decabits/vwo-golang-sdk/lib/schema"
	"github.com/decabits/vwo-golang-sdk/lib/utils"
)

// SettingsFileM interface 
type SettingsFileM interface {
	ProcessSettingsFile(settingsFileLocation string)
	GetSettingsFile() schema.SettingsFile
} 

// SettingsFileManager struct to implement SettingsFileM
type SettingsFileManager struct {
	SettingsFile schema.SettingsFile
}

// ProcessSettingsFile Processes the settings_file, assigns variation allocation range
func (sfm *SettingsFileManager) ProcessSettingsFile(settingsFileLocation string) {
	settingsFile, err := ioutil.ReadFile(settingsFileLocation)
	if err != nil {
		log.Println("Error: ", err)
	}

	if err = json.Unmarshal(settingsFile, &sfm.SettingsFile); err != nil {
		log.Println("Error: ", err)
	}

	for _, campaign := range sfm.SettingsFile.Campaigns {
		utils.SetVariationAllocation(campaign)
	}
}

//GetSettingsFile ...
func (sfm *SettingsFileManager) GetSettingsFile() schema.SettingsFile {
	return sfm.SettingsFile
}
