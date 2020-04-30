# VWO-Golang-SDK

[![Build Status](https://img.shields.io/travis/decabits/vwo-golang-sdk)](http://travis-ci.org/decabits/vwo-golang-sdk)
![Size in Bytes](https://img.shields.io/github/languages/code-size/decabits/vwo-golang-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)
[![Coverage Status](https://coveralls.io/repos/github/decabits/vwo-golang-sdk/badge.svg)](https://coveralls.io/github/decabits/vwo-golang-sdk)

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

// Get SettingsFile
settingsFile := vwo.GetSettingsFile("accountID", "SDKKey")

// Declaration of VwoInstance
vwoInstance := vwo.VWOInstance{}

// Create VwoInstance and handle error if any
err := vwoInstance.Launch("isDevelopmentMode", settingsFile, nil, nil)

// Activate API
// With Custom Variables
options := make(map[string]interface{})
options["customVariables"] = map[string]interface{}{"a": "x"}
options["variationTargetingVariables"] = map[string]interface{}{"a": "x"}
options["revenueGoal"] = 12
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
options["revenueGoal"] = 12
isSuccessful = vwoInstance.Track(campaignKey, userID, goalIdentifier, options)

// With both Custom Variables and Revenue Value
options := make(map[string]interface{})
options["customVariables"] = map[string]interface{}{"a": "x"}
options["revenueGoal"] = 12
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

1. `accountID` - Account for which sdk needs to be initialized
2. `SDKKey` - SDK key for that account
3. `logger` - If you need to pass your own logger. Check documentation below
4. `UserStorage.new` - An object allowing `get` and `set` for maintaining user storage
5. `developmentMode` - on/off (true/false). Default - false
6. `settingsFile` - Settings file if already present during initialization. Its stringified JSON format.

**User Storage**

```go
import "github.com/decabits/vwo-golang-sdk/schema"

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
	v.vwoInstance = vwo.VWOInstance{}
	err := v.vwoInstance.Launch(config.GetBool("isDevelopmentMode"), settingsFile, storage)
	if err != nil {
		fmt.Println("error intialising sdk")
	}
}

```

**Custom Logger**

```go
import vwo "github.com/decabits/vwo-golang-sdk"

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
	v.vwoInstance = vwo.VWOInstance{}
	err := v.vwoInstance.LaunchWithLogger(config.GetBool("isDevelopmentMode"), settingsFile, nil, logger)
	if err != nil {
		fmt.Println("error intialising sdk")
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
