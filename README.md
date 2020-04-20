# VWO-Golang-SDK

[![Build Status](http://img.shields.io/travis/decabits/vwo-golang-sdk/master.svg?style=flat)](https://img.shields.io/travis/decabits/vwo-golang-sdk)
[![Coverage Status](https://coveralls.io/repos/github/decabits/vwo-golang-sdk/badge.svg)](https://img.shields.io/coveralls/GitHub/decabits/vwo-golang-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)

This open source library allows you to A/B Test your Website at server-side.
https://img.shields.io/github/license/decabits/vwo-golang-sdk
## Requirements

* Works with Go 1.1x


## Installation

```go
go get "github.com/decabits/vwo-golang-sdk"
```


## Basic usage

**Importing and Instantiation**

```go
import (
	vwo "github.com/decabits/vwo-golang-sdk"
    "github.com/decabits/vwo-golang-sdk/schema"
    "github.com/decabits/vwo-golang-sdk/api"
    "github.com/decabits/vwo-golang-sdk/service"
)

// Initialize client
// storage should be of type schema.UserStorage
VWO := vwo.Default("accountID", "SDKKey", storage)

// Get Settings
settingsFile = VWO.SettingsFile

// Activate API
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" : "x"},
    }
variationName = api.ActivateWithOptions(VWO, campaignKey, userID, options)

// Without Custom Variables
variationName = api.Activate(VWO, campaignKey, userID)


// GetVariation
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" : "x"},
    }
variationName = api.GetVariationName(VWO, campaignKey, userID, options)

//Without Custom Variables
options = {}
variationName = api.GetVariationName(VWO, campaignKey, userID, options)


// Track API
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" : "x"},
    }
isSuccessful = api.TrackWithOptions(VWO, campaignKey, userID, goalIdentifier, options)

// With Revenue Value
options = schema.Options{
        RevenueGoal => 10.23,
    }
isSuccessful = api.TrackWithOptions(VWO, campaignKey, userID, goalIdentifier, options)

// With both Custom Variables and Revenue Value
options = schema.Options{
        CustomVariables : { "a" : "x"},
        RevenueGoal : 10.23,
    }
isSuccessful = api.TrackWithOptions(VWO, campaignKey, userID, goalIdentifier, options)


// FeatureEnabled API
// With Custom Varibles
options = schema.Options{
        CustomVariables : { "a" : "x"},
    }
isSuccessful = api.IsFeatureEnabled(VWO, campaignKey, userID, options)

// Without Custom Variables
options = {}
isSuccessful = api.IsFeatureEnabled(VWO, campaignKey, userID, options)


// GetFeatureVariableValue API
// With Custom Variables
options = schema.Options{
        CustomVariables : { "a" : "x"},
    }
variableValue = api.GetFeatureVariableValue(VWO, campaignKey, variableKey, userID, options)

// Without Custom Variables
options = {}
variableValue = api.GetFeatureVariableValue(VWO, campaignKey, variableKey, userID, options)

// Push API
isSuccessful = api.Push(tagKey, tagValue, userID)
```

1. `accountID` - Account for which sdk needs to be initialized
1. `SDKKey` - SDK key for that account
1. `logger` - If you need to pass your own logger. Check documentation below
1. `UserStorage.new` - An object allowing `get` and `set` for maintaining user storage
1. `developmentMode` - on/off (true/false). Default - false
1. `settingsFile` - Settings file if already present during initialization. Its stringified JSON format.


**User Storage**

```go

import "github.com/decabits/vwo-golang-sdk/schema"

// UserStorage interface
type UserStorage schema.UserStorage

// UserStorageData struct
type UserStorageData struct{}

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
	return schema.UserData{}
}

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

```


## Documentation

Refer [Official VWO Documentation](https://developers.vwo.com/reference#server-side-introduction)


## Running Unit Tests

```shell
./test.sh
```


## Authors



## Contributing

Please go through our [contributing guidelines](CONTRIBUTING.md)


## Code of Conduct

[Code of Conduct](CODE_OF_CONDUCT.md)


## License

[Apache License, Version 2.0](LICENSE)
