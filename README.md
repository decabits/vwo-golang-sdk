# VWO-Golang-SDK

[![Build Status](https://img.shields.io/travis/decabits/vwo-golang-sdk)](http://travis-ci.org/decabits/vwo-golang-sdk)
![Size in Bytes](https://img.shields.io/github/languages/code-size/decabits/vwo-golang-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)
![Coveralls github](https://img.shields.io/coveralls/github/decabits/vwo-golang-sdk)

This open source library allows you to A/B Test your Website at server-side.

## Requirements

- Works with Go 1.11 +

## Installation

```go
go get "github.com/decabits/vwo-golang-sdk"
```

## Basic usage

**Importing and Instantiation**

```go
import vwo "github.com/decabits/vwo-golang-sdk"
import "github.com/decabits/vwo-golang-sdk/pkg/api"

// Get SettingsFile
settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")

// Default instance of VwoInstance
instance, err := vwo.Launch(settingsFile)
if err != nil {
	//handle err
}

// Instance with custom options
instance, err := vwo.Launch(settingsFile, api.WithDevelopmentMode())
if err != nil {
	//handle err
}

// Activate API
// With Custom Variables
options := make(map[string]interface{})
options["customVariables"] = map[string]interface{}{"a": "x"}
options["variationTargetingVariables"] = map[string]interface{}{"a": "x"}
options["revenueValue"] = 12
variationName = vwoInstance.Activate(campaignKey, userID, options)

// Without Custom Variables
variationName = vwoInstance.Activate(campaignKey, userID, nil)


// GetVariation
// With Custom Variables
options := make(map[string]interface{})
options["customVariables"] = map[string]interface{}{"a": "x"}
variationName = vwoInstance.GetVariationName(campaignKey, userID, options)

//Without Custom Variables
variationName = vwoInstance.GetVariationName(campaignKey, userID, nil)


// Track API
// With Custom Variables
options := make(map[string]interface{})
options["customVariables"] = map[string]interface{}{"a": "x"}
isSuccessful = vwoInstance.Track(campaignKey, userID, goalIdentifier, options)

// With Revenue Value
options := make(map[string]interface{})
options["revenueValue"] = 12
isSuccessful = vwoInstance.Track(campaignKey, userID, goalIdentifier, options)

// With both Custom Variables and Revenue Value
options := make(map[string]interface{})
options["customVariables"] = map[string]interface{}{"a": "x"}
options["revenueValue"] = 12
isSuccessful = vwoInstance.Track(campaignKey, userID, goalIdentifier, options)

//Without Custom Variables
isSuccessful = vwoInstance.Track(campaignKey, userID, goalIdentifier, nil)

// FeatureEnabled API
// With Custom Varibles
options := make(map[string]interface{})
options["customVariables"] = map[string]interface{}{"a": "x"}
isSuccessful = vwoInstance.IsFeatureEnabled(campaignKey, userID, options)

// Without Custom Variables
isSuccessful = vwoInstance.IsFeatureEnabled(campaignKey, userID, nil)

// GetFeatureVariableValue API
// With Custom Variables
options := make(map[string]interface{})
options["customVariables"] = map[string]interface{}{"a": "x"}
variableValue = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, options)

// Without Custom Variables
variableValue = vwoInstance.GetFeatureVariableValue(campaignKey, variableKey, userID, nil)

// Push API
isSuccessful = vwoInstance.Push(tagKey, tagValue, userID)
```

**User Storage**

```go
import vwo "github.com/decabits/vwo-golang-sdk/"
import "github.com/decabits/vwo-golang-sdk/pkg/api"
import "github.com/decabits/vwo-golang-sdk/pkg/schema"

// declare UserStorage interface with the following Get & Set function signature
type UserStorage interface{
    Get(userID, campaignKey string) UserData
    Set(string, string, string)
}

// declare a UserStorageData struct to implement UserStorage interface
type UserStorageData struct{}

// Get method to fetch user variation from storage
func (us *UserStorageData) Get(userID, campaignKey string) schema.UserData {
    //Example code showing how to get userData  from DB
    userData, ok := userDatas[campaignKey]
    if ok {
		for _, userdata := range userData {
			if userdata.UserID == userID {
				return userdata
			}
		}
    }
    /*
    // UserData  struct
    type UserData struct {
        UserID        string
        CampaignKey   string
        VariationName string
    }
    */
	return schema.UserData{}
}

// Set method to save user variation to storage
func (us *UserStorageData) Set(userID, campaignKey, variationName string) {
    //Example code showing how to store userData in DB
    userdata := schema.UserData{
		UserID:        userID,
		CampaignKey:   campaignKey,
		VariationName: variationName,
	}
	flag := false
	userData, ok := userDatas[userdata.CampaignKey]
	if ok {
		for _, user := range userData {
			if user.UserID == userdata.UserID {
				flag = true
			}
		}
		if !flag {
			userDatas[userdata.CampaignKey] = append(userDatas[userdata.CampaignKey], userdata)
		}
	} else {
		userDatas[userdata.CampaignKey] = []schema.UserData{
			userdata,
		}
	}
}

func main() {
	settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")
	// create UserStorageData object
	storage := &UserStorageData{}

	instance, err := vwo.Launch(settingsFile, api.WithStorage(storage))
	if err != nil {
		//handle err
	}
}

```

**Custom Logger**

```go
import vwo "github.com/decabits/vwo-golang-sdk"
import "github.com/decabits/vwo-golang-sdk/pkg/api"

// declare Log interface with the following CustomLog function signature
type Log interface {
	CustomLog(level, errorMessage string)
}

// declare a LogS struct to implement Log interface
type LogS struct{}

// Get function to handle logs
func (c *LogS) CustomLog(level, errorMessage string) {}

func main() {
	settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")
	// create LogS object
	logger := &LogS{}

	instance, err := vwo.Launch(settingsFile, api.WithLogger(logger))
	if err != nil {
		//handle err
	}
}
```

## Demo App

[Example](https://github.com/decabits/vwo-golang-example-app)

## Documentation

Refer [Official VWO Documentation](https://developers.vwo.com/reference#server-side-introduction)

## Running Unit Tests

```shell
./test.sh
```

## Third-party Resources and Credits

Refer [third-party-attributions.txt](third-party-attribution.txt)

## Authors

- [Piyushh bhutoria](https://github.com/Piyushhbhutoria)
- [Vaibhav Sethia](https://github.com/vaibhavsethia)

## Contributing

Please go through our [contributing guidelines](CONTRIBUTING.md)

## Code of Conduct

[Code of Conduct](CODE_OF_CONDUCT.md)

## License

[Apache License, Version 2.0](LICENSE)
