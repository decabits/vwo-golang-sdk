package vwo

import (
	log "github.com/golang/glog"

	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/service"
)

// VWO struct
type VWO struct {
	vwoInstance schema.VwoInstance
}

// New Function
func New(settingsFileLocation string, storage schema.UserStorage) schema.VwoInstance {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.ProcessSettingsFile(settingsFileLocation); err != nil {
		log.Error("Error Processing Settings File: " + err.Error())
	}
	settingsFile := settingsFileManager.GetSettingsFile()
	vwoInstance := schema.VwoInstance{
		SettingsFile:      settingsFile,
		UserStorage:       storage,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}

// Default function
func Default(accountID, SDKKey string, storage schema.UserStorage) schema.VwoInstance {
	settingsFileManager := service.SettingsFileManager{}
	if err := settingsFileManager.FetchSettingsFile(accountID, SDKKey); err != nil {
		log.Error("Error Processing Settings File: " + err.Error())
	}
	settingsFile := settingsFileManager.GetSettingsFile()
	vwoInstance := schema.VwoInstance{
		SettingsFile:      settingsFile,
		UserStorage:       storage,
		IsDevelopmentMode: true,
	}
	return vwoInstance
}

// func main() {
// 	settingsFileManager := service.SettingsFileManager{}
// 	if err := settingsFileManager.FetchSettingsFile("89499", "7aeed7f67f5a0b0fbe476c1f086a7038"); err != nil {
// 		log.Error("Error Processing Settings File: " + err.Error())
// 	}
// 	settingsFile := settingsFileManager.GetSettingsFile()

// 	fmt.Println(settingsFile)
// 	fmt.Println(settingsFile.Campaigns[0].Goals)
// 	fmt.Println(settingsFile.Campaigns[1].Variations[0])
// 	fmt.Println(settingsFile.Campaigns[1].Variations[1])
// 	fmt.Println(settingsFile.Campaigns[2].Variations[0])
// 	fmt.Println(settingsFile.Campaigns[2].Variations[2])
// }
