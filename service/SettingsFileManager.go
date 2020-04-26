package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/decabits/vwo-golang-sdk/utils"
	"github.com/google/logger"
)

// SettingsFileManager struct to implement SettingsFileM
type SettingsFileManager struct {
	SettingsFile schema.SettingsFile
}

// FetchSettingsFile function makes call to VWO server to fetch the settings file
func (sfm *SettingsFileManager) FetchSettingsFile(accountID, SDKKey string) error {
	/*
		Args:
			accountID: Config account ID
			SDKKey: Config SDK Key

		Returns: 
			error: nil if the settings file id fetched else the error
	*/

	if accountID == "" {
		return fmt.Errorf(constants.ErrorMessagesInvalidAccountID)
	}
	if SDKKey == "" {
		return fmt.Errorf(constants.ErrorMessagesInvalidSDKKey)
	}

	protocol := constants.HTTPSProtocol
	hostname := constants.BaseURL
	path := constants.AccountSettings + "?" +
		"a=" + accountID +
		"&i=" + SDKKey +
		"&r=" + strconv.FormatFloat(float64(rand.Float32()), 'f', -1, 64) +
		"&platform=server&" +
		"&sdk=" + constants.SDKName +
		"&sdk-v=" + constants.SDKVersion +
		"&api-version=1"

	resp, err := utils.GetRequest(protocol + hostname + path)
	if err != nil {
		return fmt.Errorf(constants.ErrorMessagesSettingsFileCorrupted, err.Error())
	}

	if err = json.Unmarshal([]byte(resp), &sfm.SettingsFile); err != nil {
		return fmt.Errorf(constants.ErrorMessagesInvalidSettingsFile, err.Error())
	}
	logger.Warning(constants.DebugMessagesValidConfiguration)
	return nil
}

// ProcessSettingsFile Processes the settings_file, assigns variation allocation range
func (sfm *SettingsFileManager) ProcessSettingsFile(settingsFileLocation string) error {
	/*
		Args:
			settingsFileLocation: Location of the settings file on system

		Returns: 
			error: nil if the settings file id fetched else the error
	*/

	settingsFile, err := ioutil.ReadFile(settingsFileLocation)
	if err != nil {
		return fmt.Errorf(constants.ErrorMessagesCannotReadSettingsFile, err.Error())
	}

	if err = json.Unmarshal(settingsFile, &sfm.SettingsFile); err != nil {
		return fmt.Errorf(constants.ErrorMessagesInvalidSettingsFile, err.Error())
	}

	return nil
}

//Process function processes campaigns in the settings file and sets the variation allocation ranges to all variations 
func (sfm *SettingsFileManager) Process() {
	logs := logger.Init(constants.SDKName, true, false, ioutil.Discard)
	logger.SetFlags(log.LstdFlags)
	defer logger.Close()
	for i, campaign := range sfm.SettingsFile.Campaigns {
		var (
			currentAllocation         = 0
			variationAllocationRanges []schema.Variation
		)
		for _, variation := range campaign.Variations {
			stepFactor := utils.GetVariationBucketingRange(variation.Weight)
			if stepFactor != 0 {
				variation.StartVariationAllocation = currentAllocation + 1
				variation.EndVariationAllocation = currentAllocation + stepFactor
				currentAllocation += stepFactor
			} else {
				variation.StartVariationAllocation = -1
				variation.EndVariationAllocation = -1
			}
			logs.Infof(constants.InfoMessageVariationRageAllocation, variation.Name, variation.Weight, variation.StartVariationAllocation, variation.EndVariationAllocation)
			variationAllocationRanges = append(variationAllocationRanges, variation)
		}
		sfm.SettingsFile.Campaigns[i].Variations = variationAllocationRanges
	}
}

// GetSettingsFile returns the settings file 
func (sfm *SettingsFileManager) GetSettingsFile() schema.SettingsFile {
	return sfm.SettingsFile
}
