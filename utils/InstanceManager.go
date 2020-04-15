package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/decabits/vwo-golang-sdk/constants"
	"github.com/decabits/vwo-golang-sdk/schema"
	"github.com/google/logger"
)

// UserStorage interface for testing
type UserStorage schema.UserStorage

// UserStorageData struct for testing
type UserStorageData struct{}

// Get function is used to get the data from user storage 
func (us *UserStorageData) Get(userID, campaignKey string) schema.UserData {
	return schema.UserData{}
}

// Set function
func (us *UserStorageData) Set(userID, campaignKey, variationName string) {
	_ = schema.UserData{
		UserID:        userID,
		CampaignKey:   campaignKey,
		VariationName: variationName,
	}
}

// Exist function
func (us *UserStorageData) Exist() bool {
	return false
}

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

	storage := &UserStorageData{}

	for i, campaign := range settings.Campaigns {
		var (
			currentAllocation         = 0
			variationAllocationRanges []schema.Variation
		)
		for _, variation := range campaign.Variations {
			stepFactor := GetVariationBucketingRange(variation.Weight)
			if stepFactor != 0 {
				variation.StartVariationAllocation = currentAllocation + 1
				variation.EndVariationAllocation = currentAllocation + stepFactor
				currentAllocation += stepFactor
			} else {
				variation.StartVariationAllocation = -1
				variation.EndVariationAllocation = -1
			}
			variationAllocationRanges = append(variationAllocationRanges, variation)
		}
		settings.Campaigns[i].Variations = variationAllocationRanges
	}

	vwoInstance := schema.VwoInstance{
		SettingsFile:      settings,
		UserStorage:       storage,
		Logger:            logger,
		IsDevelopmentMode: true,
	}

	return vwoInstance
}
