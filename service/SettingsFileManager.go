package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
)

// SettingsFileM interface
type SettingsFileM interface {
	FetchSettingsFile(accountID, SDKKey string) error
	ProcessSettingsFile(settingsFileLocation string) error
	GetSettingsFile() schema.SettingsFile
}

// SettingsFileManager struct to implement SettingsFileM
type SettingsFileManager struct {
	SettingsFile schema.SettingsFile
}

// FetchSettingsFile function
func (sfm *SettingsFileManager) FetchSettingsFile(accountID, SDKKey string) error {
	if accountID == "" || SDKKey == "" {
		return errors.New("AccountId and sdkKey are required for fetching account settings. Aborting")
	}

	protocol := constants.HTTPSProtocol
	hostname := constants.BaseURL
	path := constants.AccountSettings + "?" +
		"a=" + accountID +
		"&i=" + SDKKey +
		"&r=" + strconv.FormatFloat(float64(rand.Float32()), 'f', -1, 32) +
		"&platform=server&" +
		"&sdk=" + constants.SDKName +
		"&sdk-v=" + constants.SDKVersion +
		"&api-version=1"

	resp, err := utils.Get(protocol + hostname + path)
	if err != nil {
		return errors.New("Error fetching Settings File: " + err.Error())
	}

	fmt.Println(resp)

	if err = json.Unmarshal([]byte(resp), &sfm.SettingsFile); err != nil {
		return errors.New("Error: " + err.Error())
	}

	return nil
}

// ProcessSettingsFile Processes the settings_file, assigns variation allocation range
func (sfm *SettingsFileManager) ProcessSettingsFile(settingsFileLocation string) error {
	settingsFile, err := ioutil.ReadFile(settingsFileLocation)
	if err != nil {
		return errors.New("Error: " + err.Error())
	}

	if err = json.Unmarshal(settingsFile, &sfm.SettingsFile); err != nil {
		return errors.New("Error: " + err.Error())
	}

	// for _, campaign := range sfm.SettingsFile.Campaigns {
	// 	utils.SetVariationAllocation(campaign)
	// }
	return nil
}

// GetSettingsFile ...
func (sfm *SettingsFileManager) GetSettingsFile() schema.SettingsFile {
	return sfm.SettingsFile
}
