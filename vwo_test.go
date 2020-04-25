package vwo

import (
	"log"

	"github.com/decabits/vwo-golang-sdk/service"
)

// GetInstance function creates and return a temporary VWO instance for testing
func (vwo *VWOInstance) getInstance(path string) {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.ProcessSettingsFile(path); err != nil {
		log.Println("Error Processing Settings File: ", err)
	}
	settingsFileManager.Process()
	settingsFile := settingsFileManager.GetSettingsFile()

	vwo.Launch(true, settingsFile, nil)
}
